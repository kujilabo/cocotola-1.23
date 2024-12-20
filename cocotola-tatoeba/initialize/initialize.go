package initialize

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"
	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rsuserservice "github.com/kujilabo/cocotola-1.23/redstart/user/service"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/usecase"
)

// const readHeaderTimeout = time.Duration(30) * time.Second
// const authClientTimeout = time.Duration(5) * time.Second

func InitTransactionManager(db *gorm.DB, rff gateway.RepositoryFactoryFunc) service.TransactionManager {
	appTransactionManager, err := gateway.NewTransactionManager(db, rff)
	if err != nil {
		panic(err)
	}

	return appTransactionManager
}

func InitAppServer(ctx context.Context, rootRouterGroup gin.IRouter, internalAuthConfig config.InternalAuthConfig, corsConfig *rslibconfig.CORSConfig, debugConfig *libconfig.DebugConfig, appName string, txManager, nonTxManager service.TransactionManager, rsrf rsuserservice.RepositoryFactory) error {
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
