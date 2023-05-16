package data_test

import (
	"testing"

	"github.com/0rcastra/Orca/internal/data"
)

func TestSet(t *testing.T) {
	db := data.NewDatabase()

	db.Set("name", "seongbin")
	if value, exists := db.Get("name"); !exists || value != "seongbin" {
		t.Errorf("Set failed: expected value 'seongbin', got '%s'", value)
	}
	db.Set("age", "23")

	if value, exists := db.Get("age"); !exists || value != "23" {
		t.Errorf("Set failed: expected value '23', got '%s'", value)
	}

	db.Set("", "chobobdev")
	if value, exists := db.Get(""); !exists || value != "chobobdev" {
		t.Errorf("Set failed: expected value 'chobobdev', got '%s'", value)
	}
}
