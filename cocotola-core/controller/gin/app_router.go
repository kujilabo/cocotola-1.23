package controller

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"gorm.io/gorm"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"

	"github.com/kujilabo/cocotola-1.23/cocotola-core/config"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/controller/gin/middleware"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/gateway"
	studentusecasegateway "github.com/kujilabo/cocotola-1.23/cocotola-core/gateway/usecase/student"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/service"
	studentusecase "github.com/kujilabo/cocotola-1.23/cocotola-core/usecase/student"
)

const authClientTimeout = time.Duration(5) * time.Second

// type NewIteratorFunc func(ctx context.Context, workbookID appD.WorkbookID, problemType appD.ProblemTypeName, reader io.Reader) (appS.ProblemAddParameterIterator, error)

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
		// controller.NewInitTestRouterFunc(),
	}
}

func GetPrivateRouterGroupFuncs(db *gorm.DB, txManager, nonTxManager service.TransactionManager) []libcontroller.InitRouterGroupFunc {
	// - workbookQueryUsecase
	workbookQuerySerivce := studentusecasegateway.NewWorkbookQueryService(db)
	workbookQueryUsecase := studentusecase.NewWorkbookQueryUsecase(txManager, nonTxManager, workbookQuerySerivce)
	// - workbookCommandUsecase
	workbookCommandUsecase := studentusecase.NewWorkbookCommandUsecase(txManager, nonTxManager)

	// private router
	return []libcontroller.InitRouterGroupFunc{
		NewInitWorkbookRouterFunc(workbookQueryUsecase, workbookCommandUsecase),
	}
}

func InitAuthMiddleware(authAPIConfig *config.AuthAPIonfig) (gin.HandlerFunc, error) {
	// middleware
	httpClient := http.Client{
		Timeout:   authClientTimeout,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	authEndpoint, err := url.Parse(authAPIConfig.Endpoint)
	if err != nil {
		return nil, err
	}
	cocotolaAuthClient := gateway.NewCocotolaAuthClient(&httpClient, authEndpoint, authAPIConfig.Username, authAPIConfig.Password)
	return middleware.NewAuthMiddleware(cocotolaAuthClient), nil
}
