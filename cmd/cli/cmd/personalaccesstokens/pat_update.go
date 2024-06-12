package datumtokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var patUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a new datum personal access token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updatePat(cmd.Context())
	},
}

func init() {
	patCmd.AddCommand(patUpdateCmd)

	patUpdateCmd.Flags().StringP("id", "i", "", "pat id to update")
	patUpdateCmd.Flags().StringP("name", "n", "", "name of the personal access token")
	patUpdateCmd.Flags().StringP("description", "d", "", "description of the pat")
	patUpdateCmd.Flags().StringSliceP("add-organizations", "o", []string{}, "add organization(s) id to associate the pat with")
	patUpdateCmd.Flags().StringSliceP("remove-organizations", "r", []string{}, "remove organization(s) id to associate the pat with")
}

func updatePat(ctx context.Context) error {
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
	input := datumclient.UpdatePersonalAccessTokenInput{}

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

	o, err := client.UpdatePersonalAccessToken(ctx, pID, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(o)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
