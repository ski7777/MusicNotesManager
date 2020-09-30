package migration20201001001m

import (
	"time"
)

type CompositionPart struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	PartID    int
}
