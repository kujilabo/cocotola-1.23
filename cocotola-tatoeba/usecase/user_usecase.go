package usecase

import (
	"context"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
)

type UserUsecase interface {
	FindSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchCondition) (service.TatoebaSentencePairSearchResult, error)

	FindSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (service.TatoebaSentence, error)
}

type userUsecase struct {
	transaction service.Transaction
}

func NewUserUsecase(transaction service.Transaction) UserUsecase {
	return &userUsecase{
		transaction: transaction,
	}
}

func (u *userUsecase) FindSentencePairs(ctx context.Context, param service.TatoebaSentenceSearchCondition) (service.TatoebaSentencePairSearchResult, error) {
	var result service.TatoebaSentencePairSearchResult
	if err := u.transaction.Do(ctx, func(rf service.RepositoryFactory) error {
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

func (u *userUsecase) FindSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (service.TatoebaSentence, error) {
	var result service.TatoebaSentence
	if err := u.transaction.Do(ctx, func(rf service.RepositoryFactory) error {
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
