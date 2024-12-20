package config

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"
	"os"

	"gorm.io/gorm"

	libdomain "github.com/kujilabo/cocotola-1.23/redstart/lib/domain"
	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
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
			return nil, err
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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	mergedFS, err := newMergedFS(cfg.DriverName, sqlFSs...)
	if err != nil {
		return nil, nil, nil, err
	}

	switch cfg.DriverName {
	case "sqlite3":
		return initSqlite3(ctx, cfg, mergedFS, logger)
	case "mysql":
		return initMySQL(ctx, cfg, mergedFS, logger)
	case "postgres":
		return initPostgres(ctx, cfg, mergedFS, logger)
	default:
		return nil, nil, nil, libdomain.ErrInvalidArgument
	}
}

func initSqlite3(ctx context.Context, cfg *DBConfig, fs fs.FS, logger *slog.Logger) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := libgateway.OpenSQLite3("./"+cfg.SQLite3.File, logger)
	if err != nil {
		return nil, nil, nil, liberrors.Errorf("OpenSQLite. err: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, nil, nil, err
	}

	if cfg.Migration {
		if err := libgateway.MigrateSQLite3DB(db, fs); err != nil {
			return nil, nil, nil, err
		}
	}

	dialect := libgateway.DialectMySQL{}
	return &dialect, db, sqlDB, nil
}

func initMySQL(ctx context.Context, cfg *DBConfig, fs fs.FS, logger *slog.Logger) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := libgateway.OpenMySQL(cfg.MySQL.Username, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Database, logger)
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

func initPostgres(ctx context.Context, cfg *DBConfig, fs fs.FS, logger *slog.Logger) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := libgateway.OpenPostgres(cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database, logger)
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
