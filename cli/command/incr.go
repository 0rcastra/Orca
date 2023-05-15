package command

import (
	"fmt"

	"github.com/0rcastra/Orca/internal/data"
)

type IncrCommand struct {
	Database *data.Database
}

func (c *IncrCommand) Name() string {
	return "incr"
}

func (c *IncrCommand) Execute(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments for INCR command")
	}

	key := args[0]

	newValue, err := c.Database.Incr(key)
	if err != nil {
		return fmt.Errorf("failed to increment value for key %s: %w", key, err)
	}

	fmt.Printf("New Value: %d\n", newValue)

	return nil
}
