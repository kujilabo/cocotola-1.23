package gateway

type DialectRDBMS interface {
	Name() string
	BoolDefaultValue() string
}

type DialectMySQL struct {
}

func (d *DialectMySQL) Name() string {
	return "mysql"
}

func (d *DialectMySQL) BoolDefaultValue() string {
	return "0"
}

type DialectPostgres struct {
}

func (d *DialectPostgres) Name() string {
	return "postgres"
}

func (d *DialectPostgres) BoolDefaultValue() string {
	return "false"
}

type DialectSQLite3 struct {
}

func (d *DialectSQLite3) Name() string {
	return "sqlite3"
}

func (d *DialectSQLite3) BoolDefaultValue() string {
	return "0"
}
