package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type Collection struct {
	templates.TID
	templates.TTimestamps
	Name            string
	CollectionParts []CollectionPart `gorm:"many2many:collection_collection_parts;"`
}
