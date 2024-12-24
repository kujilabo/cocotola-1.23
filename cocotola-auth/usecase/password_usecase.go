package usecase

import (
	"context"
	"errors"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/domain"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/service"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rslibservice "github.com/kujilabo/cocotola-1.23/redstart/lib/service"
	rsuserservice "github.com/kujilabo/cocotola-1.23/redstart/user/service"
)

type PasswordUsecae struct {
	txManager        service.TransactionManager
	nonTxManager     service.TransactionManager
	authTokenManager service.AuthTokenManager
}

func NewPassword(txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager) *PasswordUsecae {
	return &PasswordUsecae{
		txManager:        txManager,
		nonTxManager:     nonTxManager,
		authTokenManager: authTokenManager,
	}
}

func (u *PasswordUsecae) Authenticate(ctx context.Context, loginID, password, organizationName string) (*domain.AuthTokenSet, error) {
	var tokenSet *domain.AuthTokenSet

	targetOorganization, targetAppUser, err := rslibservice.Do2(ctx, u.txManager, func(rf service.RepositoryFactory) (*organization, *appUser, error) {
		action, err := NewOrganizationAction(ctx, rf,
			WithOrganizationRepository(),
			WithOrganization(organizationName),
			WithAppUserRepository(),
		)
		if err != nil {
			return nil, nil, err
		}

		verified, err := action.appUserRepo.VerifyPassword(ctx, action.systemAdmin, action.organization.OrganizationModel.OrganizationID, loginID, password)
		if err != nil {
			return nil, nil, err
		}

		if !verified {
			return nil, nil, domain.ErrUnauthenticated
		}

		tmpAppUser, err := action.appUserRepo.FindAppUserByLoginID(ctx, action.systemOwner, loginID)
		if err != nil {
			return nil, nil, err
		}

		targetOorganization := &organization{
			organizationID: action.organization.OrganizationModel.OrganizationID,
			name:           action.organization.OrganizationModel.Name,
		}

		targetAppUser := &appUser{
			appUserID:      tmpAppUser.AppUserModel.AppUserID,
			organizationID: tmpAppUser.AppUserModel.OrganizationID,
			loginID:        tmpAppUser.AppUserModel.LoginID,
			username:       tmpAppUser.AppUserModel.Username,
		}

		return targetOorganization, targetAppUser, nil
	})

	if err != nil {
		if errors.Is(err, rsuserservice.ErrAppUserNotFound) {
			return nil, rsliberrors.Errorf("AppUserNotFound. err: %w", domain.ErrUnauthenticated)
		}
		return nil, rsliberrors.Errorf("RegisterAppUser. err: %w", err)
	}

	tokenSetTmp, err := u.authTokenManager.CreateTokenSet(ctx, targetAppUser, targetOorganization)
	if err != nil {
		return nil, rsliberrors.Errorf("s.authTokenManager.CreateTokenSet. err: %w", err)
	}
	tokenSet = tokenSetTmp
	return tokenSet, nil
}
