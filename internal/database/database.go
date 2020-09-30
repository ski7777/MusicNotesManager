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

var drivers = map[string]func(string) gorm.Dialector{
	"mysql":    mysql.Open,
	"sqlite":   sqlite.Open,
	"postgres": postgres.Open,
	"mssql":    sqlserver.Open,
}

func GetDB(driver, dsn string) (*gorm.DB, error) {
	d, ok := drivers[driver]
	if !ok {
		return nil, errors.New("unknown database driver")
	}
	db, err := gorm.Open(d(dsn), nil)
	if err != nil {
		return nil, err
	}
	if err = migration.Migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}
