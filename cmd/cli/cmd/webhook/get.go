package datumwebhooks

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get details of existing datum webhooks",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "get a specific webhook by ID")
}

// get retrieves webhooks from the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id := datum.Config.String("id")

	if id != "" {
		o, err := client.GetWebhookByID(ctx, id)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	o, err := client.GetAllWebhooks(ctx)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
