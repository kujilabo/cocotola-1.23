package initialize

import (
	"context"

	"github.com/gin-gonic/gin"

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"
	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"
)

// const readHeaderTimeout = time.Duration(30) * time.Second
// const authClientTimeout = time.Duration(5) * time.Second

// func InitTransactionManager(db *gorm.DB, rff gateway.RepositoryFactoryFunc) service.TransactionManager {
// 	appTransactionManager, err := gateway.NewTransactionManager(db, rff)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return appTransactionManager
// }

// type systemOwnerByOrganizationName struct {
// }

// func (s systemOwnerByOrganizationName) Get(ctx context.Context, rf service.RepositoryFactory, organizationName string) (*rsuserservice.SystemOwner, error) {
// 	rsrf, err := rf.NewRedstartRepositoryFactory(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	systemAdmin, err := rsuserservice.NewSystemAdmin(ctx, rsrf)
// 	if err != nil {
// 		return nil, err
// 	}

// 	systemOwner, err := systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return systemOwner, nil
// }

func InitAppServer(ctx context.Context, rootRouterGroup gin.IRouter, corsConfig *rslibconfig.CORSConfig, debugConfig *libconfig.DebugConfig, appName string, authMiddleware gin.HandlerFunc, publicRouterGroupFuncs, privateRouterGroupFuncs []libcontroller.InitRouterGroupFunc) error {
	// cors
	ginCorsConfig := rslibconfig.InitCORS(corsConfig)

	// root
	libcontroller.InitRootRouterGroup(ctx, rootRouterGroup, ginCorsConfig, debugConfig)

	// api
	api := libcontroller.InitAPIRouterGroup(ctx, rootRouterGroup, appName)

	// v1
	v1 := api.Group("v1")

	// public router
	if err := libcontroller.InitPrivateAPIRouterGroup(ctx, v1, authMiddleware, privateRouterGroupFuncs); err != nil {
		return err
	}

	// private router
	if err := libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs); err != nil {
		return err
	}
	return nil
}
