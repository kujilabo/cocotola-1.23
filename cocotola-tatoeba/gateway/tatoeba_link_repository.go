package gateway

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rsliblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"

	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

type tatoebaLinkRepository struct {
	db        *gorm.DB
	rf        service.RepositoryFactory
	linkCache *cache.Cache
	logger    *slog.Logger
}

type tatoebaLinkEntity struct {
	Src int
	Dst int
}

func (e *tatoebaLinkEntity) TableName() string {
	return "tatoeba_link"
}

func newTatoebaLinkRepository(db *gorm.DB, rf service.RepositoryFactory) service.TatoebaLinkRepository {
	slog.Default().InfoContext(context.Background(), "newTatoebaLinkRepository")
	return &tatoebaLinkRepository{
		db:        db,
		rf:        rf,
		linkCache: cache.New(5*time.Minute, 10*time.Minute),
		logger:    slog.Default().With(slog.String(rsliblog.LoggerNameKey, "TatoebaLinkRepository")),
	}
}

func (r *tatoebaLinkRepository) Add(ctx context.Context, param service.TatoebaLinkAddParameter) error {
	sentenceRepo := r.rf.NewTatoebaSentenceRepository(ctx)
	isSrcContainedCache, srcFound := r.linkCache.Get(strconv.Itoa(param.GetSrc()))
	if !srcFound {
		isSrcContainedInDB, err := sentenceRepo.ContainsSentenceBySentenceNumber(ctx, param.GetSrc())
		if err != nil {
			return err
		}
		r.linkCache.Set(strconv.Itoa(param.GetSrc()), isSrcContainedInDB, cache.DefaultExpiration)
		isSrcContainedCache = isSrcContainedInDB
	}

	isDstContainedCache, dstFound := r.linkCache.Get(strconv.Itoa(param.GetDst()))
	if !dstFound {
		isDstContainedInDB, err := sentenceRepo.ContainsSentenceBySentenceNumber(ctx, param.GetDst())
		if err != nil {
			return err
		}
		r.linkCache.Set(strconv.Itoa(param.GetDst()), isDstContainedInDB, cache.DefaultExpiration)
		isDstContainedCache = isDstContainedInDB
	}

	isSrcContained, ok := isSrcContainedCache.(bool)
	if !ok {
		return rsliberrors.Errorf("failed to Add tatoebaLink. err: %w", service.ErrTatoebaSentenceNotFound)
	}

	isDstContained, ok := isDstContainedCache.(bool)
	if !ok {
		return rsliberrors.Errorf("failed to Add tatoebaLink. err: %w", service.ErrTatoebaSentenceNotFound)
	}

	if !isSrcContained || !isDstContained {
		return service.ErrTatoebaSentenceNotFound
	}

	entity := tatoebaLinkEntity{
		Src: param.GetSrc(),
		Dst: param.GetDst(),
	}

	if result := r.db.Create(&entity); result.Error != nil {
		if err := libgateway.ConvertDuplicatedError(result.Error, service.ErrTatoebaLinkAlreadyExists); errors.Is(err, service.ErrTatoebaLinkAlreadyExists) {
			return rsliberrors.Errorf("failed to Add tatoebaLink. err: %w", err)
		}

		if err := libgateway.ConvertRelationError(result.Error, service.ErrTatoebaLinkSourceNotFound); errors.Is(err, service.ErrTatoebaLinkSourceNotFound) {
			// fmt.Printf("relation %v, %v\n", fromContained, toContained)
			// nothing
			return nil
		}

		return rsliberrors.Errorf("failed to Add tatoebaLink. err: %w", result.Error)
	}

	return nil
}
