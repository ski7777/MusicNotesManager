package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/ski7777/MusicNotesManager/internal/database/migration/migrations/20201001001"
	"gorm.io/gorm"
	"log"
)

var migrations = []*gormigrate.Migration{
	{
		ID: "20201001001",
		Migrate: func(db *gorm.DB) error {
			log.Println("Migrating", 20201001001)
			return migration20201001001.Migrate(db)
		},
	},
}
