package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"gorm.io/gorm"

	rsmysqllibgateway "github.com/kujilabo/cocotola-1.23/redstart-mysql/lib/gateway"
	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libgateway "github.com/kujilabo/cocotola-1.23/lib/gateway"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/initialize"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/sqls"
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
		log.Fatalf("err: %+v", err)
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
	dialect, db, sqlDB, err := rslibconfig.InitDB(ctx, cfg.DB, map[string]rslibconfig.DBInitializer{
		"mysql": rsmysqllibgateway.InitMySQL,
	}, sqls.SQL)
	checkError(err)
	defer sqlDB.Close()
	defer tp.ForceFlush(ctx) // flushes any pending spans

	// init repository factory function
	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		return gateway.NewRepositoryFactory(ctx, dialect, cfg.DB.DriverName, db, time.UTC) // nolint:wrapcheck
	}

	// init repository factory
	rf, err := rff(ctx, db)
	checkError(err)

	// init transaction manager
	txManager, err := rslibgateway.NewTransactionManagerT(db, rff)
	checkError(err)

	// init non transaction manager
	nonTxManager, err := rslibgateway.NewNonTransactionManagerT(rf)
	checkError(err)

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
	router := gin.New()

	authMiddleware := controller.InitAuthMiddleware(cfg.InternalAuth)
	publicRouterGroupFuncs := controller.GetPublicRouterGroupFuncs()
	privateRouterGroupFuncs := controller.GetPrivateRouterGroupFuncs(txManager, nonTxManager)

	initialize.InitAppServer(ctx, router, cfg.CORS, cfg.Debug, cfg.App.Name, authMiddleware, publicRouterGroupFuncs, privateRouterGroupFuncs)
	return router
}
