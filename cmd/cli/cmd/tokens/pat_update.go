package datumtokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
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
	datum.ViperBindFlag("pat.update.id", patUpdateCmd.Flags().Lookup("id"))

	patUpdateCmd.Flags().StringP("name", "n", "", "name of the personal access token")
	datum.ViperBindFlag("pat.update.name", patUpdateCmd.Flags().Lookup("name"))

	patUpdateCmd.Flags().StringP("description", "d", "", "description of the pat")
	datum.ViperBindFlag("pat.update.description", patUpdateCmd.Flags().Lookup("description"))

	patUpdateCmd.Flags().StringSliceP("add-organizations", "o", []string{}, "add organization(s) id to associate the pat with")
	datum.ViperBindFlag("pat.update.add-organizations", patUpdateCmd.Flags().Lookup("add-organizations"))

	patUpdateCmd.Flags().StringSliceP("remove-organizations", "r", []string{}, "remove organization(s) id to associate the pat with")
	datum.ViperBindFlag("pat.update.remove-organizations", patUpdateCmd.Flags().Lookup("remove-organizations"))
}

func updatePat(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	pID := viper.GetString("pat.update.id")
	if pID == "" {
		return datum.NewRequiredFieldMissingError("token id")
	}

	// Craft update input
	input := datumclient.UpdatePersonalAccessTokenInput{}

	name := viper.GetString("pat.update.name")
	if name != "" {
		input.Name = &name
	}

	description := viper.GetString("pat.update.description")
	if description != "" {
		input.Description = &description
	}

	addOrgs := viper.GetStringSlice("pat.update.add-organizations")
	if addOrgs != nil {
		input.AddOrganizationIDs = addOrgs
	}

	removeOrgs := viper.GetStringSlice("pat.update.remove-organizations")
	if removeOrgs != nil {
		input.RemoveOrganizationIDs = removeOrgs
	}

	o, err := cli.Client.UpdatePersonalAccessToken(ctx, pID, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
