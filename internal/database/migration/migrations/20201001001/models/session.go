package migration20201001001m

import (
	"time"
)

type Session struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
}
