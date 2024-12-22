package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

const (
	commitSize = 1000
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

func (u *AdminUsecase) ImportSentences(ctx context.Context, iterator service.TatoebaSentenceAddParameterIterator) error {
	ctx, span := tracer.Start(ctx, "adminUsecase.ImportSentences")
	defer span.End()

	var readCount = 0
	var importCount = 0
	var skipCount = 0
	var loop = true
	for loop {
		if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
			repo := rf.NewTatoebaSentenceRepository(ctx)

			eof, readCountTmp, skipCountTmp, importCountTmp, err := u.importSentences(ctx, iterator, repo)
			if err != nil {
				return err
			}

			readCount += readCountTmp
			skipCount += skipCountTmp
			importCount += importCountTmp

			if eof {
				loop = false
				return nil
			}

			return nil
		}); err != nil {
			return rsliberrors.Errorf("import sentence. err: %w", err)
		}
	}

	u.logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importCount))
	u.logger.InfoContext(ctx, fmt.Sprintf("skipped count: %d", skipCount))
	u.logger.InfoContext(ctx, fmt.Sprintf("read count: %d", readCount))

	return nil
}

func (u *AdminUsecase) importSentences(ctx context.Context, iterator service.TatoebaSentenceAddParameterIterator, repo service.TatoebaSentenceRepository) (bool, int, int, int, error) {
	readCount := 0
	skipCount := 0
	importCount := 0
	i := 0
	for {
		param, err := iterator.Next(ctx)
		if errors.Is(err, io.EOF) {
			return true, readCount, skipCount, importCount, nil
		}
		if err != nil {
			return false, readCount, skipCount, importCount, rsliberrors.Errorf("read next line. read count: %d, err: %w", readCount, err)
		}
		readCount++

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

			return false, readCount, skipCount, importCount, rsliberrors.Errorf("failed to Add. read count: %d, err: %w", readCount, err)
		}

		i++
		importCount++
		if i >= commitSize {
			if importCount%logSize == 0 {
				u.logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importCount))
			}
			return false, readCount, skipCount, importCount, nil
		}
	}
}

func (u *AdminUsecase) ImportLinks(ctx context.Context, iterator service.TatoebaLinkAddParameterIterator) error {
	var readCount = 0
	var importCount = 0
	var skipCount = 0
	var loop = true
	for loop {
		if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
			repo := rf.NewTatoebaLinkRepository(ctx)

			eof, readCountTmp, skipCountTmp, importCountTmp, err := u.importLinks(ctx, iterator, repo)
			if err != nil {
				return err
			}
			if eof {
				loop = false
				return nil
			}
			readCount += readCountTmp
			skipCount += skipCountTmp
			importCount += importCountTmp

			// i := 0
			// for {
			// 	param, err := iterator.Next(ctx)
			// 	if errors.Is(err, io.EOF) {
			// 		loop = false
			// 		break
			// 	}
			// 	readCount++
			// 	if err != nil {
			// 		return rsliberrors.Errorf("read next line. read count: %d, err: %w", readCount, err)
			// 	}
			// 	if param == nil {
			// 		skipCount++
			// 		continue
			// 	}

			// 	if err := repo.Add(ctx, param); err != nil {
			// 		if !errors.Is(err, service.ErrTatoebaSentenceNotFound) {
			// 			logger.WarnContext(ctx, fmt.Sprintf("failed to Add. read count: %d, err: %v", readCount, err))
			// 		}
			// 		skipCount++
			// 		continue
			// 	}
			// 	i++
			// 	importCount++
			// 	if i >= commitSize {
			// 		if importCount%logSize == 0 {
			// 			logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importCount))
			// 		}
			// 		break
			// 	}
			// }

			return nil
		}); err != nil {
			return rsliberrors.Errorf("import sentence. err: %w", err)
		}
	}

	u.logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importCount))
	u.logger.InfoContext(ctx, fmt.Sprintf("skipped count: %d", skipCount))
	u.logger.InfoContext(ctx, fmt.Sprintf("read count: %d", readCount))

	return nil
}

func (u *AdminUsecase) importLinks(ctx context.Context, iterator service.TatoebaLinkAddParameterIterator, repo service.TatoebaLinkRepository) (bool, int, int, int, error) {
	eof := false
	readCount := 0
	skipCount := 0
	importCount := 0
	i := 0
	for {
		param, err := iterator.Next(ctx)
		if errors.Is(err, io.EOF) {
			eof = true
			return eof, readCount, skipCount, importCount, nil
		}
		readCount++
		if err != nil {
			return eof, readCount, skipCount, importCount, rsliberrors.Errorf("read next line. read count: %d, err: %w", readCount, err)
		}

		if param == nil {
			skipCount++
			return eof, readCount, skipCount, importCount, nil
		}

		if err := repo.Add(ctx, param); err != nil {
			if !errors.Is(err, service.ErrTatoebaSentenceNotFound) {
				u.logger.WarnContext(ctx, fmt.Sprintf("failed to Add. read count: %d, err: %v", readCount, err))
			}
			skipCount++
			return eof, readCount, skipCount, importCount, nil
		}
		i++
		importCount++
		if i >= commitSize {
			if importCount%logSize == 0 {
				u.logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importCount))
			}
			return eof, readCount, skipCount, importCount, nil
		}
	}
}
