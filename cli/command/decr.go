package command

import (
	"fmt"

	"github.com/0rcastra/Orca/internal/data"
)

type DecrCommand struct {
	Database *data.Database
}

func (c *DecrCommand) Name() string {
	return "decr"
}

func (c *DecrCommand) Execute(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments for DECR command")
	}

	key := args[0]

	newValue, err := c.Database.Decr(key)
	if err != nil {
		return fmt.Errorf("failed to decrement value for key %s: %w", key, err)
	}

	fmt.Printf("New Value: %d\n", newValue)

	return nil
}
