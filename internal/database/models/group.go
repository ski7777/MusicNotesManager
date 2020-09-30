package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type Group struct {
	templates.TID
	templates.TTimestamps
	Name        string
	Permissions []Permission `gorm:"many2many:group_permissions;"`
}
