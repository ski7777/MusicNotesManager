package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type CompositionPartPage struct {
	templates.TID
	templates.TTimestamps
	CompositionPartID int
	PageID            int
	Index             int
}
