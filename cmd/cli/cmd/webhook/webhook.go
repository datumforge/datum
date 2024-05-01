package datumwebhooks

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// webhookCmd represents the base webhook command when called without any subcommands
var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "The subcommands for working with datum webhooks",
}

func init() {
	datum.RootCmd.AddCommand(webhookCmd)
}
