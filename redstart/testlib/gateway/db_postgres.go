package gateway

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4/database"
	migrate_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/gorm"

	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
)

var testPostgresHost string
var testPostgresPort int

func openPostgresForTest() (*gorm.DB, error) {
	return libgateway.OpenPostgres(&libgateway.PostgresConfig{
		Username: "user",
		Password: "password",
		Host:     testPostgresHost,
		Port:     testPostgresPort,
		Database: "postgres",
	})
}

func setupPostgres(sqlFS embed.FS, db *gorm.DB) error {
	driverName := "postgres"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}

	return setupDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
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
