package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/domain"
	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

const (
	commitSize = 5000
	logSize    = 100000
)

type AdminUsecase struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	logger       *slog.Logger
}

func NewAdminUsecase(txManager, nonTxManager service.TransactionManager) *AdminUsecase {
	adminHandlerLoggerName := "adminUsecase"
	return &AdminUsecase{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		logger:       slog.Default().With(slog.String(rsliblog.LoggerNameKey, adminHandlerLoggerName)),
	}
}

func (u *AdminUsecase) ImportSentences(ctx context.Context, iterator service.TatoebaSentenceAddParameterIterator) (*domain.ImportResult, error) {
	ctx, span := tracer.Start(ctx, "adminUsecase.ImportSentences")
	defer span.End()

	u.logger.InfoContext(ctx, "ImportSentences")
	var readCount = 0
	var importedCount = 0
	var skippedCount = 0
	var loop = true
	for loop {
		if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
			repo := rf.NewTatoebaSentenceRepository(ctx)

			eof, readCountTmp, skipCountTmp, importCountTmp, err := u.importSentences(ctx, iterator, repo)
			if err != nil {
				return err
			}

			readCount += readCountTmp
			skippedCount += skipCountTmp
			importedCount += importCountTmp

			if eof {
				loop = false
				return nil
			}

			return nil
		}); err != nil {
			return nil, rsliberrors.Errorf("import sentence. err: %w", err)
		}
	}

	u.logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importedCount))
	u.logger.InfoContext(ctx, fmt.Sprintf("skipped count: %d", skippedCount))
	u.logger.InfoContext(ctx, fmt.Sprintf("read count: %d", readCount))

	return &domain.ImportResult{
		ImportedCount: importedCount,
		SkippedCount:  skippedCount,
		ReadCount:     readCount,
	}, nil
}

func (u *AdminUsecase) importSentences(ctx context.Context, iterator service.TatoebaSentenceAddParameterIterator, repo service.TatoebaSentenceRepository) (bool, int, int, int, error) {
	readCount := 0
	skipCount := 0
	importedCount := 0
	i := 0
	for {
		param, err := iterator.Next(ctx)
		if errors.Is(err, io.EOF) {
			return true, readCount, skipCount, importedCount, nil
		}
		if err != nil {
			return false, readCount, skipCount, importedCount, rsliberrors.Errorf("read next line. read count: %d, err: %w", readCount, err)
		}
		readCount++
		if readCount%logSize == 0 {
			u.logger.InfoContext(ctx, fmt.Sprintf("read count: %d, imported count: %d", readCount, importedCount))
		}

		if param == nil {
			skipCount++
			continue
		}

		if err := repo.Add(ctx, param); err != nil {
			if errors.Is(err, service.ErrTatoebaSentenceAlreadyExists) {
				u.logger.WarnContext(ctx, fmt.Sprintf("failed to Add. read count: %d, err: %v", readCount, err))
				skipCount++
				continue
			}

			return false, readCount, skipCount, importedCount, rsliberrors.Errorf("failed to Add. read count: %d, err: %w", readCount, err)
		}

		i++
		importedCount++
		if i >= commitSize {
			if importedCount%logSize == 0 {
				u.logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importedCount))
			}
			return false, readCount, skipCount, importedCount, nil
		}
	}
}

func (u *AdminUsecase) ImportLinks(ctx context.Context, iterator service.TatoebaLinkAddParameterIterator) (*domain.ImportResult, error) {
	var readCount = 0
	var importedCount = 0
	var skippedCount = 0
	var loop = true
	for loop {
		if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
			repo := rf.NewTatoebaLinkRepository(ctx)

			eof, readCountTmp, skipCountTmp, importCountTmp, err := u.importLinks(ctx, iterator, repo)
			if err != nil {
				return err
			}
			readCount += readCountTmp
			skippedCount += skipCountTmp
			importedCount += importCountTmp

			if eof {
				loop = false
				return nil
			}

			return nil
		}); err != nil {
			return nil, rsliberrors.Errorf("import sentence. err: %w", err)
		}
	}

	u.logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importedCount))
	u.logger.InfoContext(ctx, fmt.Sprintf("skipped count: %d", skippedCount))
	u.logger.InfoContext(ctx, fmt.Sprintf("read count: %d", readCount))

	return &domain.ImportResult{
		ImportedCount: importedCount,
		SkippedCount:  skippedCount,
		ReadCount:     readCount,
	}, nil
}

func (u *AdminUsecase) importLinks(ctx context.Context, iterator service.TatoebaLinkAddParameterIterator, repo service.TatoebaLinkRepository) (bool, int, int, int, error) {
	eof := false
	readCount := 0
	skippedCount := 0
	importedCount := 0
	i := 0
	for {
		param, err := iterator.Next(ctx)
		if errors.Is(err, io.EOF) {
			eof = true
			return eof, readCount, skippedCount, importedCount, nil
		}
		if err != nil {
			return eof, readCount, skippedCount, importedCount, rsliberrors.Errorf("read next line. read count: %d, err: %w", readCount, err)
		}
		readCount++
		if readCount%logSize == 0 {
			u.logger.InfoContext(ctx, fmt.Sprintf("read count: %d, imported count: %d", readCount, importedCount))
		}

		if param == nil {
			skippedCount++
			continue
		}

		if err := repo.Add(ctx, param); err != nil {
			if !errors.Is(err, service.ErrTatoebaSentenceNotFound) {
				u.logger.WarnContext(ctx, fmt.Sprintf("failed to Add. read count: %d, err: %v", readCount, err))
			}
			skippedCount++
			continue
		}

		i++
		importedCount++
		if i >= commitSize {
			if importedCount%logSize == 0 {
				u.logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importedCount))
			}
			return eof, readCount, skippedCount, importedCount, nil
		}
	}
}
