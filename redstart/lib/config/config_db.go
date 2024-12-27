package config

import (
	"io/fs"

	liberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
	// libgatewaysqlite3 "github.com/kujilabo/cocotola-1.23/redstart/lib/gateway/sqlite3"
)

type SQLite3Config struct {
	File string `yaml:"file" validate:"required"`
}

type MySQLConfig struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	Database string `yaml:"database" validate:"required"`
}

type PostgresConfig struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	Database string `yaml:"database" validate:"required"`
}

type DBConfig struct {
	DriverName string          `yaml:"driverName"`
	SQLite3    *SQLite3Config  `yaml:"sqlite3"`
	MySQL      *MySQLConfig    `yaml:"mysql"`
	Postgres   *PostgresConfig `yaml:"postgres"`
	Migration  bool            `yaml:"migration"`
}

type mergedFS struct {
	fss     []fs.FS
	entries []fs.DirEntry
}

func MergeFS(driverName string, fss ...fs.FS) (*mergedFS, error) {
	entries := make([]fs.DirEntry, 0)
	for i := range fss {
		e, err := fs.ReadDir(fss[i], driverName)
		if err != nil {
			return nil, liberrors.Errorf("read %q directory: %w", driverName, err)
		}
		entries = append(entries, e...)
	}

	return &mergedFS{
		fss:     fss,
		entries: entries,
	}, nil
}

func (f *mergedFS) Open(name string) (fs.File, error) {
	var file fs.File
	var err error
	for i := range f.fss {
		file, err = f.fss[i].Open(name)
		if err == nil {
			return file, nil
		}
	}

	return nil, err
}

func (f *mergedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return f.entries, nil
}
