package data_test

import (
	"testing"

	"github.com/0rcastra/Orca/internal/data"
)

func TestIncr(t *testing.T) {
	db := data.NewDatabase()

	newValue, err := db.Incr("count")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if newValue != 1 {
		t.Errorf("unexpected new value: got %d, want %d", newValue, 1)
	}

	db.Set("count", "5")
	existingValue, err := db.Incr("count")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if existingValue != 6 {
		t.Errorf("unexpected existing value: got %d, want %d", existingValue, 6)
	}
}
