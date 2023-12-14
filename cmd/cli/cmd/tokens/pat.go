package datumtokens

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// patCmd represents the base patCmd command when called without any subcommands
var patCmd = &cobra.Command{
	Use:   "pat",
	Short: "The subcommands for working with personal access tokens",
}

func init() {
	datum.RootCmd.AddCommand(patCmd)
}
