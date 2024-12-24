package student

import (
	"context"

	rslibservice "github.com/kujilabo/cocotola-1.23/redstart/lib/service"

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
	return rslibservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) (*domain.WorkbookID, error) {
		workbookRepo, err := rf.NewWorkbookRepository(ctx)
		if err != nil {
			return nil, err
		}
		workbookID, err := workbookRepo.AddWorkbook(ctx, operator, param)
		if err != nil {
			return nil, err
		}
		return workbookID, nil
	})
}

func (u *WorkbookCommandUsecase) UpdateWorkbook(ctx context.Context, operator service.OperatorInterface, workbookID *domain.WorkbookID, version int, param *service.WorkbookUpdateParameter) error {
	return rslibservice.Do0(ctx, u.txManager, func(rf service.RepositoryFactory) error {
		workbookRepo, err := rf.NewWorkbookRepository(ctx)
		if err != nil {
			return nil
		}

		if err := workbookRepo.UpdateWorkbook(ctx, operator, workbookID, version, param); err != nil {
			return err
		}

		return nil
	})
}
