package usecase

import (
	"context"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
)

type UserUsecase struct {
	txManager service.TransactionManager
}

func NewUserUsecase(txManager service.TransactionManager) *UserUsecase {
	return &UserUsecase{
		txManager: txManager,
	}
}

func (u *UserUsecase) FindSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchCondition) (service.TatoebaSentencePairSearchResult, error) {
	var result service.TatoebaSentencePairSearchResult
	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
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

func (u *UserUsecase) FindSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (service.TatoebaSentence, error) {
	var result service.TatoebaSentence
	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		repo := rf.NewTatoebaSentenceRepository(ctx)

		tmpResult, err := repo.FindTatoebaSentenceBySentenceNumber(ctx, sentenceNumber)
		if err != nil {
			return rsliberrors.Errorf("execute FindTatoebaSentenceBySentenceNumber. err: %w", err)
		}
		result = tmpResult
		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}
