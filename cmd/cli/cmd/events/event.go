package datumevents

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// eventCmd represents the base event command when called without any subcommands
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "The subcommands for working with datum events",
}

func init() {
	datum.RootCmd.AddCommand(eventCmd)
}
