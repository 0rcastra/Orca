package main

import (
	"github.com/0rcastra/Orca/cli"
	"github.com/0rcastra/Orca/cli/command"
)

func main() {
	cli := cli.NewCLI()

	cli.RegisterCommand(&command.SetCommand{})

	cli.Run()
}
