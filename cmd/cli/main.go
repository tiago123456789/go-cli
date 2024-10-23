package main

import (
	"github.com/tiago123456789/go-cli/internal/command"
	"github.com/tiago123456789/go-cli/pkg/cli"
	_ "github.com/tiago123456789/go-cli/plugins"
)

var rootCommand = *cli.Get()

func main() {

	rootCommand.AddCommand(command.NewListTodoCommand().Get())
	rootCommand.AddCommand(command.NewGetTodoByIdCommand().Get())
	rootCommand.AddCommand(command.NewDeleteTodoByIdCommand().Get())
	rootCommand.AddCommand(command.NewCreateTodoCommand().Get())
	rootCommand.AddCommand(command.NewUpdateTodoCommand().Get())

	rootCommand.Execute()

}
