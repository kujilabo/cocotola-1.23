package initialize

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	rsuserservice "github.com/kujilabo/cocotola-1.23/redstart/user/service"

	libcontroller "github.com/kujilabo/cocotola-1.23/lib/controller/gin"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/config"
	controller "github.com/kujilabo/cocotola-1.23/cocotola-auth/controller/gin"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/gateway"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/service"
)

const AppName = "cocotola-auth"

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

	// init public and private router group functions
	publicRouterGroupFuncs, err := controller.GetPublicRouterGroupFuncs(ctx, cfg.Auth, txManager, nonTxManager)
	if err != nil {
		return err
	}

	initApiServer(ctx, parent, AppName, publicRouterGroupFuncs)

	initApp1(ctx, txManager, nonTxManager, "cocotola", cfg.OwnerLoginID, cfg.OwnerPassword)

	return nil
}

func initApiServer(ctx context.Context, root gin.IRouter, appName string, publicRouterGroupFuncs []libcontroller.InitRouterGroupFunc) {
	// api
	api := libcontroller.InitAPIRouterGroup(ctx, root, appName)

	// v1
	v1 := api.Group("v1")

	// public router
	libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs)
}

func initApp1(ctx context.Context, txManager, nonTxManager service.TransactionManager, organizationName, loginID, password string) {
	logger := slog.Default().With(slog.String(rsliblog.LoggerNameKey, "InitApp1"))

	addOrganizationFunc := func(ctx context.Context, systemAdmin *rsuserservice.SystemAdmin) error {
		organization, err := systemAdmin.FindOrganizationByName(ctx, organizationName)
		if err == nil {
			logger.InfoContext(ctx, fmt.Sprintf("organization: %d", organization.OrganizationID().Int()))
			return nil
		} else if !errors.Is(err, rsuserservice.ErrOrganizationNotFound) {
			return rsliberrors.Errorf("failed to FindOrganizationByName. err: %w", err)
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

// func InitAppServer(ctx context.Context, rootRouterGroup gin.IRouter, corsConfig *rslibconfig.CORSConfig, debugConfig *libconfig.DebugConfig, appName string, publicRouterGroupFuncs []libcontroller.InitRouterGroupFunc) {
// 	// cors
// 	ginCorsConfig := rslibconfig.InitCORS(corsConfig)

// 	// root
// 	libcontroller.InitRootRouterGroup(ctx, rootRouterGroup, ginCorsConfig, debugConfig)

// 	InitApiServer(ctx, rootRouterGroup, appName, publicRouterGroupFuncs)
// }
