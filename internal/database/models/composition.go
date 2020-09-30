package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type Composition struct {
	templates.TID
	templates.TTimestamps
	Name, Composer   string
	CompositionParts []CompositionPart `gorm:"many2many:composition_composition_parts;"`
}
