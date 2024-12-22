package initialize

import (
	"context"

	"github.com/gin-gonic/gin"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"
	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/usecase"
)

func InitAppServer(ctx context.Context, rootRouterGroup gin.IRouter, internalAuthConfig config.InternalAuthConfig, corsConfig *rslibconfig.CORSConfig, debugConfig *libconfig.DebugConfig, appName string, txManager, nonTxManager service.TransactionManager) error {
	// cors
	ginCorsConfig := rslibconfig.InitCORS(corsConfig)

	// usecase
	adminUsecase := usecase.NewAdminUsecase(txManager, nonTxManager)

	// middleware
	authMiddleware := gin.BasicAuth(gin.Accounts{
		internalAuthConfig.Username: internalAuthConfig.Password,
	})

	// public router
	privateRouterGroupFunc := []controller.InitRouterGroupFunc{
		controller.NewInitAdminRouterFunc(adminUsecase),
	}

	// private router
	publicRouterGroupFunc := []controller.InitRouterGroupFunc{
		controller.NewInitTestRouterFunc(),
	}

	controller.InitRootRouterGroup(ctx, rootRouterGroup, ginCorsConfig, debugConfig)

	if err := controller.InitAPIRouterGroup(ctx, rootRouterGroup, authMiddleware, publicRouterGroupFunc, privateRouterGroupFunc, appName); err != nil {
		return err
	}

	return nil
}
