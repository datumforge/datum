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
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	// filter options
	tokenID := datum.Config.String("id")

	var s []byte

	// if an api token ID is provided, filter on that api token, otherwise get all
	if tokenID == "" {
		tokens, err := client.GetAllAPITokens(ctx)
		if err != nil {
			return err
		}

		s, err = json.Marshal(tokens)
		if err != nil {
			return err
		}
	} else {
		token, err := client.GetAPITokenByID(ctx, tokenID)
		if err != nil {
			return err
		}

		s, err = json.Marshal(token)
		if err != nil {
			return err
		}
	}

	return datum.JSONPrint(s)
}
