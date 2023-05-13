package command

import (
	"fmt"

	"github.com/0rcastra/Orca/internal/data"
)

type GetCommand struct {
	Database *data.Database
}

func (c *GetCommand) Name() string {
	return "get"
}

func (c *GetCommand) Execute(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments for GET command")
	}

	key := args[0]

	value, exists := data.Get(c.Database, key)
	if !exists {
		return fmt.Errorf("key %s not found", key)
	}

	fmt.Printf("%s\n", value)

	return nil
}
