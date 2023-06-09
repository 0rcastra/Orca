package command_test

import (
	"testing"

	"github.com/0rcastra/Orca/cli/command"
	"github.com/0rcastra/Orca/internal/data"
)

func TestSetCommand_Execute(t *testing.T) {
	db := data.NewDatabase()

	cmd := &command.SetCommand{
		Database: db,
	}

	args := []string{"name", "seongbin"}
	err := cmd.Execute(args)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	args = []string{"age"}
	err = cmd.Execute(args)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}

func TestSetCommand_Name(t *testing.T) {
	cmd := &command.SetCommand{}
	expectedName := "set"

	name := cmd.Name()

	if name != expectedName {
		t.Errorf("unexpected command name: got %s, want %s", name, expectedName)
	}
}
