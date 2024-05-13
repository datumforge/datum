package datumintegrations

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// integrationCmd represents the base integration command when called without any subcommands
var integrationCmd = &cobra.Command{
	Use:   "integration",
	Short: "The subcommands for working with datum integrations",
}

func init() {
	datum.RootCmd.AddCommand(integrationCmd)
}
