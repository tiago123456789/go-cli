package types

import (
	"github.com/spf13/cobra"
)

type CustomCommand interface {
	Get() *cobra.Command
}
