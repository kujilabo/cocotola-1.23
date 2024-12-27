package gateway

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log/slog"
	"time"

	"github.com/golang-migrate/migrate/v4/database"
	migrate_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	slog_gorm "github.com/orandin/slog-gorm"
	gorm_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"

	libconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
	liblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
)

func OpenPostgres(username, password, host string, port int, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, username, password, database, port, "disable", time.UTC.String())

	gormDialector := gorm_postgres.Open(dsn)

	gormConfig := gorm.Config{
		Logger: slog_gorm.New(
			slog_gorm.WithTraceAll(), // trace all messages
			slog_gorm.WithContextFunc(liblog.LoggerNameKey, func(ctx context.Context) (slog.Value, bool) {
				return slog.StringValue("gorm"), true
			}),
		),
	}

	return gorm.Open(gormDialector, &gormConfig)
}

func MigratePostgresDB(db *gorm.DB, sqlFS fs.FS) error {
	driverName := "postgres"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}

	return libgateway.MigrateDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_postgres.WithInstance(sqlDB, &migrate_postgres.Config{})
	})
}

func initPostgres(ctx context.Context, cfg *libconfig.DBConfig, fs fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := OpenPostgres(cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database)
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

	if err := MigratePostgresDB(db, fs); err != nil {
		return nil, nil, nil, liberrors.Errorf("failed to MigrateMySQLDB. err: %w", err)
	}

	dialect := libgateway.DialectPostgres{}
	return &dialect, db, sqlDB, nil
}
