package usecase

import (
	// "context"
	// "fmt"

	"context"

	"github.com/golang-jwt/jwt"

	// rsuserdomain "github.com/kujilabo/redstart/user/domain"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/service"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"
	// "github.com/kujilabo/cocotola-1.21/cocotola-auth/domain"
	// liblog "github.com/kujilabo/cocotola-1.21/lib/log"
	// rsliblog "github.com/kujilabo/redstart/lib/log"
)

type AppUserClaims struct {
	LoginID          string `json:"loginId"`
	AppUserID        int    `json:"appUserId"`
	Username         string `json:"username"`
	OrganizationID   int    `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
	TokenType        string `json:"tokenType"`
	jwt.StandardClaims
}

type Authentication struct {
	transactionManager            service.TransactionManager
	authTokenManager              service.AuthTokenManager
	systemOwnerByOrganizationName SystemOwnerByOrganizationName
}

func NewAuthentication(transactionManager service.TransactionManager, authTokenManager service.AuthTokenManager, systemOwnerByOrganizationName SystemOwnerByOrganizationName) *Authentication {
	return &Authentication{
		transactionManager:            transactionManager,
		authTokenManager:              authTokenManager,
		systemOwnerByOrganizationName: systemOwnerByOrganizationName,
	}
}

func (u *Authentication) GetUserInfo(ctx context.Context, bearerToken string) (*rsuserdomain.AppUserModel, error) {
	// TODO: Check whether the token is registered in the Database

	appUserInfo, err := u.authTokenManager.GetUserInfo(ctx, bearerToken)
	if err != nil {
		return nil, err
	}

	var targetAppUserModel *rsuserdomain.AppUserModel

	if err := u.transactionManager.Do(ctx, func(rf service.RepositoryFactory) error {
		systemOwner, err := u.systemOwnerByOrganizationName.Get(ctx, rf, appUserInfo.OrganizationName)
		if err != nil {
			return err
		}

		appUser, err := systemOwner.FindAppUserByLoginID(ctx, appUserInfo.LoginID)
		if err != nil {
			return err
		}

		targetAppUserModel = appUser.AppUserModel
		return nil
	}); err != nil {
		return nil, rsliberrors.Errorf("RegisterAppUser. err: %w", err)
	}

	return targetAppUserModel, nil
}

func (u *Authentication) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	accessToken, err := u.authTokenManager.RefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	// TODO: Save the token to the database

	return accessToken, nil
}

// func (u *Authentication) Authenticate(ctx context.Context, bearerToken string) (*rsuserdomain.AppUserModel, error) {
// 	logger := rsliblog.GetLoggerFromContext(ctx, liblog.AppUsecaseLoggerContextKey)

// 	token, err := jwt.ParseWithClaims(bearerToken, &AppUserClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return u.signingKey, nil
// 	})
// 	if err != nil {
// 		logger.InfoContext(ctx, fmt.Sprintf("invalid token. err: %v", err))
// 		return nil, domain.ErrUnauthenticated
// 	}

// 	claims, ok := token.Claims.(*AppUserClaims)
// 	if !ok || !token.Valid {
// 		// logger.InfoContext(ctx, "invalid token")
// 		return nil, domain.ErrUnauthenticated
// 	}

// 	systemAdmin, err := rsuserservice.NewSystemAdmin(ctx, u.rf)
// 	if err != nil {
// 		return nil, err
// 	}

// 	organizationID, err := rsuserdomain.NewOrganizationID(claims.OrganizationID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	appUserID, err := rsuserdomain.NewAppUserID(claims.AppUserID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// # TODO Check whether the token is registered in the Database

// 	appUserRepo := u.rf.NewAppUserRepository(ctx)
// 	systemOwner, err := appUserRepo.FindSystemOwnerByOrganizationID(ctx, systemAdmin, organizationID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	appUser, err := systemOwner.FindAppUserByID(ctx, appUserID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return appUser.AppUserModel, nil
// }
