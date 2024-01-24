package datuminvite

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// inviteCmd represents the base invite command when called without any subcommands
var inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "The subcommands for working with the invitations of a datum organization",
}

func init() {
	datum.RootCmd.AddCommand(inviteCmd)
}
