package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/initialize"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/service"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/sqls"
)

// const (
// 	loggerKey = liblog.SynthesizerMainLoggerContextKey
// )

func getValue(values ...string) string {
	for _, v := range values {
		if len(v) != 0 {
			return v
		}
	}
	return ""
}

func main() {
	ctx := context.Background()
	env := flag.String("env", "", "environment")
	flag.Parse()
	appEnv := getValue(*env, os.Getenv("APP_ENV"), "local")
	slog.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	rsliberrors.UseXerrorsErrorf()

	// load config
	cfg, err := config.LoadConfig(appEnv)
	if err != nil {
		panic(err)
	}

	// init log
	if err := rslibconfig.InitLog(cfg.Log); err != nil {
		panic(err)
	}

	// init tracer
	tp, err := rslibconfig.InitTracerProvider(ctx, cfg.App.Name, cfg.Trace)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("err. err: %v", err))
		slog.ErrorContext(ctx, fmt.Sprintf("err. err: %+v", err))
		panic(err)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// init db
	dialect, db, sqlDB, err := rslibconfig.InitDB(ctx, cfg.DB, sqls.SQL)
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	defer tp.ForceFlush(ctx) // flushes any pending spans

	logger := slog.Default().With(slog.String(rsliblog.LoggerNameKey, "main"))

	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		return gateway.NewRepositoryFactory(ctx, dialect, cfg.DB.DriverName, db, time.UTC) // nolint:wrapcheck
	}
	rf, err := rff(ctx, db)
	if err != nil {
		panic(err)
	}

	// init transaction manager
	txManager, err := rslibgateway.NewTransactionManagerT(db, rff)
	if err != nil {
		panic(err)
	}

	// init non transaction manager
	nonTxManager, err := rslibgateway.NewNonTransactionManagerT(rf)
	if err != nil {
		panic(err)
	}

	// init gin
	if !cfg.Debug.Gin {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	authMiddleware := controller.InitAuthMiddleware(cfg.InternalAuth)
	publicRouterGroupFuncs := controller.GetPublicRouterGroupFuncs()
	privateRouterGroupFuncs := controller.GetPrivateRouterGroupFuncs(cfg.TTS, txManager, nonTxManager)
	initialize.InitAppServer(ctx, router, cfg.CORS, cfg.Debug, cfg.App.Name, authMiddleware, publicRouterGroupFuncs, privateRouterGroupFuncs)

	// run
	result := run(ctx, cfg, router)

	gracefulShutdownTime2 := time.Duration(cfg.Shutdown.TimeSec2) * time.Second
	time.Sleep(gracefulShutdownTime2)
	logger.InfoContext(ctx, "exited")
	os.Exit(result)
}

func run(ctx context.Context, cfg *config.Config, router http.Handler) int {
	var eg *errgroup.Group
	eg, ctx = errgroup.WithContext(ctx)

	eg.Go(func() error {
		return libcontroller.AppServerProcess(ctx, router, cfg.App.HTTPPort, time.Duration(cfg.App.ReadHeaderTimeoutSec)*time.Second, time.Duration(cfg.Shutdown.TimeSec1)*time.Second) // nolint:wrapcheck
	})
	eg.Go(func() error {
		return rslibgateway.MetricsServerProcess(ctx, cfg.App.MetricsPort, cfg.Shutdown.TimeSec1) // nolint:wrapcheck
	})
	eg.Go(func() error {
		return rslibgateway.SignalWatchProcess(ctx) // nolint:wrapcheck
	})
	eg.Go(func() error {
		<-ctx.Done()
		return ctx.Err() // nolint:wrapcheck
	})

	if err := eg.Wait(); err != nil {
		return 1
	}
	return 0
}
