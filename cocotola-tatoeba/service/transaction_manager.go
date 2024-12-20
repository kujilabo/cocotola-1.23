package service

import "context"

// type TransactionManager interface {
// 	Do(ctx context.Context, fn func(rf RepositoryFactory) error) error
// }

type TransactionManagerT[T any] interface {
	Do(ctx context.Context, fn func(rf T) error) error
}

type TransactionManager TransactionManagerT[RepositoryFactory]
