package command_test

import (
	"testing"

	"github.com/0rcastra/Orca/cli/command"
	"github.com/0rcastra/Orca/internal/data"
	"github.com/0rcastra/Orca/utils"
)

func TestIncrCommand_Execute(t *testing.T) {
	db := data.NewDatabase()
	cmd := &command.IncrCommand{Database: db}

	// Test case 1: Increment existing key
	key := "count"
	db.Set(key, "5")

	expectedOutput := "New Value: 6\n"
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

	// Test case 2: Invalid number of arguments
	args := []string{}
	err = cmd.Execute(args)
	expectedErrorMsg := "invalid number of arguments for INCR command"
	if err == nil || err.Error() != expectedErrorMsg {
		t.Errorf("unexpected error: got %v, want %s", err, expectedErrorMsg)
	}
}
