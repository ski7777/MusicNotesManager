package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type CompositionPart struct {
	templates.TID
	templates.TTimestamps
	PartID int
}
