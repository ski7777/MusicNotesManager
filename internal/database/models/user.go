package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type User struct {
	templates.TID
	templates.TTimestamps
	FirstName, LastName, Mail, Password string
	Groups                              []Group          `gorm:"many2many:user_groups;"`
	Permissions                         []Permission     `gorm:"many2many:user_permissions;"`
	CollectionParts                     []CollectionPart `gorm:"many2many:user_collection_parts;"`
}
