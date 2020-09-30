package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type File struct {
	templates.TID
	templates.TTimestamps
	Database bool
	Data     []byte
	Type     string
}
