package initialize

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/service"
	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/usecase"
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

func InitAppServer(ctx context.Context, rootRouterGroup gin.IRouter, internalAuthConfig config.InternalAuthConfig, corsConfig *rslibconfig.CORSConfig, debugConfig *libconfig.DebugConfig, ttsConfig *config.GoogleTextToSpeechConfig, appName string, txManager, nonTxManager service.TransactionManager) error {
	// cors
	ginCorsConfig := rslibconfig.InitCORS(corsConfig)

	// usecase
	httpClient := http.Client{
		Timeout:   time.Duration(ttsConfig.APITimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	synthesizerClient := gateway.NewGoogleTTSClient(&httpClient, ttsConfig.APIKey)
	audioFile := gateway.NewAudioFile()
	synthesizerUsecase := usecase.NewSynthesizerUsecase(txManager, nonTxManager, synthesizerClient, audioFile)

	// middleware
	authMiddleware := gin.BasicAuth(gin.Accounts{
		internalAuthConfig.Username: internalAuthConfig.Password,
	})

	// public router
	publicRouterGroupFunc := []controller.InitRouterGroupFunc{
		controller.NewInitTestRouterFunc(),
	}

	// private router
	privateRouterGroupFunc := []controller.InitRouterGroupFunc{
		controller.NewInitSynthesizerRouterFunc(synthesizerUsecase),
	}

	// rout
	controller.InitRootRouterGroup(ctx, rootRouterGroup, ginCorsConfig, debugConfig)

	// api
	if err := controller.InitAPIRouterGroup(ctx, rootRouterGroup, authMiddleware, publicRouterGroupFunc, privateRouterGroupFunc, appName); err != nil {
		return err
	}

	return nil
}
