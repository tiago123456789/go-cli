package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tiago123456789/go-cli/internal/types"
	httpclient "github.com/tiago123456789/go-cli/pkg/http_client"
)

type ListTodoCommand struct {
}

func NewListTodoCommand() *ListTodoCommand {
	return &ListTodoCommand{}
}

func (cP *ListTodoCommand) Get() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-todo",
		Short: "List todos",
		Run: func(cmd *cobra.Command, args []string) {
			var todos types.Todos
			err := httpclient.GetOne("https://dummyjson.com/todos", &todos)

			if err != nil {
				fmt.Println(err)
				fmt.Println("Couldn't found the todos")
				return
			}

			for _, todo := range todos.Items {
				fmt.Println("Id => ", todo.Id, " Todo => ", todo.Todo, " Done => ", todo.Completed)
			}
		},
	}

	return cmd
}
