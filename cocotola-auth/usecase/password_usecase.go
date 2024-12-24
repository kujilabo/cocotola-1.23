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

type organizationAction struct {
	rf               service.RepositoryFactory
	rsrf             rsuserservice.RepositoryFactory
	organizationRepo rsuserservice.OrganizationRepository
	appUserRepo      rsuserservice.AppUserRepository
	organization     *rsuserservice.Organization
	systemAdmin      *rsuserservice.SystemAdmin
	systemOwner      *rsuserservice.SystemOwner
}

type OrganizationActionOption func(context.Context, *organizationAction) error

func (a *organizationAction) initRsrf(ctx context.Context) error {
	if a.rsrf != nil {
		return nil
	}

	rsrf, err := a.rf.NewRedstartRepositoryFactory(ctx)
	if err != nil {
		return err
	}
	a.rsrf = rsrf
	return nil
}

func (a *organizationAction) initSystemAdmin(ctx context.Context) error {
	if a.systemAdmin != nil {
		return nil
	}

	systemAdmin, err := rsuserservice.NewSystemAdmin(ctx, a.rsrf)
	if err != nil {
		return err
	}
	a.systemAdmin = systemAdmin
	return nil
}

func (a *organizationAction) initSystemOwnerByOrganizationName(ctx context.Context, organizationName string) error {
	if a.systemOwner != nil {
		return nil
	}

	if err := a.initSystemAdmin(ctx); err != nil {
		return err
	}
	systemOwner, err := a.systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
	if err != nil {
		return err
	}
	a.systemOwner = systemOwner
	return nil
}

func WithSystemAdmin() OrganizationActionOption {
	return func(ctx context.Context, action *organizationAction) error {
		if err := action.initSystemAdmin(ctx); err != nil {
			return err
		}
		return nil
	}
}

func WithOrganizationRepository() OrganizationActionOption {
	return func(ctx context.Context, action *organizationAction) error {
		if err := action.initRsrf(ctx); err != nil {
			return err
		}
		action.organizationRepo = action.rsrf.NewOrganizationRepository(ctx)
		return nil
	}
}

func WithAppUserRepository() OrganizationActionOption {
	return func(ctx context.Context, action *organizationAction) error {
		if err := action.initRsrf(ctx); err != nil {
			return err
		}
		action.appUserRepo = action.rsrf.NewAppUserRepository(ctx)
		return nil
	}
}

func WithOrganization(organizationName string) OrganizationActionOption {
	return func(ctx context.Context, action *organizationAction) error {
		if err := action.initSystemOwnerByOrganizationName(ctx, organizationName); err != nil {
			return err
		}
		return nil
	}
}

func NewOrganizationAction(ctx context.Context, rf service.RepositoryFactory, options ...OrganizationActionOption) (*organizationAction, error) {
	action := organizationAction{}
	action.rf = rf
	for _, option := range options {
		if err := option(ctx, &action); err != nil {
			return nil, err
		}
	}
	return &action, nil
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
