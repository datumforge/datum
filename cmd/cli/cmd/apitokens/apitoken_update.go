package datumapitokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

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
	apiTokenUpdateCmd.Flags().StringP("name", "n", "", "name of the api token token")
	apiTokenUpdateCmd.Flags().StringP("description", "d", "", "description of the api token")
	apiTokenUpdateCmd.Flags().StringSlice("scopes", []string{}, "scopes to add to the api token")
}

func updateAPIToken(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	var s []byte

	pID := datum.Config.String("id")
	if pID == "" {
		return datum.NewRequiredFieldMissingError("token id")
	}

	// Craft update input
	input := datumclient.UpdateAPITokenInput{}

	name := datum.Config.String("name")
	if name != "" {
		input.Name = &name
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	scopes := datum.Config.Strings("scopes")
	if len(scopes) > 0 {
		input.Scopes = scopes
	}

	o, err := client.UpdateAPIToken(ctx, pID, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(o)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
