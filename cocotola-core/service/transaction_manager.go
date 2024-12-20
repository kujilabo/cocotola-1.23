package service

import "context"

type TransactionManager interface {
	Do(ctx context.Context, fn func(rf RepositoryFactory) error) error
}

func Do1[T any](ctx context.Context, txManager TransactionManager, fn func(rf RepositoryFactory) (T, error)) (T, error) {
	var t1 T
	err := txManager.Do(ctx, func(rf RepositoryFactory) error {
		var t1tmp T
		t1tmp, err := fn(rf)
		t1 = t1tmp
		return err
	})
	return t1, err
}
