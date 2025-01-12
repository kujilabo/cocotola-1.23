package config

import (
	"context"
	"database/sql"
	"io/fs"

	"gorm.io/gorm"

	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
)

func initDBSQLite3(ctx context.Context, cfg *DBConfig, fs fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	return libgateway.InitSqlite3(ctx, cfg.SQLite3, cfg.Migration, fs)
}
