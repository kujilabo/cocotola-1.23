package gateway

import (
	"context"
	"time"

	"gorm.io/gorm"

	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"

	"github.com/kujilabo/cocotola-1.23/cocotola-tatoeba/service"
)

type repositoryFactory struct {
	dialect    rslibgateway.DialectRDBMS
	driverName string
	db         *gorm.DB
	location   *time.Location
}

func NewRepositoryFactory(ctx context.Context, dialect rslibgateway.DialectRDBMS, driverName string, db *gorm.DB, location *time.Location) (service.RepositoryFactory, error) {
	if db == nil {
		return nil, rsliberrors.Errorf("db is nil. err: %w", rslibdomain.ErrInvalidArgument)
	}

	return &repositoryFactory{
		dialect:    dialect,
		driverName: driverName,
		db:         db,
		location:   location,
	}, nil
}

func (f *repositoryFactory) NewTatoebaSentenceRepository(ctx context.Context) service.TatoebaSentenceRepository {
	return newTatoebaSentenceRepository(f.db)
}

func (f *repositoryFactory) NewTatoebaLinkRepository(ctx context.Context) service.TatoebaLinkRepository {
	return newTatoebaLinkRepository(f.db, f)
}

type RepositoryFactoryFunc func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error)
