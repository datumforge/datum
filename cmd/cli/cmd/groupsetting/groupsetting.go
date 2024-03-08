package datumgroupsetting

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// groupSettingCmd represents the base group setting command when called without any subcommands
var groupSettingCmd = &cobra.Command{
	Use:   "groupsetting",
	Short: "The subcommands for working with the datum group settings",
}

func init() {
	datum.RootCmd.AddCommand(groupSettingCmd)
}
