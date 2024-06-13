package datumtokens

import (
	"context"
	"encoding/json"
	"time"

	"github.com/spf13/cobra"

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
	patCreateCmd.Flags().StringP("description", "d", "", "description of the pat")
	patCreateCmd.Flags().StringSliceP("organizations", "o", []string{}, "organization(s) id to associate the pat with")
	patCreateCmd.Flags().DurationP("expiration", "e", 0, "duration of the pat to be valid, defaults to 7 days")
}

func createPat(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := datum.Config.String("name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("token name")
	}

	input := datumclient.CreatePersonalAccessTokenInput{
		Name: name,
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	organizations := datum.Config.Strings("organizations")
	if organizations != nil {
		input.OrganizationIDs = organizations
	}

	expiration := datum.Config.Duration("expiration")
	if expiration != 0 {
		input.ExpiresAt = time.Now().Add(expiration)
	}

	o, err := client.CreatePersonalAccessToken(ctx, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(o)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
