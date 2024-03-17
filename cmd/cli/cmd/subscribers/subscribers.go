package datumsubscribers

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// subscribersCmd represents the base subscribers command when called without any subcommands
var subscribersCmd = &cobra.Command{
	Use:   "subscribers",
	Short: "The subcommands for working with the subscribers of a datum organization",
}

func init() {
	datum.RootCmd.AddCommand(subscribersCmd)
}
