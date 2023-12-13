// Package datumgroup is our cobra/viper cli for group endpoints
package datumgroup

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// groupCmd represents the base group command when called without any subcommands
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "The subcommands for working with datum groups",
}

func init() {
	datum.RootCmd.AddCommand(groupCmd)
}
