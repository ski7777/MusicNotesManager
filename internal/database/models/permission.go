package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type Permission struct {
	templates.TID
	Name, FriendlyName string
}
