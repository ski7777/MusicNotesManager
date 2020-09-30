package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type Page struct {
	templates.TID
	templates.TTimestamps
	UploadID int
	PageNo   int
	Rotation int
}
