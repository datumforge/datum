package datumtokens

import (
	"context"
	"time"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum personal access token",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "name of the personal access token")
	createCmd.Flags().StringP("description", "d", "", "description of the pat")
	createCmd.Flags().StringSliceP("organizations", "o", []string{}, "organization(s) id to associate the pat with")
	createCmd.Flags().DurationP("expiration", "e", 0, "duration of the pat to be valid, defaults to 7 days")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreatePersonalAccessTokenInput, err error) {
	input.Name = datum.Config.String("name")
	if input.Name == "" {
		return input, datum.NewRequiredFieldMissingError("token name")
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

	return input, nil
}

// create a new datum personal access token
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreatePersonalAccessToken(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
