package datumapitokens

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a datum api token token",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "api token id to update")
	updateCmd.Flags().StringP("name", "n", "", "name of the api token token")
	updateCmd.Flags().StringP("description", "d", "", "description of the api token")
	updateCmd.Flags().StringSlice("scopes", []string{}, "scopes to add to the api token")
}

// updateValidation validates the required fields for the command
func updateValidation() (id string, input datumclient.UpdateAPITokenInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("token id")
	}

	// Craft update input
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

	return id, input, nil
}

// update an existing datum api token
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdateAPIToken(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
