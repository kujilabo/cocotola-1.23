package usecase

import (
	"context"
	"errors"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/domain"
	"github.com/kujilabo/cocotola-1.23/cocotola-auth/service"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
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

	var targetOorganization *organization
	var targetAppUser *appUser
	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		rsrf, err := rf.NewRedstartRepositoryFactory(ctx)
		if err != nil {
			return err
		}
		systemAdmin, err := rsuserservice.NewSystemAdmin(ctx, rsrf)
		if err != nil {
			return err
		}
		systemOwner, err := systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
		if err != nil {
			return err
		}

		orgRepo := rsrf.NewOrganizationRepository(ctx)
		tmpOrganization, err := orgRepo.FindOrganizationByName(ctx, systemAdmin, organizationName)
		if err != nil {
			return err
		}

		appUserRepo := rsrf.NewAppUserRepository(ctx)
		verified, err := appUserRepo.VerifyPassword(ctx, systemAdmin, tmpOrganization.OrganizationModel.OrganizationID, loginID, password)
		if err != nil {
			return err
		}

		if !verified {
			return domain.ErrUnauthenticated
		}

		tmpAppUser, err := appUserRepo.FindAppUserByLoginID(ctx, systemOwner, loginID)
		if err != nil {
			return err
		}

		targetAppUser = &appUser{
			appUserID:      tmpAppUser.AppUserModel.AppUserID,
			organizationID: tmpAppUser.AppUserModel.OrganizationID,
			loginID:        tmpAppUser.AppUserModel.LoginID,
			username:       tmpAppUser.AppUserModel.Username,
		}
		targetOorganization = &organization{
			organizationID: tmpOrganization.OrganizationModel.OrganizationID,
			name:           tmpOrganization.OrganizationModel.Name,
		}

		return nil
	}); err != nil {
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
