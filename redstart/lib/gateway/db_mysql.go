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

	libconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"

	// libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
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

	return MigrateDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_mysql.WithInstance(sqlDB, &migrate_mysql.Config{})
	})
}

func InitMySQL(ctx context.Context, cfg *libconfig.DBConfig, fs fs.FS) (DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := OpenMySQL(cfg.MySQL.Username, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Database)
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

	if err := MigrateMySQLDB(db, fs); err != nil {
		return nil, nil, nil, liberrors.Errorf("failed to MigrateMySQLDB. err: %w", err)
	}

	dialect := DialectMySQL{}
	return &dialect, db, sqlDB, nil
}
