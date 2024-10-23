package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tiago123456789/go-cli/internal/types"
	httpclient "github.com/tiago123456789/go-cli/pkg/http_client"
)

type GetTodoByIdCommand struct {
}

func NewGetTodoByIdCommand() *GetTodoByIdCommand {
	return &GetTodoByIdCommand{}
}

func (cP *GetTodoByIdCommand) Get() *cobra.Command {
	var id string

	var cmd = &cobra.Command{
		Use:   "get-todo-by-id",
		Short: "Get todo by id",
		Run: func(cmd *cobra.Command, args []string) {
			if id == "" {
				fmt.Println("You need provide the id")
				return
			}

			var todo types.Todo
			url := fmt.Sprintf("https://dummyjson.com/todos/%s", id)
			err := httpclient.GetOne(url, &todo)

			if err != nil {
				fmt.Println(err)
				fmt.Println("Couldn't found the todo")
				return
			}

			fmt.Println("Id => ", todo.Id, " Todo => ", todo.Todo, " Done => ", todo.Completed)
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Id of todo")

	return cmd
}
