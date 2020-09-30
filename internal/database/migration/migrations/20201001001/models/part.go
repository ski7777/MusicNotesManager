package migration20201001001m

import (
	"time"
)

type Part struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}
