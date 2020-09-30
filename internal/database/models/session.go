package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type Session struct {
	templates.TID
	templates.TTimestamps
	UserID int
}
