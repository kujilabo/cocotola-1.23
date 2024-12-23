package service

import (
	"context"

	rsuserservice "github.com/kujilabo/cocotola-1.23/redstart/user/service"

	rslibservice "github.com/kujilabo/cocotola-1.23/redstart/lib/service"
)

type RepositoryFactory interface {
	NewRedstartRepositoryFactory(ctx context.Context) (rsuserservice.RepositoryFactory, error)

	NewStateRepository(ctx context.Context) (StateRepository, error)
}
type TransactionManager rslibservice.TransactionManagerT[RepositoryFactory]
