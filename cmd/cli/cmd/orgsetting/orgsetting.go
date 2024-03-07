package datumorgsetting

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// orgSettingCmd represents the base org setting command when called without any subcommands
var orgSettingCmd = &cobra.Command{
	Use:   "orgsetting",
	Short: "The subcommands for working with the datum org settings",
}

func init() {
	datum.RootCmd.AddCommand(orgSettingCmd)
}
