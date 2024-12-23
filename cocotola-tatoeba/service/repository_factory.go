package service

import (
	"context"

	rslibservice "github.com/kujilabo/cocotola-1.23/redstart/lib/service"
)

type RepositoryFactory interface {
	NewTatoebaLinkRepository(ctx context.Context) TatoebaLinkRepository

	NewTatoebaSentenceRepository(ctx context.Context) TatoebaSentenceRepository
}

type TransactionManager rslibservice.TransactionManagerT[RepositoryFactory]
