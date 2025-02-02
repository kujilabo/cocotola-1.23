package initialize

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"

	"github.com/kujilabo/cocotola-1.23/cocotola-core/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-core/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/service"
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

	authMiddleware, err := controller.InitAuthMiddleware(cfg.AuthAPI)
	if err != nil {
		return err
	}

	// init public and private router group functions
	publicRouterGroupFuncs := controller.GetPublicRouterGroupFuncs()
	privateRouterGroupFuncs := controller.GetPrivateRouterGroupFuncs(db, txManager, nonTxManager)

	initApiServer(ctx, parent, AppName, authMiddleware, publicRouterGroupFuncs, privateRouterGroupFuncs)

	return nil
}

func initApiServer(ctx context.Context, root gin.IRouter, appName string, authMiddleware gin.HandlerFunc, publicRouterGroupFuncs, privateRouterGroupFuncs []libcontroller.InitRouterGroupFunc) {
	// api
	api := libcontroller.InitAPIRouterGroup(ctx, root, appName)

	// v1
	v1 := api.Group("v1")

	// public router
	libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs)

	// private router
	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, authMiddleware, privateRouterGroupFuncs)
}

// const readHeaderTimeout = time.Duration(30) * time.Second

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

// func InitAppServer(ctx context.Context, rootRouterGroup gin.IRouter, corsConfig *rslibconfig.CORSConfig, debugConfig *libconfig.DebugConfig, appName string, authMiddleware gin.HandlerFunc, publicRouterGroupFuncs, privateRouterGroupFuncs []libcontroller.InitRouterGroupFunc) {
// 	// cors
// 	ginCorsConfig := rslibconfig.InitCORS(corsConfig)

// 	// root
// 	libcontroller.InitRootRouterGroup(ctx, rootRouterGroup, ginCorsConfig, debugConfig)

// 	InitApiServer(ctx, rootRouterGroup, appName, authMiddleware, publicRouterGroupFuncs, privateRouterGroupFuncs)
// }

// func InitApp1(ctx context.Context, txManager service.TransactionManager, workbookQueryService studentusecase.WorkbookQueryService) error {
// 	if err := txManager.Do(ctx, func(rf service.RepositoryFactory) error {
// 		// rf.NewWorkbookRepository(ctx)
// 		return nil
// 	}); err != nil {
// 		return err
// 	}

// 	workbookQueryService.RetrieveWorkbookByID(ctx)

// 	type Problem struct {
// 		Type       string            `json:"type"`
// 		Properties map[string]string `json:"properties"`
// 	}

// 	type Content struct {
// 		Problems []*Problem `json:"problems"`
// 	}

// 	x := Content{
// 		Problems: []*Problem{
// 			{
// 				Type: "text",
// 				Properties: map[string]string{
// 					"srcLang":         "ja",
// 					"srcAudioContent": audioContentJa1,
// 					"srcAudioLength":  strconv.Itoa(audioLengthJa1),
// 					"srcText":         "こんにちは",
// 					"dstLang":         "en",
// 					"dstAudioContent": audioContentEn1,
// 					"dstAudioLength":  strconv.Itoa(audioLengthEn1),
// 					"dstText":         "Hello",
// 				},
// 			},
// 			{
// 				Type: "text",
// 				Properties: map[string]string{
// 					"srcLang":         "ja",
// 					"srcAudioContent": audioContentJa2,
// 					"srcAudioLength":  strconv.Itoa(audioLengthJa2),
// 					"srcText":         "さようなら",
// 					"dstLang":         "en",
// 					"dstAudioContent": audioContentEn2,
// 					"dstAudioLength":  strconv.Itoa(audioLengthEn2),
// 					"dstText":         "Goodbye",
// 				},
// 			},
// 		},
// 	}

// 	_, err := json.Marshal(x)
// 	if err != nil {
// 		return err
// 	}

// 	// fmt.Println(jsonBytes)

// 	return nil
// }

// func systemOwnerAction(ctx context.Context, organizationName string, txManager service.TransactionManager, fn func(context.Context, *rsuserservice.SystemOwner) error) error {
// 	return txManager.Do(ctx, func(rf service.RepositoryFactory) error {
// 		rsrf, err := rf.NewRedstartRepositoryFactory(ctx)
// 		if err != nil {
// 			return rsliberrors.Errorf(". err: %w", err)
// 		}

// 		systemAdmin, err := rsuserservice.NewSystemAdmin(ctx, rsrf)
// 		if err != nil {
// 			return rsliberrors.Errorf(". err: %w", err)
// 		}
// 		systemOwner, err := systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
// 		if err != nil {
// 			return rsliberrors.Errorf(". err: %w", err)
// 		}

// 		return fn(ctx, systemOwner)
// 	})
// }
