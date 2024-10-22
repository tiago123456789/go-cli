package cli

import "github.com/spf13/cobra"

var rootCommand = &cobra.Command{
	Use:   "app",
	Short: "Study how to build CLI using Golang + how to use plugins too",
	Long:  "Study how to build CLI using Golang + how to use plugins too",
}

func Get() *cobra.Command {
	return rootCommand
}
