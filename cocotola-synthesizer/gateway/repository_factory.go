package gateway

import (
	"context"
	"time"

	"gorm.io/gorm"

	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/service"
)

type RepositoryFactory struct {
	dialect    rslibgateway.DialectRDBMS
	driverName string
	db         *gorm.DB
	location   *time.Location
}

func NewRepositoryFactory(ctx context.Context, dialect rslibgateway.DialectRDBMS, driverName string, db *gorm.DB, location *time.Location) (*RepositoryFactory, error) {
	if db == nil {
		return nil, rsliberrors.Errorf("db is nil. err: %w", rslibdomain.ErrInvalidArgument)
	}

	return &RepositoryFactory{
		dialect:    dialect,
		driverName: driverName,
		db:         db,
		location:   location,
	}, nil
}

func (f *RepositoryFactory) NewAudioRepository(ctx context.Context) service.AudioRepository {
	return newAudioRepository(ctx, f.db)
}

type RepositoryFactoryFunc func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error)
