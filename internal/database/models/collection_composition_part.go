package models

import "github.com/ski7777/MusicNotesManager/internal/database/models/templates"

type CollectionCompositionPart struct {
	templates.TID
	templates.TTimestamps
	CollectionPartID, CompositionPartID int
}
