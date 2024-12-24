package service

import "context"

type TransactionManagerT[RF any] interface {
	Do(ctx context.Context, fn func(rf RF) error) error
}

func Do0[RF any](ctx context.Context, txManager TransactionManagerT[RF], fn func(rf RF) error) error {
	return txManager.Do(ctx, func(rf RF) error {
		return fn(rf)
	})
}

func Do1[RF any, T1 any](ctx context.Context, txManager TransactionManagerT[RF], fn func(rf RF) (T1, error)) (T1, error) {
	var t1 T1
	err := txManager.Do(ctx, func(rf RF) error {
		var t1tmp T1
		t1tmp, err := fn(rf)
		t1 = t1tmp
		return err
	})
	return t1, err
}
