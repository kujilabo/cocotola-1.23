package config

import (
	"context"
	"database/sql"
	"io/fs"

	"gorm.io/gorm"

	libgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"
)

func initDBPostgres(ctx context.Context, cfg *DBConfig, fs fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	return libgateway.InitPostgres(ctx, cfg.Postgres, cfg.Migration, fs)
}
