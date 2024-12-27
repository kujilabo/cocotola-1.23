package gateway

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"

	gorm_sqlite "github.com/glebarez/sqlite"
	"github.com/golang-migrate/migrate/v4/database"
	migrate_sqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	slog_gorm "github.com/orandin/slog-gorm"
	"gorm.io/gorm"

	libconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
	liblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
)

// migrate_sqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite3"
// gorm_sqlite "gorm.io/driver/sqlite"

func OpenSQLite3(filePath string) (*gorm.DB, error) {
	gormDialector := gorm_sqlite.Open(filePath)

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

func MigrateSQLite3DB(db *gorm.DB, sqlFS fs.FS) error {
	driverName := "sqlite3"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}

	return libgateway.MigrateDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_sqlite3.WithInstance(sqlDB, &migrate_sqlite3.Config{})
	})
}

func InitSqlite3(ctx context.Context, cfg *libconfig.DBConfig, fs fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := OpenSQLite3("./" + cfg.SQLite3.File)
	if err != nil {
		return nil, nil, nil, liberrors.Errorf("OpenSQLite file: %s err: %w", cfg.SQLite3.File, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, nil, liberrors.Errorf("DB. file: %s err: %w", cfg.SQLite3.File, err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, nil, nil, liberrors.Errorf("Ping. file: %s err: %w", cfg.SQLite3.File, err)
	}

	if cfg.Migration {
		if err := MigrateSQLite3DB(db, fs); err != nil {
			return nil, nil, nil, liberrors.Errorf("migrate DB. file: %s err: %w", cfg.SQLite3.File, err)
		}
	}

	dialect := libgateway.DialectMySQL{}
	return &dialect, db, sqlDB, nil
}
