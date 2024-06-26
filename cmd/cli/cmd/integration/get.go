package datumintegrations

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get details of existing datum integrations",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "get a specific integration by ID")
}

// get retrieves all integrations or a specific integration by ID
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id := datum.Config.String("id")

	if id != "" {
		o, err := client.GetIntegrationByID(ctx, id)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	o, err := client.GetAllIntegrations(ctx)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
