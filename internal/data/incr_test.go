package data_test

import (
	"fmt"
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

	db.Set("invalid", "abc")
	_, err = db.Incr("invalid")
	if err == nil || err.Error() != fmt.Sprintf("value for key 'invalid' is not a valid integer: strconv.Atoi: parsing \"abc\": invalid syntax") {
		t.Errorf("expected error: value for key 'invalid' is not a valid integer")
	}
}
