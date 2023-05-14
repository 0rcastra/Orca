package main

import (
	"github.com/0rcastra/Orca/cli"
	"github.com/0rcastra/Orca/cli/command"
	"github.com/0rcastra/Orca/internal/data"
	"github.com/inancgumus/screen"
)

func main() {
	db := data.NewDatabase()

	cli := cli.NewCLI()

	setCommand := &command.SetCommand{
		Database: db,
	}
	cli.RegisterCommand(setCommand)

	getCommand := &command.GetCommand{
		Database: db,
	}
	cli.RegisterCommand(getCommand)

	clearScreen()

	cli.Run()
}

func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}