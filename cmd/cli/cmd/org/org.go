package datumorg

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// orgCmd represents the base org command when called without any subcommands
var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "The subcommands for working with the datum organization",
}

func init() {
	datum.RootCmd.AddCommand(orgCmd)
}
