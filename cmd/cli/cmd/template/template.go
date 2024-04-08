package datumtemplates

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// templateCmd represents the base template command when called without any subcommands
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "The subcommands for working with the datum templates",
}

func init() {
	datum.RootCmd.AddCommand(templateCmd)
}
