package command

import (
	"fmt"

	"github.com/spf13/cobra"
	httpclient "github.com/tiago123456789/go-cli/pkg/http_client"
)

type DeleteTodoByIdCommand struct {
}

func NewDeleteTodoByIdCommand() *DeleteTodoByIdCommand {
	return &DeleteTodoByIdCommand{}
}

func (cP *DeleteTodoByIdCommand) Get() *cobra.Command {
	var id string

	var cmd = &cobra.Command{
		Use:   "delete-todo-by-id",
		Short: "Delete by id provided",
		Run: func(cmd *cobra.Command, args []string) {
			if id == "" {
				fmt.Println("You need provide the id")
				return
			}

			url := fmt.Sprintf("https://dummyjson.com/todos/%s", id)
			err := httpclient.Delete(url)

			if err != nil {
				fmt.Println(err)
				fmt.Printf("Error: %v", err)
				return
			}

			fmt.Println("Delete with success")
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Id of todo")

	return cmd
}
