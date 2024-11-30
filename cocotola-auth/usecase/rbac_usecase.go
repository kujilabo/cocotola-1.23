package usecase

import (
	"context"

	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"
	rsuserservice "github.com/kujilabo/cocotola-1.23/redstart/user/service"

	"github.com/kujilabo/cocotola-1.23/cocotola-auth/service"
)

type rbac struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
}

func NewRBAC(txManager, nonTxManager service.TransactionManager) *rbac {
	return &rbac{
		txManager:    txManager,
		nonTxManager: nonTxManager,
	}
}

func (u *rbac) AddPolicyToUser(ctx context.Context, organizationID *rsuserdomain.OrganizationID, subject rsuserdomain.RBACSubject, action rsuserdomain.RBACAction, object rsuserdomain.RBACObject, effect rsuserdomain.RBACEffect) error {
	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		rsrf, err := rf.NewRedstartRepositoryFactory(ctx)
		if err != nil {
			return err
		}

		sysAdmin, err := rsuserservice.NewSystemAdmin(ctx, rsrf)
		if err != nil {
			return err
		}

		authorizationManager := rsrf.NewAuthorizationManager(ctx)
		if err := authorizationManager.AddPolicyToUserBySystemAdmin(ctx, sysAdmin, organizationID, subject, action, object, effect); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
