package migration20201001001m

import (
	"time"
)

type Upload struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Processed bool
	FileID    int
	UserID    int
}
