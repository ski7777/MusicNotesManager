package migration20201001001m

import (
	"time"
)

type Composition struct {
	ID               int `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Name, Composer   string
	CompositionParts []CompositionPart `gorm:"many2many:composition_composition_parts;"`
}
