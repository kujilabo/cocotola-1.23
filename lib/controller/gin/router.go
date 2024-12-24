package controller

import (
	"context"
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/kujilabo/cocotola-1.23/lib/config"
	"github.com/kujilabo/cocotola-1.23/lib/controller/gin/middleware"
)

type InitRouterGroupFunc func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc)

func InitRootRouterGroup(ctx context.Context, rootRouterGroup gin.IRouter, corsConfig cors.Config, debugConfig *config.DebugConfig) {
	rootRouterGroup.Use(cors.New(corsConfig))
	rootRouterGroup.Use(sloggin.New(slog.Default()))

	if debugConfig.Wait {
		rootRouterGroup.Use(middleware.NewWaitMiddleware())
	}
}

func InitAPIRouterGroup(ctx context.Context, parentRouterGroup gin.IRouter, appName string) *gin.RouterGroup {
	api := parentRouterGroup.Group("api")
	api.Use(otelgin.Middleware(appName))
	api.Use(middleware.NewTraceLogMiddleware(appName))
	return api
}

func InitPublicAPIRouterGroup(ctx context.Context, parentRouterGroup gin.IRouter, initPublicRouterFunc []InitRouterGroupFunc) {
	for _, fn := range initPublicRouterFunc {
		fn(parentRouterGroup)
	}
}

func InitPrivateAPIRouterGroup(ctx context.Context, parentRouterGroup gin.IRouter, authMiddleware gin.HandlerFunc, initPrivateRouterFunc []InitRouterGroupFunc) {
	for _, fn := range initPrivateRouterFunc {
		fn(parentRouterGroup, authMiddleware)
	}
}
