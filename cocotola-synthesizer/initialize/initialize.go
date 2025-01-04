package initialize

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/service"
)

const AppName = "cocotola-core"

func Initialize(ctx context.Context, parent gin.IRouter, dialect rslibgateway.DialectRDBMS, driverName string, db *gorm.DB, cfg *config.AppConfig) error {
	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		return gateway.NewRepositoryFactory(ctx, dialect, driverName, db, time.UTC) // nolint:wrapcheck
	}
	rf, err := rff(ctx, db)
	if err != nil {
		return err
	}

	// init transaction manager
	txManager, err := rslibgateway.NewTransactionManagerT(db, rff)
	if err != nil {
		return err
	}

	// init non transaction manager
	nonTxManager, err := rslibgateway.NewNonTransactionManagerT(rf)
	if err != nil {
		return err
	}

	authMiddleware := controller.InitAuthMiddleware(cfg.InternalAuth)

	// init public and private router group functions
	publicRouterGroupFuncs := controller.GetPublicRouterGroupFuncs()
	privateRouterGroupFuncs := controller.GetPrivateRouterGroupFuncs(cfg.TTS, txManager, nonTxManager)

	initApiServer(ctx, parent, AppName, authMiddleware, publicRouterGroupFuncs, privateRouterGroupFuncs)

	return nil
}

func initApiServer(ctx context.Context, rootRouterGroup gin.IRouter, appName string, authMiddleware gin.HandlerFunc, publicRouterGroupFuncs, privateRouterGroupFuncs []libcontroller.InitRouterGroupFunc) {
	// api
	api := libcontroller.InitAPIRouterGroup(ctx, rootRouterGroup, appName)

	// v1
	v1 := api.Group("v1")

	// public router
	libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs)

	// private router
	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, authMiddleware, privateRouterGroupFuncs)
}

// func initAppServer(ctx context.Context, rootRouterGroup gin.IRouter, corsConfig *rslibconfig.CORSConfig, debugConfig *libconfig.DebugConfig, appName string, authMiddleware gin.HandlerFunc, publicRouterGroupFuncs, privateRouterGroupFuncs []libcontroller.InitRouterGroupFunc) {
// 	// cors
// 	ginCorsConfig := rslibconfig.InitCORS(corsConfig)

// 	// root
// 	libcontroller.InitRootRouterGroup(ctx, rootRouterGroup, ginCorsConfig, debugConfig)

// 	// api
// 	api := libcontroller.InitAPIRouterGroup(ctx, rootRouterGroup, appName)

// 	// v1
// 	v1 := api.Group("v1")

// 	// public router
// 	libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs)

// 	// private router
// 	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, authMiddleware, privateRouterGroupFuncs)
// }

// const readHeaderTimeout = time.Duration(30) * time.Second
// const authClientTimeout = time.Duration(5) * time.Second

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
