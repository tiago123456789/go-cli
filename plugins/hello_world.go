package plugins

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tiago123456789/go-cli/pkg/cli"
)

func init() {
	cli.Get().AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a hello message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}
