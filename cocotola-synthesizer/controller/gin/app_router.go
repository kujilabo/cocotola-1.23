package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/config"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/service"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/usecase"
	libhandler "github.com/kujilabo/cocotola-1.23/lib/controller/gin"
)

func NewInitTestRouterFunc() libhandler.InitRouterGroupFunc {
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

func GetPublicRouterGroupFuncs() []libhandler.InitRouterGroupFunc {
	// public router
	return []libhandler.InitRouterGroupFunc{
		NewInitTestRouterFunc(),
	}
}

func GetPrivateRouterGroupFuncs(ttsConfig *config.GoogleTextToSpeechConfig, txManager, nonTxManager service.TransactionManager) []libhandler.InitRouterGroupFunc {
	// usecase
	httpClient := http.Client{
		Timeout:   time.Duration(ttsConfig.APITimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	synthesizerClient := gateway.NewGoogleTTSClient(&httpClient, ttsConfig.APIKey)
	audioFile := gateway.NewAudioFile()
	synthesizerUsecase := usecase.NewSynthesizerUsecase(txManager, nonTxManager, synthesizerClient, audioFile)

	// private router
	return []libhandler.InitRouterGroupFunc{
		NewInitSynthesizerRouterFunc(synthesizerUsecase),
	}
}

func InitAuthMiddleware(internalAuthConfig *config.InternalAuthConfig) gin.HandlerFunc {
	// middleware
	return gin.BasicAuth(gin.Accounts{
		internalAuthConfig.Username: internalAuthConfig.Password,
	})
}
