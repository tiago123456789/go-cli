package command

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tiago123456789/go-cli/internal/types"
	httpclient "github.com/tiago123456789/go-cli/pkg/http_client"
)

type UpdateTodoCommand struct {
}

func NewUpdateTodoCommand() *UpdateTodoCommand {
	return &UpdateTodoCommand{}
}

func (cP *UpdateTodoCommand) Get() *cobra.Command {
	var id, userId, description string
	var completed bool = false

	var cmd = &cobra.Command{
		Use:   "update-todo",
		Short: "Update todo",
		Run: func(cmd *cobra.Command, args []string) {
			if id == "" {
				fmt.Println("You need provide the id")
				return
			}

			if userId == "" {
				fmt.Println("You need provide the userId")
				return
			}

			if description == "" {
				fmt.Println("You need provide the description")
				return
			}

			userIdValue, _ := strconv.Atoi(userId)

			todoUpdate := types.Todo{
				Todo:      description,
				Completed: completed,
				UserId:    userIdValue,
			}

			url := fmt.Sprintf("https://dummyjson.com/todos/%s", id)
			var todoReturned types.Todo
			err := httpclient.Put[types.Todo, types.Todo](
				url,
				&todoUpdate,
				&todoReturned,
			)

			if err != nil {
				fmt.Println(err)
				fmt.Printf("Error: %v", err)
				return
			}

			fmt.Println("Id => ", todoReturned.Id, " | Description => ", todoReturned.Todo)
			fmt.Println("Updated todo with success")
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Id")
	cmd.Flags().StringVarP(&userId, "user-id", "u", "", "User id")
	cmd.Flags().StringVarP(&description, "description", "d", "", "The todo description")
	cmd.Flags().BoolVarP(&completed, "completed", "c", false, "Make todo is complete")

	return cmd
}
