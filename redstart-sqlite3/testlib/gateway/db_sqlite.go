package gateway

import (
	"database/sql"
	"embed"
	"os"

	"github.com/golang-migrate/migrate/v4/database"
	migrate_sqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/gorm"

	rssqlite3libgateway "github.com/kujilabo/cocotola-1.23/redstart-sqlite3/lib/gateway"
	rstestlibgateway "github.com/kujilabo/cocotola-1.23/redstart/testlib/gateway"
)

// migrate_sqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite3"
// _ "github.com/mattn/go-sqlite3"

var testDBFile string

func openSQLiteForTest() (*gorm.DB, error) {
	return rssqlite3libgateway.OpenSQLite3(testDBFile)
	// logger := slog.Default()
	// return gorm.Open(gormSQLite.Open(testDBFile), &gorm.Config{
	// 	Logger: slog_gorm.New(
	// 		slog_gorm.WithHandler(logger.Handler()),
	// 		slog_gorm.WithTraceAll(), // trace all messages
	// 	),
	// })
}

// func OpenSQLiteInMemory(sqlFS embed.FS) (*gorm.DB, error) {
// 	logger := slog.Default()
// 	db, err := gorm.Open(gormSQLite.Open("file:memdb1?mode=memory&cache=shared"), &gorm.Config{
// 		Logger: slog_gorm.New(
// 			slog_gorm.WithLogger(logger), // Optional, use slog.Default() by default
// 			slog_gorm.WithTraceAll(),     // trace all messages
// 		),
// 	})
// 	if err != nil {
// 		return nil, liberrors.Errorf("gorm.Open. err: %w", err)
// 	}
// 	if err := setupSQLite(sqlFS, db); err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func setupSQLite(sqlFS embed.FS, db *gorm.DB) error {
	driverName := "sqlite3"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}
	return rstestlibgateway.SetupDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_sqlite3.WithInstance(sqlDB, &migrate_sqlite3.Config{})
	})
}

func InitSQLiteInFile(sqlFS embed.FS) (*gorm.DB, error) {
	testDBFile = "./test.db"
	os.Remove(testDBFile)
	db, err := openSQLiteForTest()
	if err != nil {
		return nil, err
	}
	if err := setupSQLite(sqlFS, db); err != nil {
		return nil, err
	}
	return db, nil
}
