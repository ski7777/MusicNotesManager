package migration20201001001m

import (
	"time"
)

type CollectionCompositionPart struct {
	ID                                  int `gorm:"primaryKey"`
	CreatedAt                           time.Time
	UpdatedAt                           time.Time
	CollectionPartID, CompositionPartID int
}
