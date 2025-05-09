package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"
	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"
	libgateway "github.com/kujilabo/cocotola-1.23/lib/gateway"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/config"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/initialize"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/sqls"
)

func main() {
	ctx := context.Background()
	env := flag.String("env", "", "environment")
	flag.Parse()
	appEnv := libdomain.GetNonEmptyValue(*env, os.Getenv("APP_ENV"), "local")

	rsliberrors.UseXerrorsErrorf()

	// load config
	cfg, err := config.LoadConfig(appEnv)
	libdomain.CheckError(err)

	// init log
	rslibconfig.InitLog(cfg.Log)
	logger := slog.Default().With(slog.String(rsliblog.LoggerNameKey, "main"))
	logger.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	// init tracer
	tp, err := rslibconfig.InitTracerProvider(ctx, initialize.AppName, cfg.Trace)
	libdomain.CheckError(err)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// init db
	dialect, db, sqlDB, err := rslibconfig.InitDB(ctx, cfg.DB, sqls.SQL)
	libdomain.CheckError(err)
	defer sqlDB.Close()
	defer tp.ForceFlush(ctx) // flushes any pending spans

	// init gin
	router := libcontroller.InitRootRouterGroup(ctx, cfg.CORS, cfg.Debug)

	if err := initialize.Initialize(ctx, router, dialect, cfg.DB.DriverName, db, cfg.App); err != nil {
		libdomain.CheckError(err)
	}

	// run
	readHeaderTimeout := time.Duration(cfg.Server.ReadHeaderTimeoutSec) * time.Second
	shutdownTime := time.Duration(cfg.Shutdown.TimeSec1) * time.Second
	result := libgateway.Run(ctx,
		libgateway.WithAppServerProcess(router, cfg.Server.HTTPPort, readHeaderTimeout, shutdownTime),
		libgateway.WithSignalWatchProcess(),
		libgateway.WithMetricsServerProcess(cfg.Server.MetricsPort, cfg.Shutdown.TimeSec1),
	)

	gracefulShutdownTime2 := time.Duration(cfg.Shutdown.TimeSec2) * time.Second
	time.Sleep(gracefulShutdownTime2)
	logger.InfoContext(ctx, "exited")
	os.Exit(result)
}

// func initGin(ctx context.Context, cfg *config.Config, txManager, nonTxManager service.TransactionManager) http.Handler {
// 	if !cfg.Debug.Gin {
// 		gin.SetMode(gin.ReleaseMode)
// 	}
// 	router := gin.New()

// 	authMiddleware := controller.InitAuthMiddleware(cfg.App.InternalAuth)
// 	publicRouterGroupFuncs := controller.GetPublicRouterGroupFuncs()
// 	privateRouterGroupFuncs := controller.GetPrivateRouterGroupFuncs(cfg.TTS, txManager, nonTxManager)
// 	initialize.InitAppServer(ctx, router, cfg.CORS, cfg.Debug, cfg.App.Name, authMiddleware, publicRouterGroupFuncs, privateRouterGroupFuncs)
// 	return router
// }
