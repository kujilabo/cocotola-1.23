package service

import (
	"context"

	rsuserdomain "github.com/kujilabo/cocotola-1.23/redstart/user/domain"
	rsuserservice "github.com/kujilabo/cocotola-1.23/redstart/user/service"

	rslibservice "github.com/kujilabo/cocotola-1.23/redstart/lib/service"
)

type OperatorInterface interface {
	AppUserID() *rsuserdomain.AppUserID
	OrganizationID() *rsuserdomain.OrganizationID
	// LoginID() string
	// Username() string
}
type RepositoryFactory interface {
	NewRedstartRepositoryFactory(ctx context.Context) (rsuserservice.RepositoryFactory, error)

	NewStateRepository(ctx context.Context) (StateRepository, error)
}
type TransactionManager rslibservice.TransactionManagerT[RepositoryFactory]
