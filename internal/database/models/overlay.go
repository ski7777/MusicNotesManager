package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type Overlay struct {
	templates.TID
	templates.TTimestamps
	UserID int
	PageID int
	FileID int
}
