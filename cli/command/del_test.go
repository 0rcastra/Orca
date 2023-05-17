package command_test

import (
	"testing"

	"github.com/0rcastra/Orca/cli/command"
	"github.com/0rcastra/Orca/internal/data"
)

func TestDelCommand_Execute(t *testing.T) {
	db := data.NewDatabase()
	setCmd := &command.SetCommand{Database: db}
	setErr := setCmd.Execute([]string{"name", "Seongbin"})
	if setErr != nil {
		t.Fatalf("unexpected error setting value: %s", setErr.Error())
	}

	cmd := &command.DelCommand{Database: db}

	args := []string{"name"}
	err := cmd.Execute(args)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if value, exists := db.Get("name"); exists {
		t.Errorf("expected key to be deleted, but value still exists: %s", value)
	}

	args = []string{"nonexistent"}
	err = cmd.Execute(args)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestDelCommand_Name(t *testing.T) {
	cmd := &command.DelCommand{}
	expectedName := "del"

	name := cmd.Name()

	if name != expectedName {
		t.Errorf("unexpected command name: got %s, want %s", name, expectedName)
	}
}
