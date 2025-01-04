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

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	"github.com/kujilabo/cocotola-1.23/redstart/sqls"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"
	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"
	libgateway "github.com/kujilabo/cocotola-1.23/lib/gateway"

	authinit "github.com/kujilabo/cocotola-1.23/cocotola-auth/initialize"

	synthesizerinit "github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/initialize"

	coreinit "github.com/kujilabo/cocotola-1.23/cocotola-core/initialize"
	tatoebainit "github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/initialize"

	"github.com/kujilabo/cocotola-1.23/cocotola-app/config"
)

//go:embed web_dist
var web embed.FS

const AppName = "cocotola-app"

func main() {
	ctx := context.Background()
	env := flag.String("env", "", "environment")
	flag.Parse()
	appEnv := libdomain.GetNonEmptyValue(*env, os.Getenv("APP_ENV"), "local")
	slog.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	rsliberrors.UseXerrorsErrorf()

	cfg, err := config.LoadConfig(appEnv)
	libdomain.CheckError(err)

	// init log
	rslibconfig.InitLog(cfg.Log)
	logger := slog.Default().With(slog.String(rsliblog.LoggerNameKey, "main"))
	logger.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	// init tracer
	tp, err := rslibconfig.InitTracerProvider(ctx, AppName, cfg.Trace)
	libdomain.CheckError(err)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// init db
	dialect, db, sqlDB, err := rslibconfig.InitDB(ctx, cfg.DB, sqls.SQL)
	libdomain.CheckError(err)
	defer sqlDB.Close()
	defer tp.ForceFlush(ctx) // flushes any pending spans

	router := libcontroller.InitRootRouterGroup(ctx, cfg.CORS, cfg.Debug)

	// web
	{
		viteStaticFS, err := fs.Sub(web, "web_dist")
		libdomain.CheckError(err)
		initGinWeb(ctx, router, viteStaticFS)
	}
	// auth
	{
		auth := router.Group("auth")
		authinit.Initialize(ctx, auth, dialect, cfg.DB.DriverName, db, cfg.Auth)
	}
	// core
	{
		core := router.Group("core")
		coreinit.Initialize(ctx, core, dialect, cfg.DB.DriverName, db, cfg.Core)
	}
	// synthesizer
	{
		synthesizer := router.Group("synthesizer")
		synthesizerinit.Initialize(ctx, synthesizer, dialect, cfg.DB.DriverName, db, cfg.Synthesizer)
	}
	// tatoeba
	{
		tatoeba := router.Group("tatoeba")
		tatoebainit.Initialize(ctx, tatoeba, dialect, cfg.DB.DriverName, db, cfg.Tatoeba)
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

// func initGinRouter(ctx context.Context, cfg *config.Config) *gin.Engine {
// 	if !cfg.Debug.Gin {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	router := gin.New()

// 	// cors
// 	ginCorsConfig := rslibconfig.InitCORS(cfg.CORS)

// 	// root
// 	libcontroller.InitRootRouterGroup(ctx, router, ginCorsConfig, cfg.Debug)

// 	return router
// }

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
