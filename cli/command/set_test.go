package command_test

import (
	"testing"

	"github.com/0rcastra/Orca/cli/command"
)

func TestSetCommand_Execute(t *testing.T) {
	cmd := &command.SetCommand{}

	args := []string{"key1", "value1"}
	err := cmd.Execute(args)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	args = []string{"key2"}
	err = cmd.Execute(args)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
