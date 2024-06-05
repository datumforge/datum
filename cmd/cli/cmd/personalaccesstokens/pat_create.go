package datumtokens

import (
	"context"
	"encoding/json"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var patCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum personal access token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createPat(cmd.Context())
	},
}

func init() {
	patCmd.AddCommand(patCreateCmd)

	patCreateCmd.Flags().StringP("name", "n", "", "name of the personal access token")
	datum.ViperBindFlag("pat.create.name", patCreateCmd.Flags().Lookup("name"))

	patCreateCmd.Flags().StringP("description", "d", "", "description of the pat")
	datum.ViperBindFlag("pat.create.description", patCreateCmd.Flags().Lookup("description"))

	patCreateCmd.Flags().StringSliceP("organizations", "o", []string{}, "organization(s) id to associate the pat with")
	datum.ViperBindFlag("pat.create.organizations", patCreateCmd.Flags().Lookup("organizations"))

	patCreateCmd.Flags().DurationP("expiration", "e", 0, "duration of the pat to be valid, defaults to 7 days")
	datum.ViperBindFlag("pat.create.expiration", patCreateCmd.Flags().Lookup("expiration"))
}

func createPat(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := viper.GetString("pat.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("token name")
	}

	input := datumclient.CreatePersonalAccessTokenInput{
		Name: name,
	}

	description := viper.GetString("pat.create.description")
	if description != "" {
		input.Description = &description
	}

	organizations := viper.GetStringSlice("pat.create.organizations")
	if organizations != nil {
		input.OrganizationIDs = organizations
	}

	expiration := viper.GetDuration("pat.create.expiration")
	if expiration != 0 {
		input.ExpiresAt = time.Now().Add(expiration)
	}

	o, err := client.CreatePersonalAccessToken(ctx, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
