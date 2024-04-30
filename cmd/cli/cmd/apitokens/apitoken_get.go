package datumapitokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
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
	datum.ViperBindFlag("apitoken.get.id", apiTokenGetCmd.Flags().Lookup("id"))
}

func apiTokens(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	// filter options
	tokenID := viper.GetString("apitoken.get.id")

	var s []byte

	// if an api token ID is provided, filter on that api token, otherwise get all
	if tokenID == "" {
		tokens, err := cli.Client.GetAllAPITokens(ctx, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err = json.Marshal(tokens)
		if err != nil {
			return err
		}
	} else {
		token, err := cli.Client.GetAPITokenByID(ctx, tokenID, cli.Interceptor)
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
