package datumorgmembers

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// orgMembersCmd represents the base orgMembers command when called without any subcommands
var orgMembersCmd = &cobra.Command{
	Use:   "orgmembers",
	Short: "The subcommands for working with the members of a datum organization",
}

func init() {
	datum.RootCmd.AddCommand(orgMembersCmd)
}
