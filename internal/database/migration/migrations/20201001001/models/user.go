package migration20201001001m

import (
	"time"
)

type User struct {
	ID                                  int `gorm:"primaryKey"`
	CreatedAt                           time.Time
	UpdatedAt                           time.Time
	FirstName, LastName, Mail, Password string
	Groups                              []Group          `gorm:"many2many:user_groups;"`
	Permissions                         []Permission     `gorm:"many2many:user_permissions;"`
	CollectionParts                     []CollectionPart `gorm:"many2many:user_collection_parts;"`
}
