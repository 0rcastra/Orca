package data_test

import (
	"fmt"
	"testing"

	"github.com/0rcastra/Orca/internal/data"
)

func TestDecr(t *testing.T) {
	db := data.NewDatabase()

	newValue, err := db.Decr("age")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if newValue != 1 {
		t.Errorf("unexpected new value: got %d, want %d", newValue, 1)
	}

	db.Set("age", "23")
	existingValue, err := db.Decr("age")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if existingValue != 22 {
		t.Errorf("unexpected existing value: got %d, want %d", existingValue, 22)
	}

	db.Set("name", "Seongbin")
	_, err = db.Decr("name")
	if err == nil || err.Error() != fmt.Sprintf("value for key 'name' is not a valid integer: strconv.Atoi: parsing \"Seongbin\": invalid syntax") {
		t.Errorf("expected error: value for key 'name' is not a valid integer")
	}
}
