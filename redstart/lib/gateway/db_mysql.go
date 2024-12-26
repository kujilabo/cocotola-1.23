package gateway

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log/slog"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	slog_gorm "github.com/orandin/slog-gorm"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	liblog "github.com/kujilabo/cocotola-1.23/redstart/lib/log"
)

func OpenMySQL(username, password, host string, port int, database string) (*gorm.DB, error) {
	c := mysql.Config{
		DBName:               database,
		User:                 username,
		Passwd:               password,
		Addr:                 fmt.Sprintf("%s:%d", host, port),
		Net:                  "tcp",
		ParseTime:            true,
		MultiStatements:      true,
		Params:               map[string]string{"charset": "utf8mb4"},
		Collation:            "utf8mb4_bin",
		AllowNativePasswords: true,
		Loc:                  time.UTC,
	}
	dsn := c.FormatDSN()

	gormDialector := gorm_mysql.Open(dsn)

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

func MigrateMySQLDB(db *gorm.DB, sqlFS fs.FS) error {
	driverName := "mysql"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}

	return migrateDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_mysql.WithInstance(sqlDB, &migrate_mysql.Config{})
	})
}
