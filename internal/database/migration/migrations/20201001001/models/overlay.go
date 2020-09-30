package migration20201001001m

import (
	"time"
)

type Overlay struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	PageID    int
	FileID    int
}
