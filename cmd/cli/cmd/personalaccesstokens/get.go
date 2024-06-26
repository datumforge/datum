package datumtokens

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get personal access tokens",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "personal access token id to query")
}

func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	id := datum.Config.String("id")

	// if an id is provided, filter on the id, otherwise get all
	if id == "" {
		o, err := client.GetAllPersonalAccessTokens(ctx)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	o, err := client.GetPersonalAccessTokenByID(ctx, id)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
