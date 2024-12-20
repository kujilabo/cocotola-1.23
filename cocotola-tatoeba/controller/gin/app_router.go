package handler

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"
	libmiddleware "github.com/kujilabo/cocotola-1.23/lib/controller/gin/middleware"
)

type InitRouterGroupFunc func(parentRouterGroup *gin.RouterGroup, middleware ...gin.HandlerFunc) error

func NewInitTestRouterFunc() InitRouterGroupFunc {
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
	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	rootRouterGroup.Use(cors.New(corsConfig))
	rootRouterGroup.Use(sloggin.New(logger))

	if debugConfig.Wait {
		rootRouterGroup.Use(libmiddleware.NewWaitMiddleware())
	}
}

func InitAPIRouterGroup(ctx context.Context, apiRouterGroup gin.IRouter, authMiddleware gin.HandlerFunc, initPublicRouterFunc []InitRouterGroupFunc, initPrivateRouterFunc []InitRouterGroupFunc, appName string) error {
	v1 := apiRouterGroup.Group("v1")
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
