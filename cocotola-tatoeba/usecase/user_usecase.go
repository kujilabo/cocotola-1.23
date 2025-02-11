package usecase

import (
	"context"
	"log/slog"
	"time"

	"github.com/patrickmn/go-cache"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	rslibservice "github.com/kujilabo/cocotola-1.23/redstart/lib/service"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

type UserUsecase struct {
	txManager             service.TransactionManager
	nonTxManager          service.TransactionManager
	findSentencePairCache *cache.Cache
	logger                *slog.Logger
}

func NewUserUsecase(txManager, nonTxManager service.TransactionManager) *UserUsecase {
	userUsecaseLoggerName := "userUsecase"
	return &UserUsecase{
		txManager:             txManager,
		nonTxManager:          nonTxManager,
		findSentencePairCache: cache.New(5*60*time.Second, 10*60*time.Second),
		logger:                slog.Default().With(slog.String(rsliblog.LoggerNameKey, userUsecaseLoggerName)),
	}
}

func (u *UserUsecase) FindSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchConditionInterface) (*service.TatoebaSentencePairSearchResult, error) {
	result, err := rslibservice.Do1(ctx, u.nonTxManager, func(rf service.RepositoryFactory) (*service.TatoebaSentencePairSearchResult, error) {
		repo := rf.NewTatoebaSentenceRepository(ctx)

		sentences, err := repo.FindTatoebaSentencePairs(ctx, param)
		if err != nil {
			return nil, rsliberrors.Errorf("execute FindTatoebaSentencePairs. err: %w", err)
		}

		if cachedCount, ok := u.findSentencePairCache.Get(param.ToString()); ok {
			if _cachedCountInt, ok := cachedCount.(int); ok {
				return service.NewTatoebaSentencePairSearchResult(_cachedCountInt, sentences), nil
			}
			u.logger.ErrorContext(ctx, "cachedCount is not int")
		}

		count, err := repo.CountTatoebaSentencePairs(ctx, param)
		if err != nil {
			return nil, rsliberrors.Errorf("execute CountTatoebaSentencePairs. err: %w", err)
		}
		u.findSentencePairCache.Set(param.ToString(), count, cache.DefaultExpiration)

		return service.NewTatoebaSentencePairSearchResult(count, sentences), nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserUsecase) FindSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (*service.TatoebaSentence, error) {
	return rslibservice.Do1(ctx, u.nonTxManager, func(rf service.RepositoryFactory) (*service.TatoebaSentence, error) {
		repo := rf.NewTatoebaSentenceRepository(ctx)

		result, err := repo.FindTatoebaSentenceBySentenceNumber(ctx, sentenceNumber)
		if err != nil {
			return nil, rsliberrors.Errorf("execute FindTatoebaSentenceBySentenceNumber. err: %w", err)
		}
		return result, nil
	})
}
