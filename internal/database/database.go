package database

import (
	"errors"
	"github.com/ski7777/MusicNotesManager/internal/database/migration"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// ts == thread-safe
var drivers = map[string]struct {
	f  func(string) gorm.Dialector
	ts bool
}{
	"mysql":    {f: mysql.Open, ts: true},
	"sqlite":   {f: sqlite.Open, ts: false},
	"postgres": {f: postgres.Open, ts: true},
	"mssql":    {f: sqlserver.Open, ts: true},
}

func GetDB(driver, dsn string) (*gorm.DB, bool, error) {
	d, ok := drivers[driver]
	if !ok {
		return nil, false, errors.New("unknown database driver")
	}
	db, err := gorm.Open(d.f(dsn), nil)
	if err != nil {
		return nil, false, err
	}
	if err = migration.Migrate(db); err != nil {
		return nil, false, err
	}
	return db, d.ts, nil
}
