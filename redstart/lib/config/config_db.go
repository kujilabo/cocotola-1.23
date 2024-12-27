package config

import (
	"context"
	"database/sql"
	"io/fs"

	"gorm.io/gorm"

	libdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
	// libgatewaysqlite3 "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway/sqlite3"
)

type SQLite3Config struct {
	File string `yaml:"file" validate:"required"`
}

type MySQLConfig struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	Database string `yaml:"database" validate:"required"`
}

type PostgresConfig struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	Database string `yaml:"database" validate:"required"`
}

type DBConfig struct {
	DriverName string          `yaml:"driverName"`
	SQLite3    *SQLite3Config  `yaml:"sqlite3"`
	MySQL      *MySQLConfig    `yaml:"mysql"`
	Postgres   *PostgresConfig `yaml:"postgres"`
	Migration  bool            `yaml:"migration"`
}

type mergedFS struct {
	fss     []fs.FS
	entries []fs.DirEntry
}

func newMergedFS(driverName string, fss ...fs.FS) (*mergedFS, error) {
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

type DBInitializer func(context.Context, *DBConfig, fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error)

func InitDB(ctx context.Context, cfg *DBConfig, initializer map[string]DBInitializer, sqlFSs ...fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	mergedFS, err := newMergedFS(cfg.DriverName, sqlFSs...)
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

func InitMySQL(ctx context.Context, cfg *DBConfig, fs fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := libgateway.OpenMySQL(cfg.MySQL.Username, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Database)
	if err != nil {
		return nil, nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, nil, nil, err
	}

	if err := libgateway.MigrateMySQLDB(db, fs); err != nil {
		return nil, nil, nil, liberrors.Errorf("failed to MigrateMySQLDB. err: %w", err)
	}

	dialect := libgateway.DialectMySQL{}
	return &dialect, db, sqlDB, nil
}

func initPostgres(ctx context.Context, cfg *DBConfig, fs fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := libgateway.OpenPostgres(cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database)
	if err != nil {
		return nil, nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, nil, nil, err
	}

	if err := libgateway.MigratePostgresDB(db, fs); err != nil {
		return nil, nil, nil, liberrors.Errorf("failed to MigrateMySQLDB. err: %w", err)
	}

	dialect := libgateway.DialectPostgres{}
	return &dialect, db, sqlDB, nil
}
