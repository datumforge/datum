package datumgroupmembers

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// groupMembersCmd represents the base groupMembers command when called without any subcommands
var groupMembersCmd = &cobra.Command{
	Use:   "groupmembers",
	Short: "The subcommands for working with the members of a datum group",
}

func init() {
	datum.RootCmd.AddCommand(groupMembersCmd)
}
