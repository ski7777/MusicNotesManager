package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type Part struct {
	templates.TID
	templates.TTimestamps
	Name string
}
