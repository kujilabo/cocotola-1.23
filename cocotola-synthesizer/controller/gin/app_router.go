package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/config"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/service"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/usecase"
)

func NewInitTestRouterFunc() libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		test := parentRouterGroup.Group("test")
		for _, m := range middleware {
			test.Use(m)
		}
		test.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}
}

func GetPublicRouterGroupFuncs() []libcontroller.InitRouterGroupFunc {
	// public router
	return []libcontroller.InitRouterGroupFunc{
		NewInitTestRouterFunc(),
	}
}

func GetPrivateRouterGroupFuncs(ttsConfig *config.GoogleTextToSpeechConfig, txManager, nonTxManager service.TransactionManager) []libcontroller.InitRouterGroupFunc {
	// usecase
	httpClient := http.Client{
		Timeout:   time.Duration(ttsConfig.APITimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	synthesizerClient := gateway.NewGoogleTTSClient(&httpClient, ttsConfig.APIKey)
	audioFile := gateway.NewAudioFile()
	synthesizerUsecase := usecase.NewSynthesizerUsecase(txManager, nonTxManager, synthesizerClient, audioFile)

	// private router
	return []libcontroller.InitRouterGroupFunc{
		NewInitSynthesizerRouterFunc(synthesizerUsecase),
	}
}

func InitAuthMiddleware(internalAuthConfig *config.InternalAuthConfig) gin.HandlerFunc {
	// middleware
	return gin.BasicAuth(gin.Accounts{
		internalAuthConfig.Username: internalAuthConfig.Password,
	})
}
