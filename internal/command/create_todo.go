package command

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tiago123456789/go-cli/internal/types"
	httpclient "github.com/tiago123456789/go-cli/pkg/http_client"
)

type CreateTodoCommand struct {
}

func NewCreateTodoCommand() *CreateTodoCommand {
	return &CreateTodoCommand{}
}

func (cP *CreateTodoCommand) Get() *cobra.Command {
	var userId, description string
	var completed bool = false

	var cmd = &cobra.Command{
		Use:   "create-todo",
		Short: "Create todo",
		Run: func(cmd *cobra.Command, args []string) {
			if userId == "" {
				fmt.Println("You need provide the userId")
				return
			}

			if description == "" {
				fmt.Println("You need provide the description")
				return
			}

			userIdValue, _ := strconv.Atoi(userId)

			newTodo := types.Todo{
				Todo:      description,
				Completed: completed,
				UserId:    userIdValue,
			}

			var todoReturned types.Todo
			err := httpclient.Post[types.Todo, types.Todo](
				"https://dummyjson.com/todos/add",
				&newTodo,
				&todoReturned,
			)

			if err != nil {
				fmt.Println(err)
				fmt.Printf("Error: %v", err)
				return
			}

			fmt.Println("Id => ", todoReturned.Id, " | Description => ", todoReturned.Todo)
			fmt.Println("Created todo with success")
		},
	}

	cmd.Flags().StringVarP(&userId, "user-id", "u", "", "User id")
	cmd.Flags().StringVarP(&description, "description", "d", "", "The todo description")
	cmd.Flags().BoolVarP(&completed, "completed", "c", false, "Make todo is complete")

	return cmd
}
