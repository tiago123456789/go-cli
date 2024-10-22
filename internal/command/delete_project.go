package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

type DeleteProjectCommand struct {
}

func NewDeleteProjectCommand() *DeleteProjectCommand {
	return &DeleteProjectCommand{}
}

func (cP *DeleteProjectCommand) Get() *cobra.Command {
	var projectName, projectPath string

	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete boilerplate for a new project",
		Run: func(cmd *cobra.Command, args []string) {
			if projectName == "" {
				fmt.Println("You must supply a project name.")
				return
			}
			if projectPath == "" {
				fmt.Println("You must supply a project path.")
				return
			}
			fmt.Println("Creating project...")
		},
	}

	cmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project")
	cmd.Flags().StringVarP(&projectPath, "path", "p", "", "Path where the project will be created")

	return cmd
}
