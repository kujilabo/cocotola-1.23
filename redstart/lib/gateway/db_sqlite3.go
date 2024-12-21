package gateway

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"

	"github.com/golang-migrate/migrate/v4/database"
	migrate_sqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	liblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
	slog_gorm "github.com/orandin/slog-gorm"
	gorm_sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSQLite3(filePath string, logger *slog.Logger) (*gorm.DB, error) {
	return gorm.Open(gorm_sqlite.Open(filePath), &gorm.Config{
		Logger: slog_gorm.New(
			slog_gorm.WithTraceAll(), // trace all messages
			slog_gorm.WithContextFunc(liblog.LoggerNameKey, func(ctx context.Context) (slog.Value, bool) {
				return slog.StringValue("gorm"), true
			}),
		),
	})
}

func MigrateSQLite3DB(db *gorm.DB, sqlFS fs.FS) error {
	driverName := "sqlite3"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}

	return migrateDB(db, "sqlite3", sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_sqlite3.WithInstance(sqlDB, &migrate_sqlite3.Config{})
	})
}
