package datumapitokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var apiTokenDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum api token token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteAPIToken(cmd.Context())
	},
}

func init() {
	apiTokenCmd.AddCommand(apiTokenDeleteCmd)

	apiTokenDeleteCmd.Flags().StringP("id", "i", "", "api token id to delete")
	datum.ViperBindFlag("apitoken.delete.id", apiTokenDeleteCmd.Flags().Lookup("id"))
}

func deleteAPIToken(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	tokenID := viper.GetString("apitoken.delete.id")
	if tokenID == "" {
		return datum.NewRequiredFieldMissingError("token id")
	}

	o, err := client.DeleteAPIToken(ctx, tokenID, client.Config().Interceptors...)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
