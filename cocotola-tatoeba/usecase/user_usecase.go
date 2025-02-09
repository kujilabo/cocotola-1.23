package usecase

import (
	"context"
	"log/slog"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	rslibservice "github.com/kujilabo/cocotola-1.23/redstart/lib/service"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

type UserUsecase struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	logger       *slog.Logger
}

func NewUserUsecase(txManager, nonTxManager service.TransactionManager) *UserUsecase {
	userUsecaseLoggerName := "userUsecase"
	return &UserUsecase{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		logger:       slog.Default().With(slog.String(rsliblog.LoggerNameKey, userUsecaseLoggerName)),
	}
}

func (u *UserUsecase) FindSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchConditionInterface) (*service.TatoebaSentencePairSearchResult, error) {
	var result *service.TatoebaSentencePairSearchResult
	if err := u.nonTxManager.Do(ctx, func(rf service.RepositoryFactory) error {
		repo := rf.NewTatoebaSentenceRepository(ctx)

		tmpResult, err := repo.FindTatoebaSentencePairs(ctx, param)
		if err != nil {
			return rsliberrors.Errorf("execute FindTatoebaSentencePairs. err: %w", err)
		}
		result = tmpResult
		return nil
	}); err != nil {
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
