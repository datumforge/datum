package datumusersetting

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// userSettingCmd represents the base user setting command when called without any subcommands
var userSettingCmd = &cobra.Command{
	Use:   "usersetting",
	Short: "The subcommands for working with the datum user settings",
}

func init() {
	datum.RootCmd.AddCommand(userSettingCmd)
}
