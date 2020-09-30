package migration20201001001m

import (
	"time"
)

type CollectionContent struct {
	ID                          int `gorm:"primaryKey"`
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
	CollectionID, CompositionID int
	Index                       int
	CollectionCompositionPart   []CollectionCompositionPart `gorm:"many2many:collection_content_collection_composition_parts;"`
}
