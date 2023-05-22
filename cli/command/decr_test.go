package command_test

import (
	"strings"
	"testing"

	"github.com/0rcastra/Orca/cli/command"
	"github.com/0rcastra/Orca/internal/data"
	"github.com/0rcastra/Orca/utils"
)

func TestDecrCommand_Execute_Error(t *testing.T) {
	db := data.NewDatabase()
	setCmd := &command.SetCommand{Database: db}
	setErr := setCmd.Execute([]string{"name", "Seongbin"})
	if setErr != nil {
		t.Fatalf("unexpected error setting value: %s", setErr.Error())
	}

	cmd := &command.DecrCommand{Database: db}

	expectedErrorMessage := "failed to decrement value for key name: value for key 'name' is not a valid integer"
	_, err := utils.CaptureOutput(func() {
		cmdErr := cmd.Execute([]string{"name"})
		if cmdErr == nil {
			t.Errorf("expected error, got nil")
		}
		if !strings.Contains(cmdErr.Error(), expectedErrorMessage) {
			t.Errorf("unexpected error message: got %q, want %q", cmdErr.Error(), expectedErrorMessage)
		}
	})
	if err != nil {
		t.Fatalf("failed to capture output: %s", err.Error())
	}
}

func TestDecrCommand_Execute(t *testing.T) {
	db := data.NewDatabase()
	cmd := &command.DecrCommand{Database: db}

	key := "age"
	db.Set(key, "23")

	expectedOutput := "New Value: 22\n"
	output, err := utils.CaptureOutput(func() {
		err := cmd.Execute([]string{key})
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}
	})
	if err != nil {
		t.Fatalf("failed to capture output: %s", err.Error())
	}
	if output != expectedOutput {
		t.Errorf("unexpected output: got %q, want %q", output, expectedOutput)
	}

	args := []string{}
	err = cmd.Execute(args)
	expectedErrorMsg := "invalid number of arguments for DECR command"
	if err == nil || err.Error() != expectedErrorMsg {
		t.Errorf("unexpected error: got %v, want %s", err, expectedErrorMsg)
	}
}

func TestDecrCommand_Name(t *testing.T) {
	cmd := &command.DecrCommand{}
	expectedName := "decr"

	name := cmd.Name()

	if name != expectedName {
		t.Errorf("unexpected command name: got %s, want %s", name, expectedName)
	}
}
