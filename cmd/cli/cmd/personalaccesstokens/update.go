package datumtokens

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum personal access token",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "pat id to update")
	updateCmd.Flags().StringP("name", "n", "", "name of the personal access token")
	updateCmd.Flags().StringP("description", "d", "", "description of the pat")
	updateCmd.Flags().StringSliceP("add-organizations", "o", []string{}, "add organization(s) id to associate the pat with")
	updateCmd.Flags().StringSliceP("remove-organizations", "r", []string{}, "remove organization(s) id to associate the pat with")
}

// updateValidation validates the required fields for the command
func updateValidation() (id string, input datumclient.UpdatePersonalAccessTokenInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("token id")
	}

	name := datum.Config.String("name")
	if name != "" {
		input.Name = &name
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	addOrgs := datum.Config.Strings("add-organizations")
	if addOrgs != nil {
		input.AddOrganizationIDs = addOrgs
	}

	removeOrgs := datum.Config.Strings("remove-organizations")
	if removeOrgs != nil {
		input.RemoveOrganizationIDs = removeOrgs
	}

	return id, input, nil
}

// update an existing datum personal access token
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdatePersonalAccessToken(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
