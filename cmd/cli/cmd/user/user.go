// Package datumuser is our cobra/viper cli for user endpoints
package datumuser

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// userCmd represents the base user command when called without any subcommands
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "The subcommands for working with the datum user",
}

func init() {
	datum.RootCmd.AddCommand(userCmd)
}
