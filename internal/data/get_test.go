package data_test

import (
	"testing"

	"github.com/0rcastra/Orca/internal/data"
)

func TestGet(t *testing.T) {
	db := data.NewDatabase()

	db.Set("name", "Seongbin")

	value, exists := db.Get("name")
	if !exists {
		t.Errorf("expected key to exist, but it doesn't")
	}

	if value != "Seongbin" {
		t.Errorf("unexpected value: got %s, want %s", value, "Seongbin")
	}

	value, exists = db.Get("nonexistent")
	if exists {
		t.Errorf("expected key to not exist, but it exists")
	}

	if value != "" {
		t.Errorf("unexpected value: got %s, want empty string", value)
	}
}
