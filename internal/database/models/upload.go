package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type Upload struct {
	templates.TID
	templates.TCreatedAt
	Processed bool
	FileID    int
	UserID    int
}
