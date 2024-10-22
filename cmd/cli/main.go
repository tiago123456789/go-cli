package main

import (
	"github.com/tiago123456789/go-cli/internal/command"
	"github.com/tiago123456789/go-cli/pkg/cli"
	_ "github.com/tiago123456789/go-cli/plugins"
)

var rootCommand = *cli.Get()

func main() {

	rootCommand.AddCommand(command.NewCreateProjectCommand().Get())
	rootCommand.AddCommand(command.NewDeleteProjectCommand().Get())
	rootCommand.AddCommand(command.NewSearchGithubRepositoriesCommand().Get())

	rootCommand.Execute()

}
