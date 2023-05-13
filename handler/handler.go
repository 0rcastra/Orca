package handler

import (
	"github.com/0rcastra/Orca/internal/data"
)

// Handler handles HTTP requests.
type Handler struct {
	db *data.Database
}

// NewHandler creates a new instance of the Handler.
func NewHandler(db *data.Database) *Handler {
	return &Handler{
		db: db,
	}
}
