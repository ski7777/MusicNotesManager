package migration20201001001m

import (
	"time"
)

type Page struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UploadID  int
	PageNo    int
	Rotation  int
}
