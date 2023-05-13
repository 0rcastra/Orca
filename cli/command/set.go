package command

import (
	"fmt"

	"github.com/0rcastra/Orca/internal/data"
)

type SetCommand struct {
	Database *data.Database
}

func (c *SetCommand) Name() string {
	return "set"
}

func (c *SetCommand) Execute(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("invalid number of arguments for SET command")
	}

	key := args[0]
	value := args[1]

	c.Database.Set(key, value)

	fmt.Printf("OK\n")

	return nil
}
