package sqls

import (
	"embed"
	_ "embed"
)

//go:embed sqlite3/*.sql
var SQL embed.FS
