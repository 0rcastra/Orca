package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	commands map[string]Command
}

type Command interface {
	Name() string
	Execute(args []string) error
}

// To creates a new CLI instance.
func NewCLI() *CLI {
	return &CLI{
		commands: make(map[string]Command),
	}
}

// To registers a command with the CLI.
func (cli *CLI) RegisterCommand(command Command) {
	cli.commands[command.Name()] = command
}

// starts the CLI.
func (cli *CLI) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Orca> ")
		if !scanner.Scan() {
			return
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "exit" {
			return
		}

		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			continue
		}

		commandName := parts[0]
		commandArgs := parts[1:]

		command, exists := cli.commands[commandName]
		if !exists {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}

		if err := command.Execute(commandArgs); err != nil {
			fmt.Printf("Error executing command: %s\n", err)
		}
	}
}
