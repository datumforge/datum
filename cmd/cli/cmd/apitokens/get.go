package datumapitokens

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get or list api tokens",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "api token id to query")
}

// get retrieves api tokens from the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	tokenID := datum.Config.String("id")

	// if an api token ID is provided, filter on that api token, otherwise get all
	if tokenID == "" {
		tokens, err := client.GetAllAPITokens(ctx)
		cobra.CheckErr(err)

		return consoleOutput(tokens)
	}

	token, err := client.GetAPITokenByID(ctx, tokenID)
	cobra.CheckErr(err)

	return consoleOutput(token)
}
