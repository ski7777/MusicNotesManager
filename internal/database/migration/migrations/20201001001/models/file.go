package migration20201001001m

import (
	"time"
)

type File struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Database  bool
	Data      []byte
	Type      string
}
