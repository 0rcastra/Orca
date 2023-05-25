package data_test

import (
	"testing"

	"github.com/0rcastra/Orca/internal/data"
)

func TestExists(t *testing.T) {
	db := data.NewDatabase()

	// Test case 1: Key exists
	db.Set("name", "Seongbin")

	exists := db.Exists("name")
	if !exists {
		t.Errorf("unexpected result: key 'name' should exist")
	}

	// Test case 2: Key does not exist
	exists = db.Exists("nonexistent")
	if exists {
		t.Errorf("unexpected result: key 'nonexistent' should not exist")
	}
}
