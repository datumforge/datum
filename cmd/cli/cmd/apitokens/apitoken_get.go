package datumapitokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var apiTokenGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get or list api tokens",
	RunE: func(cmd *cobra.Command, args []string) error {
		return apiTokens(cmd.Context())
	},
}

func init() {
	apiTokenCmd.AddCommand(apiTokenGetCmd)

	apiTokenGetCmd.Flags().StringP("id", "i", "", "api token id to query")
}

func apiTokens(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	tokenID := datum.Config.String("id")

	var s []byte

	// if an api token ID is provided, filter on that api token, otherwise get all
	if tokenID == "" {
		tokens, err := client.GetAllAPITokens(ctx)
		cobra.CheckErr(err)

		s, err = json.Marshal(tokens)
		cobra.CheckErr(err)
	} else {
		token, err := client.GetAPITokenByID(ctx, tokenID)
		cobra.CheckErr(err)

		s, err = json.Marshal(token)
		cobra.CheckErr(err)
	}

	return datum.JSONPrint(s)
}
