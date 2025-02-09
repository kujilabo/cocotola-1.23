package usecase

import (
	"context"

	rslibservice "github.com/kujilabo/cocotola-1.23/redstart/lib/service"
	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"
	rsuserservice "github.com/kujilabo/cocotola-1.23/redstart/user/service"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/service"
)

type RBACUsecase struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
}

func NewRBACUsecase(txManager, nonTxManager service.TransactionManager) *RBACUsecase {
	return &RBACUsecase{
		txManager:    txManager,
		nonTxManager: nonTxManager,
	}
}

func (u *RBACUsecase) AddPolicyToUser(ctx context.Context, organizationID *rsuserdomain.OrganizationID, subject rsuserdomain.RBACSubject, action rsuserdomain.RBACAction, object rsuserdomain.RBACObject, effect rsuserdomain.RBACEffect) error {
	return rslibservice.Do0(ctx, u.txManager, func(rf service.RepositoryFactory) error {
		rsrf, err := rf.NewRedstartRepositoryFactory(ctx)
		if err != nil {
			return err
		}

		sysAdmin, err := rsuserservice.NewSystemAdmin(ctx, rsrf)
		if err != nil {
			return err
		}

		authorizationManager, err := rsrf.NewAuthorizationManager(ctx)
		if err != nil {
			return err
		}

		if err := authorizationManager.AddPolicyToUserBySystemAdmin(ctx, sysAdmin, organizationID, subject, action, object, effect); err != nil {
			return err
		}

		return nil
	})
}

func (u *RBACUsecase) Authorize(ctx context.Context, operator service.OperatorInterface, action rsuserdomain.RBACAction, object rsuserdomain.RBACObject) (bool, error) {
	return rslibservice.Do1(ctx, u.nonTxManager, func(rf service.RepositoryFactory) (bool, error) {
		rsrf, err := rf.NewRedstartRepositoryFactory(ctx)
		if err != nil {
			return false, err
		}

		authorizationManager, err := rsrf.NewAuthorizationManager(ctx)
		if err != nil {
			return false, err
		}

		return authorizationManager.Authorize(ctx, operator, action, object)
	})
}
