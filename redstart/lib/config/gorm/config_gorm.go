package gorm

import (
	"context"
	"database/sql"
	"io/fs"

	"gorm.io/gorm"

	libconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	libdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
)

type DBInitializer func(context.Context, *libconfig.DBConfig, fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error)

func InitDB(ctx context.Context, cfg *libconfig.DBConfig, initializer map[string]DBInitializer, sqlFSs ...fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	mergedFS, err := libconfig.MergeFS(cfg.DriverName, sqlFSs...)
	if err != nil {
		return nil, nil, nil, liberrors.Errorf("merge sql files in %q directory: %w", cfg.DriverName, err)
	}

	initializerFunc, ok := initializer[cfg.DriverName]
	if !ok {
		return nil, nil, nil, libdomain.ErrInvalidArgument
	}
	return initializerFunc(ctx, cfg, mergedFS)
	// switch cfg.DriverName {
	// case "sqlite3":
	//
	//	return initSqlite3(ctx, cfg, mergedFS)
	//
	// case "mysql":
	//
	//	return initMySQL(ctx, cfg, mergedFS)
	//
	// case "postgres":
	//
	//	return initPostgres(ctx, cfg, mergedFS)
	//
	// default:
	//
	//		return nil, nil, nil, libdomain.ErrInvalidArgument
	//	}
}
