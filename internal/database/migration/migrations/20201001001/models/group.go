package migration20201001001m

import (
	"time"
)

type Group struct {
	ID          int `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Permissions []Permission `gorm:"many2many:group_permissions;"`
}
