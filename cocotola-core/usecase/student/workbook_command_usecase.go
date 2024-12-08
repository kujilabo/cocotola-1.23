package student

import (
	"context"

	"github.com/kujilabo/cocotola-1.23/cocotola-core/domain"
	"github.com/kujilabo/cocotola-1.23/cocotola-core/service"
)

type WorkbookCommandUsecase struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
}

func NewWorkbookCommandUsecase(txManager, nonTxManager service.TransactionManager) *WorkbookCommandUsecase {
	return &WorkbookCommandUsecase{
		txManager:    txManager,
		nonTxManager: nonTxManager,
	}
}

func (u *WorkbookCommandUsecase) AddWorkbook(ctx context.Context, operator service.OperatorInterface, param *service.WorkbookAddParameter) (*domain.WorkbookID, error) {
	var workbookID *domain.WorkbookID

	fn := func(workbookRepository service.WorkbookRepository) error {
		tmpWorkbookID, err := workbookRepository.AddWorkbook(ctx, operator, param)
		if err != nil {
			return err
		}

		workbookID = tmpWorkbookID
		return nil
	}

	if err := u.workbookFunction(ctx, operator, fn); err != nil {
		return nil, err
	}

	return workbookID, nil
}

func (u *WorkbookCommandUsecase) UpdateWorkbook(ctx context.Context, operator service.OperatorInterface, workbookID *domain.WorkbookID, version int, param *service.WorkbookUpdateParameter) error {
	fn := func(workbookRepository service.WorkbookRepository) error {
		if err := workbookRepository.UpdateWorkbook(ctx, operator, workbookID, version, param); err != nil {
			return err
		}

		return nil
	}

	if err := u.workbookFunction(ctx, operator, fn); err != nil {
		return err
	}

	return nil
}

func (u *WorkbookCommandUsecase) workbookFunction(ctx context.Context, operator service.OperatorInterface, fn func(workbookRepository service.WorkbookRepository) error) error {
	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		workbookRepo, err := rf.NewWorkbookRepository(ctx)
		if err != nil {
			return err
		}
		if err := fn(workbookRepo); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
