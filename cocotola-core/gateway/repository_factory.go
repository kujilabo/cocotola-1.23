package gateway

import (
	"context"
	"time"

	"gorm.io/gorm"

	rslibdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"

	"github.com/kujilabo/cocotola-1.23/cocotola-core/service"
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

func (f *RepositoryFactory) NewWorkbookRepository(ctx context.Context) (service.WorkbookRepository, error) {
	return NewWorkbookRepository(f.db), nil
}

type RepositoryFactoryFunc func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error)