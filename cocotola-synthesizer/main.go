package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	nurl "net/url"
	"os"
	"strings"

	"go.uber.org/atomic"

	"gorm.io/gorm"

	"github.com/hashicorp/go-multierror"

	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"

	rslibconfig "github.com/kujilabo/cocotola-1.23/redstart/lib/config"
	rslibgateway "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway"

	// rslibgatewaysqlite3 "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway/sqlite3"

	// migrate_sqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	// "github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/config"
	// "github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/sqls"
)

func getValue(values ...string) string {
	for _, v := range values {
		if len(v) != 0 {
			return v
		}
	}
	return ""
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}

type Config struct {
	MigrationsTable string
	DatabaseName    string
	NoTxWrap        bool
}

type Sqlite struct {
	db       *sql.DB
	isLocked atomic.Bool

	config *Config
}

func main() {
	var db *gorm.DB
	log.Println("Hello, World!!!!!!!!")
	database.Register("sqlite", &Sqlite{})
	iofs.New(nil, "")
	var _ = db

	// var _ = migrate_sqlite3.Config{}

	ctx := context.Background()
	env := flag.String("env", "", "environment")
	flag.Parse()
	appEnv := getValue(*env, os.Getenv("APP_ENV"), "local")

	rsliberrors.UseXerrorsErrorf()

	// load config
	// cfg, err := config.LoadConfig(appEnv)
	// checkError(err)
	// var _ = cfg

	var _ = rslibgateway.MYSQL_ER_DUP_ENTRY
	var _ = rslibconfig.DBConfig{}
	// var _ = ctx
	// var _ = cfg
	var _ = appEnv
	var _ = ctx

	// dialect, db, sqlDB, err := rslibconfig.InitDB(ctx, cfg.DB, map[string]rslibconfig.DBInitializer{
	// 	"mysql": rslibgatewaysqlite3.InitSqlite3,
	// }, sqls.SQL)
	// checkError(err)
	// var _ = dialect
	// var _ = sqlDB

	// // init log
	// rslibconfig.InitLog(cfg.Log)
	// logger := slog.Default().With(slog.String(rsliblog.LoggerNameKey, "main"))
	// logger.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	// // init tracer
	// tp, err := rslibconfig.InitTracerProvider(ctx, cfg.App.Name, cfg.Trace)
	// checkError(err)
	// otel.SetTracerProvider(tp)
	// otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	// if config == nil {
	// 	return nil, ErrNilConfig
	// }

	// if err := instance.Ping(); err != nil {
	// 	return nil, err
	// }

	// if len(config.MigrationsTable) == 0 {
	// 	config.MigrationsTable = DefaultMigrationsTable
	// }

	mx := &Sqlite{
		db:     instance,
		config: config,
	}
	if err := mx.ensureVersionTable(); err != nil {
		return nil, err
	}
	return mx, nil
}

// ensureVersionTable checks if versions table exists and, if not, creates it.
// Note that this function locks the database, which deviates from the usual
// convention of "caller locks" in the Sqlite type.
func (m *Sqlite) ensureVersionTable() (err error) {
	if err = m.Lock(); err != nil {
		return err
	}

	defer func() {
		if e := m.Unlock(); e != nil {
			if err == nil {
				err = e
			} else {
				err = multierror.Append(err, e)
			}
		}
	}()

	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (version uint64,dirty bool);
  CREATE UNIQUE INDEX IF NOT EXISTS version_unique ON %s (version);
  `, m.config.MigrationsTable, m.config.MigrationsTable)

	if _, err := m.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (m *Sqlite) Open(url string) (database.Driver, error) {
	purl, err := nurl.Parse(url)
	if err != nil {
		return nil, err
	}
	dbfile := strings.Replace(migrate.FilterCustomQuery(purl).String(), "sqlite://", "", 1)
	db, err := sql.Open("sqlite", dbfile)
	if err != nil {
		return nil, err
	}
	var _ = db

	// qv := purl.Query()

	// migrationsTable := qv.Get("x-migrations-table")
	// if len(migrationsTable) == 0 {
	// 	migrationsTable = DefaultMigrationsTable
	// }

	// noTxWrap := false
	// if v := qv.Get("x-no-tx-wrap"); v != "" {
	// 	noTxWrap, err = strconv.ParseBool(v)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("x-no-tx-wrap: %s", err)
	// 	}
	// }
	return nil, nil

	// mx, err := WithInstance(db, &Config{
	// 	DatabaseName:    purl.Path,
	// 	MigrationsTable: migrationsTable,
	// 	NoTxWrap:        noTxWrap,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return mx, nil
}

func (m *Sqlite) Close() error {
	return m.db.Close()
}

func (m *Sqlite) Drop() (err error) {
	query := `SELECT name FROM sqlite_master WHERE type = 'table';`
	tables, err := m.db.Query(query)
	if err != nil {
		return &database.Error{OrigErr: err, Query: []byte(query)}
	}
	defer func() {
		if errClose := tables.Close(); errClose != nil {
			err = multierror.Append(err, errClose)
		}
	}()

	tableNames := make([]string, 0)
	for tables.Next() {
		var tableName string
		if err := tables.Scan(&tableName); err != nil {
			return err
		}
		if len(tableName) > 0 {
			tableNames = append(tableNames, tableName)
		}
	}
	if err := tables.Err(); err != nil {
		return &database.Error{OrigErr: err, Query: []byte(query)}
	}

	if len(tableNames) > 0 {
		for _, t := range tableNames {
			query := "DROP TABLE " + t
			err = m.executeQuery(query)
			if err != nil {
				return &database.Error{OrigErr: err, Query: []byte(query)}
			}
		}
		query := "VACUUM"
		_, err = m.db.Query(query)
		if err != nil {
			return &database.Error{OrigErr: err, Query: []byte(query)}
		}
	}

	return nil
}

func (m *Sqlite) Lock() error {
	if !m.isLocked.CAS(false, true) {
		return database.ErrLocked
	}
	return nil
}

func (m *Sqlite) Unlock() error {
	if !m.isLocked.CAS(true, false) {
		return database.ErrNotLocked
	}
	return nil
}

func (m *Sqlite) Run(migration io.Reader) error {
	migr, err := io.ReadAll(migration)
	if err != nil {
		return err
	}
	query := string(migr[:])

	if m.config.NoTxWrap {
		return m.executeQueryNoTx(query)
	}
	return m.executeQuery(query)
}

func (m *Sqlite) executeQuery(query string) error {
	tx, err := m.db.Begin()
	if err != nil {
		return &database.Error{OrigErr: err, Err: "transaction start failed"}
	}
	if _, err := tx.Exec(query); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			err = multierror.Append(err, errRollback)
		}
		return &database.Error{OrigErr: err, Query: []byte(query)}
	}
	if err := tx.Commit(); err != nil {
		return &database.Error{OrigErr: err, Err: "transaction commit failed"}
	}
	return nil
}

func (m *Sqlite) executeQueryNoTx(query string) error {
	if _, err := m.db.Exec(query); err != nil {
		return &database.Error{OrigErr: err, Query: []byte(query)}
	}
	return nil
}

func (m *Sqlite) SetVersion(version int, dirty bool) error {
	tx, err := m.db.Begin()
	if err != nil {
		return &database.Error{OrigErr: err, Err: "transaction start failed"}
	}

	query := "DELETE FROM " + m.config.MigrationsTable
	if _, err := tx.Exec(query); err != nil {
		return &database.Error{OrigErr: err, Query: []byte(query)}
	}

	// Also re-write the schema version for nil dirty versions to prevent
	// empty schema version for failed down migration on the first migration
	// See: https://github.com/golang-migrate/migrate/issues/330
	if version >= 0 || (version == database.NilVersion && dirty) {
		query := fmt.Sprintf(`INSERT INTO %s (version, dirty) VALUES (?, ?)`, m.config.MigrationsTable)
		if _, err := tx.Exec(query, version, dirty); err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				err = multierror.Append(err, errRollback)
			}
			return &database.Error{OrigErr: err, Query: []byte(query)}
		}
	}

	if err := tx.Commit(); err != nil {
		return &database.Error{OrigErr: err, Err: "transaction commit failed"}
	}

	return nil
}

func (m *Sqlite) Version() (version int, dirty bool, err error) {
	query := "SELECT version, dirty FROM " + m.config.MigrationsTable + " LIMIT 1"
	err = m.db.QueryRow(query).Scan(&version, &dirty)
	if err != nil {
		return database.NilVersion, false, nil
	}
	return version, dirty, nil
}
