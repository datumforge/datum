package datumapitokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var apiTokenUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a datum api token token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateAPIToken(cmd.Context())
	},
}

func init() {
	apiTokenCmd.AddCommand(apiTokenUpdateCmd)

	apiTokenUpdateCmd.Flags().StringP("id", "i", "", "api token id to update")
	datum.ViperBindFlag("apitoken.update.id", apiTokenUpdateCmd.Flags().Lookup("id"))

	apiTokenUpdateCmd.Flags().StringP("name", "n", "", "name of the api token token")
	datum.ViperBindFlag("apitoken.update.name", apiTokenUpdateCmd.Flags().Lookup("name"))

	apiTokenUpdateCmd.Flags().StringP("description", "d", "", "description of the api token")
	datum.ViperBindFlag("apitoken.update.description", apiTokenUpdateCmd.Flags().Lookup("description"))

	apiTokenUpdateCmd.Flags().StringSlice("scopes", []string{}, "scopes to add to the api token")
	datum.ViperBindFlag("apitoken.update.scopes", apiTokenCreateCmd.Flags().Lookup("scopes"))
}

func updateAPIToken(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	pID := viper.GetString("apitoken.update.id")
	if pID == "" {
		return datum.NewRequiredFieldMissingError("token id")
	}

	// Craft update input
	input := datumclient.UpdateAPITokenInput{}

	name := viper.GetString("apitoken.update.name")
	if name != "" {
		input.Name = &name
	}

	description := viper.GetString("apitoken.update.description")
	if description != "" {
		input.Description = &description
	}

	scopes := viper.GetStringSlice("apitoken.update.scopes")
	if len(scopes) > 0 {
		input.Scopes = scopes
	}

	o, err := cli.Client.UpdateAPIToken(ctx, pID, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
