package datumworkspace

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// workspaceCmd represents the base workspace command when called without any subcommands
var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "The subcommands for working with the datum workspace",
}

func init() {
	datum.RootCmd.AddCommand(workspaceCmd)
}
