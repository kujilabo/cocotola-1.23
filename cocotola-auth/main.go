package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"gorm.io/gorm"

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	"github.com/kujilabo/cocotola-1.23/redstart/sqls"

	libgateway "github.com/kujilabo/cocotola-1.23/lib/gateway"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-auth/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/initialize"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/service"
)

func getValue(values ...string) string {
	for _, v := range values {
		if len(v) != 0 {
			return v
		}
	}
	return ""
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	env := flag.String("env", "", "environment")
	flag.Parse()
	appEnv := getValue(*env, os.Getenv("APP_ENV"), "local")

	rsliberrors.UseXerrorsErrorf()

	// load config
	cfg, err := config.LoadConfig(appEnv)
	checkError(err)

	// init log
	rslibconfig.InitLog(cfg.Log)
	logger := slog.Default().With(slog.String(rsliblog.LoggerNameKey, "main"))
	logger.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	// init tracer
	tp, err := rslibconfig.InitTracerProvider(ctx, cfg.App.Name, cfg.Trace)
	checkError(err)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// init db
	dialect, db, sqlDB, err := rslibconfig.InitDB(ctx, cfg.DB, sqls.SQL)
	checkError(err)
	defer sqlDB.Close()
	defer tp.ForceFlush(ctx) // flushes any pending spans

	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		return gateway.NewRepositoryFactory(ctx, dialect, cfg.DB.DriverName, db, time.UTC) // nolint:wrapcheck
	}
	rf, err := rff(ctx, db)
	checkError(err)

	// init transaction manager
	txManager, err := rslibgateway.NewTransactionManagerT(db, rff)
	checkError(err)

	// init non transaction manager
	nonTxManager, err := rslibgateway.NewNonTransactionManagerT(rf)
	checkError(err)

	initialize.InitApp1(ctx, txManager, nonTxManager, "cocotola", cfg.App.OwnerLoginID, cfg.App.OwnerPassword)

	// init gin
	router := initGin(ctx, cfg, txManager, nonTxManager)

	// run
	readHeaderTimeout := time.Duration(cfg.App.ReadHeaderTimeoutSec) * time.Second
	shutdownTime := time.Duration(cfg.Shutdown.TimeSec1) * time.Second
	result := libgateway.Run(ctx,
		libgateway.WithAppServerProcess(router, cfg.App.HTTPPort, readHeaderTimeout, shutdownTime),
		libgateway.WithSignalWatchProcess(),
		libgateway.WithMetricsServerProcess(cfg.App.MetricsPort, cfg.Shutdown.TimeSec1),
	)

	gracefulShutdownTime2 := time.Duration(cfg.Shutdown.TimeSec2) * time.Second
	time.Sleep(gracefulShutdownTime2)
	logger.InfoContext(ctx, "exited")
	os.Exit(result)
}

func initGin(ctx context.Context, cfg *config.Config, txManager, nonTxManager service.TransactionManager) *gin.Engine {
	if !cfg.Debug.Gin {
		gin.SetMode(gin.ReleaseMode)
	}

	publicRouterGroupFuncs := controller.GetPublicRouterGroupFuncs(cfg.Auth, txManager, nonTxManager)
	// privateRouterGroupFuncs := controller.InitPublicRouterGroupFuncs(cfg.Auth, txManager, nonTxManager)
	router := gin.New()
	initialize.InitAppServer(ctx, router, cfg.CORS, cfg.Debug, cfg.App.Name, publicRouterGroupFuncs)
	return router
}
