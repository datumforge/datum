package datumapitokens

import (
	"context"
	"encoding/json"
	"time"

	"github.com/samber/lo"
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var apiTokenCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum api token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createAPIToken(cmd.Context())
	},
}

func init() {
	apiTokenCmd.AddCommand(apiTokenCreateCmd)

	apiTokenCreateCmd.Flags().StringP("name", "n", "", "name of the api token token")
	apiTokenCreateCmd.Flags().StringP("description", "d", "", "description of the api token")
	apiTokenCreateCmd.Flags().DurationP("expiration", "e", 0, "duration of the api token to be valid, leave empty to never expire")
	apiTokenCreateCmd.Flags().StringSlice("scopes", []string{"read", "write"}, "scopes to associate with the api token")
}

func createAPIToken(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := datum.Config.String("name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("token name")
	}

	input := datumclient.CreateAPITokenInput{
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

	o, err := client.CreateAPIToken(ctx, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(o)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
