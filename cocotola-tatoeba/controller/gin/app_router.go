package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/config"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/usecase"
)

func NewInitTestRouterFunc() libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		test := parentRouterGroup.Group("test")
		for _, m := range middleware {
			test.Use(m)
		}
		test.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}
}

func GetPublicRouterGroupFuncs() []libcontroller.InitRouterGroupFunc {
	// public router
	return []libcontroller.InitRouterGroupFunc{
		NewInitTestRouterFunc(),
	}
}

func GetPrivateRouterGroupFuncs(txManager, nonTxManager service.TransactionManager) []libcontroller.InitRouterGroupFunc {
	// usecase
	adminUsecase := usecase.NewAdminUsecase(txManager, nonTxManager)

	// private router
	return []libcontroller.InitRouterGroupFunc{
		NewInitAdminRouterFunc(adminUsecase),
	}
}

func InitAuthMiddleware(internalAuthConfig *config.InternalAuthConfig) gin.HandlerFunc {
	// middleware
	return gin.BasicAuth(gin.Accounts{
		internalAuthConfig.Username: internalAuthConfig.Password,
	})
}
