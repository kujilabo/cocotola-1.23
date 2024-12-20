package initialize

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	rsuserservice "github.com/kujilabo/cocotola-1.23/redstart/user/service"

	libconfig "github.com/kujilabo/cocotola-1.23/lib/config"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-auth/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/service"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/usecase"
)

// func InitTransactionManager(db *gorm.DB, rff gateway.RepositoryFactoryFunc) service.TransactionManager {
// 	appTransactionManager, err := gateway.NewTransactionManager(db, rff)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return appTransactionManager
// }

type systemOwnerByOrganizationName struct {
}

func (s systemOwnerByOrganizationName) Get(ctx context.Context, rf service.RepositoryFactory, organizationName string) (*rsuserservice.SystemOwner, error) {
	rsrf, err := rf.NewRedstartRepositoryFactory(ctx)
	if err != nil {
		return nil, err
	}
	systemAdmin, err := rsuserservice.NewSystemAdmin(ctx, rsrf)
	if err != nil {
		return nil, err
	}

	systemOwner, err := systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
	if err != nil {
		return nil, err
	}

	return systemOwner, nil
}

func InitAppServer(ctx context.Context, parentRouterGroup gin.IRouter, corsConfig *rslibconfig.CORSConfig, authConfig *config.AuthConfig, debugConfig *libconfig.DebugConfig, appName string, txManager, nonTxManager service.TransactionManager, rsrf rsuserservice.RepositoryFactory) error {
	// cors
	gincorsConfig := rslibconfig.InitCORS(corsConfig)
	httpClient := http.Client{
		Timeout:   time.Duration(authConfig.APITimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	googleAuthClient := gateway.NewGoogleAuthClient(&httpClient, authConfig.GoogleClientID, authConfig.GoogleClientSecret, authConfig.GoogleCallbackURL)

	signingKey := []byte(authConfig.SigningKey)
	signingMethod := jwt.SigningMethodHS256
	authTokenManager := gateway.NewAuthTokenManager(signingKey, signingMethod, time.Duration(authConfig.AccessTokenTTLMin)*time.Minute, time.Duration(authConfig.RefreshTokenTTLHour)*time.Hour)

	rbacUsecase := usecase.NewRBAC(txManager, nonTxManager)
	authenticationUsecase := usecase.NewAuthentication(txManager, authTokenManager, &systemOwnerByOrganizationName{})
	googleUserUsecase := usecase.NewGoogleUser(txManager, nonTxManager, authTokenManager, googleAuthClient)
	passwordUsecase := usecase.NewPassword(txManager, nonTxManager, authTokenManager)

	privateRouterGroupFunc := []controller.InitRouterGroupFunc{
		controller.NewInitRBACRouterFunc(rbacUsecase),
	}

	publicRouterGroupFunc := []controller.InitRouterGroupFunc{
		controller.NewInitTestRouterFunc(),
		controller.NewInitAuthRouterFunc(authenticationUsecase),
		controller.NewInitGoogleRouterFunc(googleUserUsecase),
		controller.NewInitPasswordRouterFunc(passwordUsecase),
	}

	if err := controller.InitRouter(ctx, parentRouterGroup, publicRouterGroupFunc, privateRouterGroupFunc, gincorsConfig, debugConfig, appName); err != nil {
		return err
	}

	return nil
}

func InitApp1(ctx context.Context, txManager, nonTxManager service.TransactionManager, organizationName, loginID, password string) {
	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

	addOrganizationFunc := func(ctx context.Context, systemAdmin *rsuserservice.SystemAdmin) error {
		organization, err := systemAdmin.FindOrganizationByName(ctx, organizationName)
		if err == nil {
			logger.InfoContext(ctx, fmt.Sprintf("organization: %d", organization.OrganizationID().Int()))
			return nil
		} else if !errors.Is(err, rsuserservice.ErrOrganizationNotFound) {
			return rsliberrors.Errorf("failed to AddOrganization. err: %w", err)
		}

		firstOwnerAddParam, err := rsuserservice.NewAppUserAddParameter(loginID, "Owner(cocotola)", password, "", "", "", "")
		if err != nil {
			return rsliberrors.Errorf("NewFirstOwnerAddParameter. err: %w", err)
		}

		organizationAddParameter, err := rsuserservice.NewOrganizationAddParameter(organizationName, firstOwnerAddParam)
		if err != nil {
			return rsliberrors.Errorf("NewOrganizationAddParameter. err: %w", err)
		}

		organizationID, err := systemAdmin.AddOrganization(ctx, organizationAddParameter)
		if err != nil {
			return rsliberrors.Errorf("AddOrganization. err: %w", err)
		}

		logger.InfoContext(ctx, fmt.Sprintf("organizationID: %d", organizationID.Int()))
		return nil
	}

	if err := systemAdminAction(ctx, txManager, addOrganizationFunc); err != nil {
		panic(err)
	}
}

func systemAdminAction(ctx context.Context, transactionManager service.TransactionManager, fn func(context.Context, *rsuserservice.SystemAdmin) error) error {
	return transactionManager.Do(ctx, func(rf service.RepositoryFactory) error {
		rsrf, err := rf.NewRedstartRepositoryFactory(ctx)
		if err != nil {
			return rsliberrors.Errorf(". err: %w", err)
		}

		systemAdmin, err := rsuserservice.NewSystemAdmin(ctx, rsrf)
		if err != nil {
			return rsliberrors.Errorf(". err: %w", err)
		}

		return fn(ctx, systemAdmin)
	})
}
