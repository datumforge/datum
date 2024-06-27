package datumapitokens

import (
	"context"
	"time"

	"github.com/samber/lo"
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum api token",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "name of the api token token")
	createCmd.Flags().StringP("description", "d", "", "description of the api token")
	createCmd.Flags().DurationP("expiration", "e", 0, "duration of the api token to be valid, leave empty to never expire")
	createCmd.Flags().StringSlice("scopes", []string{"read", "write"}, "scopes to associate with the api token")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateAPITokenInput, err error) {
	name := datum.Config.String("name")
	if name == "" {
		return input, datum.NewRequiredFieldMissingError("token name")
	}

	input = datumclient.CreateAPITokenInput{
		Name:   name,
		Scopes: datum.Config.Strings("scopes"),
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	expiration := datum.Config.Duration("expiration")
	if expiration != 0 {
		input.ExpiresAt = lo.ToPtr(time.Now().Add(expiration))
	}

	return input, nil
}

// create a new datum api token
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateAPIToken(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
