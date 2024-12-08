package initialize

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"gorm.io/gorm"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"
	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"

	"github.com/kujilabo/cocotola-1.23/cocotola-core/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-core/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/controller/gin/middleware"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/gateway"
	studentusecasegateway "github.com/kujilabo/cocotola-1.23/cocotola-core/gateway/usecase/student"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/service"
	studentusecase "github.com/kujilabo/cocotola-1.23/cocotola-core/usecase/student"
)

// const readHeaderTimeout = time.Duration(30) * time.Second
const authClientTimeout = time.Duration(5) * time.Second

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

func InitAppServer(ctx context.Context, parentRouterGroup gin.IRouter, authAPIConfig config.AuthAPIonfig, corsConfig *rslibconfig.CORSConfig, debugConfig *libconfig.DebugConfig, appName string, db *gorm.DB, txManager, nonTxManager service.TransactionManager) error {
	// cors
	gincorsConfig := rslibconfig.InitCORS(corsConfig)
	httpClient := http.Client{
		Timeout:   authClientTimeout,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	workbookQuerySerivce := studentusecasegateway.NewWorkbookQueryService(db)
	workbookQueryUsecase := studentusecase.NewWorkbookQueryUsecase(txManager, nonTxManager, workbookQuerySerivce)
	workbookCommandUsecase := studentusecase.NewWorkbookCommandUsecase(txManager, nonTxManager)
	privateRouterGroupFunc := []controller.InitRouterGroupFunc{
		controller.NewInitWorkbookRouterFunc(workbookQueryUsecase, workbookCommandUsecase),
	}
	authEndpoint, err := url.Parse(authAPIConfig.Endpoint)
	if err != nil {
		return err
	}

	publicRouterGroupFunc := []controller.InitRouterGroupFunc{
		controller.NewInitTestRouterFunc(),
	}
	cocotolaAuthClient := gateway.NewCocotolaAuthClient(&httpClient, authEndpoint, authAPIConfig.Username, authAPIConfig.Password)
	authMiddleware := middleware.NewAuthMiddleware(cocotolaAuthClient)

	if err := controller.InitRouter(ctx, parentRouterGroup, authMiddleware, publicRouterGroupFunc, privateRouterGroupFunc, gincorsConfig, debugConfig, appName); err != nil {
		return err
	}

	return nil
}

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
