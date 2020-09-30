package migration20201001001

import (
	"github.com/ski7777/MusicNotesManager/internal/database/migration/migrations/20201001001/models"
	"gorm.io/gorm"
)

func Migrate(d *gorm.DB) error {
	return d.AutoMigrate(
		&migration20201001001m.Permission{},
		&migration20201001001m.Part{},
		&migration20201001001m.File{},
		&migration20201001001m.Group{},
		&migration20201001001m.CompositionPart{},
		&migration20201001001m.User{},
		&migration20201001001m.Upload{},
		&migration20201001001m.Session{},
		&migration20201001001m.Page{},
		&migration20201001001m.Overlay{},
		&migration20201001001m.CollectionPart{},
		&migration20201001001m.Collection{},
		&migration20201001001m.Composition{},
		&migration20201001001m.CompositionPartPage{},
		&migration20201001001m.CollectionCompositionPart{},
		&migration20201001001m.CollectionContent{},
	)
}
