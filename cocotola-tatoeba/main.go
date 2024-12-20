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
	rsusergateway "github.com/kujilabo/cocotola-1.23/redstart/user/gateway"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller"
	liblog "github.com/kujilabo/cocotola-1.23/lib/log"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/config"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/initialize"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/sqls"
)

const (
	loggerKey = liblog.TatoebaMainLoggerContextKey
)

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

	ctx = liblog.InitLogger(ctx)
	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	// init repository factory function
	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		return gateway.NewRepositoryFactory(ctx, dialect, cfg.DB.DriverName, db, time.UTC) // nolint:wrapcheck
	}

	// init repository factory
	rf, err := rff(ctx, db)
	if err != nil {
		panic(err)
	}

	// init rs repository factory
	rsrf, err := rsusergateway.NewRepositoryFactory(ctx, dialect, cfg.DB.DriverName, db, time.UTC)
	if err != nil {
		panic(err)
	}

	// init transaction manager
	txManager, err := gateway.NewTransactionManager(db, rff)
	if err != nil {
		panic(err)
	}

	// init non transaction manager
	nonTxManager, err := gateway.NewNoneTransactionManager(rf)
	if err != nil {
		panic(err)
	}

	// init gin
	if !cfg.Debug.Gin {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	if err := initialize.InitAppServer(ctx, router, *cfg.InternalAuth, cfg.CORS, cfg.Debug, cfg.App.Name, txManager, nonTxManager, rsrf); err != nil {
		panic(err)
	}

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
		return libcontroller.AppServerProcess(ctx, loggerKey, router, cfg.App.HTTPPort, time.Duration(cfg.App.ReadHeaderTimeoutSec)*time.Second, time.Duration(cfg.Shutdown.TimeSec1)*time.Second) // nolint:wrapcheck
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
