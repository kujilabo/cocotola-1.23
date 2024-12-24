package initialize

import (
	"context"

	"github.com/gin-gonic/gin"

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"
	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"

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
	publicRouterGroupFunc := []libcontroller.InitRouterGroupFunc{
		controller.NewInitTestRouterFunc(),
	}

	// private router
	privateRouterGroupFunc := []libcontroller.InitRouterGroupFunc{
		controller.NewInitAdminRouterFunc(adminUsecase),
	}

	controller.InitRootRouterGroup(ctx, rootRouterGroup, ginCorsConfig, debugConfig)

	if err := controller.InitAPIRouterGroup(ctx, rootRouterGroup, authMiddleware, publicRouterGroupFunc, privateRouterGroupFunc, appName); err != nil {
		return err
	}

	return nil
}
