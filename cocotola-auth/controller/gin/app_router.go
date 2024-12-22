package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"
	libmiddleware "github.com/kujilabo/cocotola-1.23/lib/controller/gin/middleware"
)

type InitRouterGroupFunc func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) error

func NewInitTestRouterFunc() InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) error {
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

func InitRouter(ctx context.Context, parentRouterGroup gin.IRouter, initPublicRouterFunc []InitRouterGroupFunc, initPrivateRouterFunc []InitRouterGroupFunc, corsConfig cors.Config, debugConfig *libconfig.DebugConfig, appName string) error {
	parentRouterGroup.Use(cors.New(corsConfig))
	parentRouterGroup.Use(sloggin.New(slog.Default()))

	if debugConfig.Wait {
		parentRouterGroup.Use(libmiddleware.NewWaitMiddleware())
	}

	v1 := parentRouterGroup.Group("v1")
	{
		v1.Use(otelgin.Middleware(appName))
		v1.Use(libmiddleware.NewTraceLogMiddleware(appName))

		for _, fn := range initPublicRouterFunc {
			if err := fn(v1); err != nil {
				return err
			}
		}
	}

	return nil
}
