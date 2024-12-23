package controller

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"
	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"
	libmiddleware "github.com/kujilabo/cocotola-1.23/lib/controller/gin/middleware"
)

func NewInitTestRouterFunc() libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup *gin.RouterGroup, middleware ...gin.HandlerFunc) error {
		test := parentRouterGroup.Group("test")
		for _, m := range middleware {
			test.Use(m)
		}
		test.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		return nil
	}
}

func InitRootRouterGroup(ctx context.Context, rootRouterGroup gin.IRouter, corsConfig cors.Config, debugConfig *libconfig.DebugConfig) {
	rootRouterGroup.Use(cors.New(corsConfig))
	rootRouterGroup.Use(sloggin.New(slog.Default()))

	if debugConfig.Wait {
		rootRouterGroup.Use(libmiddleware.NewWaitMiddleware())
	}
}

func InitAPIRouterGroup(ctx context.Context, parentRouterGroup gin.IRouter, authMiddleware gin.HandlerFunc, initPublicRouterFunc []libcontroller.InitRouterGroupFunc, initPrivateRouterFunc []libcontroller.InitRouterGroupFunc, appName string) error {
	api := parentRouterGroup.Group("api")
	api.Use(otelgin.Middleware(appName))
	api.Use(libmiddleware.NewTraceLogMiddleware(appName))

	v1 := api.Group("v1")
	{
		v1.Use(otelgin.Middleware(appName))
		v1.Use(libmiddleware.NewTraceLogMiddleware(appName))

		for _, fn := range initPublicRouterFunc {
			if err := fn(v1); err != nil {
				return err
			}
		}
		for _, fn := range initPrivateRouterFunc {
			if err := fn(v1, authMiddleware); err != nil {
				return err
			}
		}
	}

	return nil
}
