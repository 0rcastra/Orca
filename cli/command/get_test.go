package command_test

import (
	"strings"
	"testing"

	"github.com/0rcastra/Orca/cli/command"
	"github.com/0rcastra/Orca/internal/data"
	"github.com/0rcastra/Orca/utils"
)

func TestGetCommand_Execute(t *testing.T) {
	db := data.NewDatabase()
	setCmd := &command.SetCommand{Database: db}
	setErr := setCmd.Execute([]string{"name", "Seongbin"})
	if setErr != nil {
		t.Fatalf("unexpected error setting value: %s", setErr.Error())
	}

	cmd := &command.GetCommand{Database: db}

	expectedOutput := "Seongbin"
	output, err := utils.CaptureOutput(func() {
		cmdErr := cmd.Execute([]string{"name"})
		if cmdErr != nil {
			t.Errorf("unexpected error: %s", cmdErr.Error())
		}
	})
	if err != nil {
		t.Fatalf("failed to capture output: %s", err.Error())
	}
	if strings.TrimSpace(output) != expectedOutput {
		t.Errorf("unexpected output: got %q, want %q", strings.TrimSpace(output), expectedOutput)
	}

	expectedErrorMessage := "invalid number of arguments for GET command"
	output, err = utils.CaptureOutput(func() {
		cmdErr := cmd.Execute([]string{})
		if cmdErr == nil || cmdErr.Error() != expectedErrorMessage {
			t.Errorf("unexpected error: got %v, want %s", cmdErr, expectedErrorMessage)
		}
	})
	if err != nil {
		t.Fatalf("failed to capture output: %s", err.Error())
	}

	expectedErrorMessage = "key nonexistent not found"
	output, err = utils.CaptureOutput(func() {
		cmdErr := cmd.Execute([]string{"nonexistent"})
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
	if output != "" {
		t.Errorf("unexpected output: got %q, want empty string", output)
	}
}

func TestGetCommand_Name(t *testing.T) {
	cmd := &command.GetCommand{}
	expectedName := "get"

	name := cmd.Name()

	if name != expectedName {
		t.Errorf("unexpected command name: got %s, want %s", name, expectedName)
	}
}
