package config

import (
	"context"
	"database/sql"
	"io/fs"

	"gorm.io/gorm"

	libdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
)

type DBConfig struct {
	DriverName string                     `yaml:"driverName"`
	MySQL      *libgateway.MySQLConfig    `yaml:"mysql"`
	Postgres   *libgateway.PostgresConfig `yaml:"postgres"`
	SQLite3    *libgateway.SQLite3Config  `yaml:"sqlite3"`
	Migration  bool                       `yaml:"migration"`
}

type mergedFS struct {
	fss     []fs.FS
	entries []fs.DirEntry
}

func MergeFS(driverName string, fss ...fs.FS) (*mergedFS, error) {
	entries := make([]fs.DirEntry, 0)
	for i := range fss {
		e, err := fs.ReadDir(fss[i], driverName)
		if err != nil {
			return nil, liberrors.Errorf("read %q directory: %w", driverName, err)
		}
		entries = append(entries, e...)
	}

	return &mergedFS{
		fss:     fss,
		entries: entries,
	}, nil
}

func (f *mergedFS) Open(name string) (fs.File, error) {
	var file fs.File
	var err error
	for i := range f.fss {
		file, err = f.fss[i].Open(name)
		if err == nil {
			return file, nil
		}
	}

	return nil, err
}

func (f *mergedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return f.entries, nil
}
func InitDB(ctx context.Context, cfg *DBConfig, sqlFSs ...fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	mergedFS, err := MergeFS(cfg.DriverName, sqlFSs...)
	if err != nil {
		return nil, nil, nil, liberrors.Errorf("merge sql files in %q directory: %w", cfg.DriverName, err)
	}

	initDBFunc, ok := initDBs[cfg.DriverName]
	if !ok {
		return nil, nil, nil, libdomain.ErrInvalidArgument
	}
	return initDBFunc(ctx, cfg, mergedFS)
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
