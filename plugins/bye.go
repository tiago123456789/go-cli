package plugins

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tiago123456789/go-cli/pkg/cli"
)

func init() {
	cli.Get().AddCommand(byeCmd)
}

var byeCmd = &cobra.Command{
	Use:   "bye",
	Short: "Say bye message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bye, World!")
	},
}
