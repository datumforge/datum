package datumtokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var patGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get or list personal access tokens",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pats(cmd.Context())
	},
}

func init() {
	patCmd.AddCommand(patGetCmd)

	patGetCmd.Flags().StringP("id", "i", "", "pat id to query")
}

func pats(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	// filter options
	pID := datum.Config.String("id")

	var s []byte

	// if an pat ID is provided, filter on that pat, otherwise get all
	if pID == "" {
		tokens, err := client.GetAllPersonalAccessTokens(ctx)
		if err != nil {
			return err
		}

		s, err = json.Marshal(tokens)
		if err != nil {
			return err
		}
	} else {
		token, err := client.GetPersonalAccessTokenByID(ctx, pID)
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
