package datumapitokens

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// apiTokenCmd represents the base apiTokenCmd command when called without any subcommands
var apiTokenCmd = &cobra.Command{
	Use:   "pat",
	Short: "The subcommands for working with api tokens",
}

func init() {
	datum.RootCmd.AddCommand(apiTokenCmd)
}
