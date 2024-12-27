package gateway

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4/database"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/gorm"

	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
)

var testDBHost string
var testDBPort int

func openMySQLForTest() (*gorm.DB, error) {
	return libgateway.OpenMySQL(&libgateway.MySQLConfig{
		Username: "user",
		Password: "password",
		Host:     testDBHost,
		Port:     testDBPort,
		Database: "testdb",
	})
}

func setupMySQL(sqlFS embed.FS, db *gorm.DB) error {
	driverName := "mysql"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}

	return setupDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_mysql.WithInstance(sqlDB, &migrate_mysql.Config{})
	})
}

func InitMySQL(sqlFS embed.FS, dbHost string, dbPort int) (*gorm.DB, error) {
	testDBHost = dbHost
	testDBPort = dbPort
	db, err := openMySQLForTest()
	if err != nil {
		return nil, err
	}

	if err := setupMySQL(sqlFS, db); err != nil {
		return nil, err
	}

	return db, nil
}
