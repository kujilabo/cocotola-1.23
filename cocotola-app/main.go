package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strings"
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

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"
	libgateway "github.com/kujilabo/cocotola-1.23/lib/gateway"

	authconfig "github.com/kujilabo/cocotola-1.23/cocotola-auth/config"
	authcontroller "github.com/kujilabo/cocotola-1.23/cocotola-auth/controller/gin"
	authgateway "github.com/kujilabo/cocotola-1.23/cocotola-auth/gateway"
	authinit "github.com/kujilabo/cocotola-1.23/cocotola-auth/initialize"
	authservice "github.com/kujilabo/cocotola-1.23/cocotola-auth/service"

	corecontroller "github.com/kujilabo/cocotola-1.23/cocotola-core/controller/gin"
	coregateway "github.com/kujilabo/cocotola-1.23/cocotola-core/gateway"
	coreinit "github.com/kujilabo/cocotola-1.23/cocotola-core/initialize"
	coreservice "github.com/kujilabo/cocotola-1.23/cocotola-core/service"

	"github.com/kujilabo/cocotola-1.23/cocotola-app/config"
)

//go:embed web_dist
var web embed.FS

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
	slog.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	rsliberrors.UseXerrorsErrorf()

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

	authRFF := func(ctx context.Context, db *gorm.DB) (authservice.RepositoryFactory, error) {
		return authgateway.NewRepositoryFactory(ctx, dialect, cfg.DB.DriverName, db, time.UTC) // nolint:wrapcheck
	}
	authRF, err := authRFF(ctx, db)
	checkError(err)

	coreRFF := func(ctx context.Context, db *gorm.DB) (coreservice.RepositoryFactory, error) {
		return coregateway.NewRepositoryFactory(ctx, dialect, cfg.DB.DriverName, db, time.UTC) // nolint:wrapcheck
	}
	coreRF, err := coreRFF(ctx, db)
	checkError(err)

	router := initGinRouter(ctx, cfg)

	// web
	{
		viteStaticFS, err := fs.Sub(web, "web_dist")
		checkError(err)
		initGinWeb(ctx, router, viteStaticFS)
	}

	// auth
	{
		authTxManager, err := rslibgateway.NewTransactionManagerT(db, authRFF)
		checkError(err)

		authNonTxManager, err := rslibgateway.NewNonTransactionManagerT(authRF)
		checkError(err)
		initGinAuth(ctx, router, cfg.Auth, authTxManager, authNonTxManager)
	}

	// core
	{
		coreTxManager, err := rslibgateway.NewTransactionManagerT(db, coreRFF)
		checkError(err)
		coreNonTxManager, err := rslibgateway.NewNonTransactionManagerT(coreRF)
		checkError(err)

		coreAuthMiddleware, err := corecontroller.InitAuthMiddleware(cfg.AuthAPI)
		checkError(err)
		initGinCore(ctx, router, coreAuthMiddleware, db, coreTxManager, coreNonTxManager)
	}

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

func initGinRouter(ctx context.Context, cfg *config.Config) *gin.Engine {
	if !cfg.Debug.Gin {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// cors
	ginCorsConfig := rslibconfig.InitCORS(cfg.CORS)

	// root
	libcontroller.InitRootRouterGroup(ctx, router, ginCorsConfig, cfg.Debug)

	return router
}

func initGinWeb(ctx context.Context, router *gin.Engine, viteStaticFS fs.FS) {
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/assets") {
			c.FileFromFS(c.Request.URL.Path, http.FS(viteStaticFS))
			return
		}
		if !strings.HasPrefix(c.Request.URL.Path, "/auth") {
			c.FileFromFS("", http.FS(viteStaticFS))
			return
		}
		if !strings.HasPrefix(c.Request.URL.Path, "/core") {
			c.FileFromFS("", http.FS(viteStaticFS))
			return
		}
		if !strings.HasPrefix(c.Request.URL.Path, "/synthesizer") {
			c.FileFromFS("", http.FS(viteStaticFS))
			return
		}
		if !strings.HasPrefix(c.Request.URL.Path, "/tatoeba") {
			c.FileFromFS("", http.FS(viteStaticFS))
			return
		}
	})

}

func initGinAuth(ctx context.Context, router gin.IRouter, authConfig *authconfig.AuthConfig, txManager, nonTxManager authservice.TransactionManager) {
	auth := router.Group("auth")
	publicRouterGroupFuncs := authcontroller.GetPublicRouterGroupFuncs(authConfig, txManager, nonTxManager)
	authinit.InitApiServer(ctx, auth, "auth", publicRouterGroupFuncs)
}

func initGinCore(ctx context.Context, router gin.IRouter, authMiddleware gin.HandlerFunc, db *gorm.DB, txManager, nonTxManager coreservice.TransactionManager) {
	core := router.Group("core")
	publicRouterGroupFuncs := corecontroller.GetPublicRouterGroupFuncs()
	privateRouterGroupFuncs := corecontroller.GetPrivateRouterGroupFuncs(db, txManager, nonTxManager)
	coreinit.InitApiServer(ctx, core, "core", authMiddleware, publicRouterGroupFuncs, privateRouterGroupFuncs)
}
