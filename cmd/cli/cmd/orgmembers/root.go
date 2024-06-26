package datumorgmembers

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// cmd represents the base orgMembers command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "org-members",
	Short: "the subcommands for working with the members of a datum organization",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}
