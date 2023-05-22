package data_test

import (
	"testing"

	"github.com/0rcastra/Orca/internal/data"
)

func TestDel(t *testing.T) {
	db := data.NewDatabase()

	db.Set("name", "Seongbin")
	exists := db.Del("name")
	if !exists {
		t.Errorf("Del failed: expected true, got false")
	}

	_, exists = db.Get("name")
	if exists {
		t.Errorf("Del failed: key 'name' still exists")
	}

	exists = db.Del("nonexistent")
	if exists {
		t.Errorf("Del failed: expected false, got true")
	}
}
