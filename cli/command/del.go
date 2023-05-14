package command

import (
	"fmt"

	"github.com/0rcastra/Orca/internal/data"
)

type DelCommand struct {
	Database *data.Database
}

func (c *DelCommand) Name() string {
	return "del"
}

func (c *DelCommand) Execute(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments for DEL command")
	}

	key := args[0]

	deleted := c.Database.Del(key)
	if deleted {
		fmt.Printf("Key '%s' deleted\n", key)
	} else {
		fmt.Printf("Key '%s' not found\n", key)
	}

	return nil
}
