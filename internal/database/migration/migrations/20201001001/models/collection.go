package migration20201001001m

import (
	"time"
)

type Collection struct {
	ID              int `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string
	CollectionParts []CollectionPart `gorm:"many2many:collection_collection_parts;"`
}
