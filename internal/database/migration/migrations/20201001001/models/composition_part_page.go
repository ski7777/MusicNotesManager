package migration20201001001m

import (
	"time"
)

type CompositionPartPage struct {
	ID                int `gorm:"primaryKey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CompositionPartID int
	PageID            int
	Index             int
}
