package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type CollectionContent struct {
	templates.TID
	templates.TTimestamps
	CollectionID, CompositionID int
	Index                       int
	CollectionCompositionPart   []CollectionCompositionPart `gorm:"many2many:collection_content_collection_composition_parts;"`
}
