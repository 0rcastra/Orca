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

	args2 := []string{}
	err2 := cmd.Execute(args2)
	expectedErrorMsg2 := "invalid number of arguments for DEL command"
	if err2 == nil || err2.Error() != expectedErrorMsg2 {
		t.Errorf("unexpected error: got %v, want %s", err2, expectedErrorMsg2)
	}

	args3 := []string{"key1", "key2"}
	err3 := cmd.Execute(args3)
	expectedErrorMsg3 := "invalid number of arguments for DEL command"
	if err3 == nil || err3.Error() != expectedErrorMsg3 {
		t.Errorf("unexpected error: got %v, want %s", err3, expectedErrorMsg3)
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
