package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
)

const (
	commitSize = 1000
	logSize    = 100000
)

type AdminUsecase interface {
	ImportSentences(ctx context.Context, iterator service.TatoebaSentenceAddParameterIterator) error

	ImportLinks(ctx context.Context, iterator service.TatoebaLinkAddParameterIterator) error
}

type adminUsecase struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
}

func NewAdminUsecase(txManager, nonTxManager service.TransactionManager) AdminUsecase {
	return &adminUsecase{
		txManager:    txManager,
		nonTxManager: nonTxManager,
	}
}

func (u *adminUsecase) ImportSentences(ctx context.Context, iterator service.TatoebaSentenceAddParameterIterator) error {
	ctx, span := tracer.Start(ctx, "adminUsecase.ImportSentences")
	defer span.End()

	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

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

	logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importCount))
	logger.InfoContext(ctx, fmt.Sprintf("skipped count: %d", skipCount))
	logger.InfoContext(ctx, fmt.Sprintf("read count: %d", readCount))

	return nil
}

func (u *adminUsecase) importSentences(ctx context.Context, iterator service.TatoebaSentenceAddParameterIterator, repo service.TatoebaSentenceRepository) (bool, int, int, int, error) {
	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

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
				logger.WarnContext(ctx, fmt.Sprintf("failed to Add. read count: %d, err: %v", readCount, err))
				skipCount++
				continue
			}

			return false, readCount, skipCount, importCount, rsliberrors.Errorf("failed to Add. read count: %d, err: %w", readCount, err)
		}

		i++
		importCount++
		if i >= commitSize {
			if importCount%logSize == 0 {
				logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importCount))
			}
			return false, readCount, skipCount, importCount, nil
		}
	}
}

func (u *adminUsecase) ImportLinks(ctx context.Context, iterator service.TatoebaLinkAddParameterIterator) error {
	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

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

	logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importCount))
	logger.InfoContext(ctx, fmt.Sprintf("skipped count: %d", skipCount))
	logger.InfoContext(ctx, fmt.Sprintf("read count: %d", readCount))

	return nil
}

func (u *adminUsecase) importLinks(ctx context.Context, iterator service.TatoebaLinkAddParameterIterator, repo service.TatoebaLinkRepository) (bool, int, int, int, error) {
	ctx = rsliblog.WithLoggerName(ctx, loggerKey)
	logger := rsliblog.GetLoggerFromContext(ctx, loggerKey)

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
				logger.WarnContext(ctx, fmt.Sprintf("failed to Add. read count: %d, err: %v", readCount, err))
			}
			skipCount++
			return eof, readCount, skipCount, importCount, nil
		}
		i++
		importCount++
		if i >= commitSize {
			if importCount%logSize == 0 {
				logger.InfoContext(ctx, fmt.Sprintf("imported count: %d", importCount))
			}
			return eof, readCount, skipCount, importCount, nil
		}
	}
}
