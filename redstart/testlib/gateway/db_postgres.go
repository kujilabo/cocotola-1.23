package gateway

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-migrate/migrate/v4/database"
	migrate_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	slog_gorm "github.com/orandin/slog-gorm"
	gorm_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"

	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
)

var testPostgresHost string
var testPostgresPort int

func openPostgresForTest() (*gorm.DB, error) {
	logger := slog.Default()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", testPostgresHost, "user", "password", "postgres", testPostgresPort, "disable", time.UTC.String())
	db, err := gorm.Open(gorm_postgres.Open(dsn), &gorm.Config{
		Logger: slog_gorm.New(
			slog_gorm.WithHandler(logger.Handler()),
			// slog_gorm.WithTraceAll(),     // trace all messages
		),
	})
	if err != nil {
		return nil, liberrors.Errorf("gorm.Open. err: %w", err)
	}
	return db, nil
}

func setupPostgres(sqlFS embed.FS, db *gorm.DB) error {
	driverName := "postgres"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}

	return SetupDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_postgres.WithInstance(sqlDB, &migrate_postgres.Config{})
	})
}

func InitPostgres(sqlFS embed.FS, dbHost string, dbPort int) (*gorm.DB, error) {
	testPostgresHost = dbHost
	testPostgresPort = dbPort
	db, err := openPostgresForTest()
	if err != nil {
		return nil, err
	}

	if err := setupPostgres(sqlFS, db); err != nil {
		return nil, err
	}

	return db, nil
}
