package service

import (
	"context"

	rslibservice "github.com/kujilabo/cocotola-1.23/redstart/lib/service"
)

type RepositoryFactory interface {
	NewWorkbookRepository(ctx context.Context) (WorkbookRepository, error)
}

type TransactionManager rslibservice.TransactionManagerT[RepositoryFactory]
