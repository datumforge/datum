package datumapitokens

import (
	"context"
	"encoding/json"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	datum.ViperBindFlag("apitoken.create.name", apiTokenCreateCmd.Flags().Lookup("name"))

	apiTokenCreateCmd.Flags().StringP("description", "d", "", "description of the api token")
	datum.ViperBindFlag("apitoken.create.description", apiTokenCreateCmd.Flags().Lookup("description"))

	apiTokenCreateCmd.Flags().DurationP("expiration", "e", 0, "duration of the api token to be valid, leave empty to never expire")
	datum.ViperBindFlag("apitoken.create.expiration", apiTokenCreateCmd.Flags().Lookup("expiration"))
}

func createAPIToken(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := viper.GetString("apitoken.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("token name")
	}

	input := datumclient.CreatePersonalAccessTokenInput{
		Name: name,
	}

	description := viper.GetString("apitoken.create.description")
	if description != "" {
		input.Description = &description
	}

	expiration := viper.GetDuration("apitoken.create.expiration")
	if expiration != 0 {
		input.ExpiresAt = time.Now().Add(expiration)
	}

	o, err := cli.Client.CreatePersonalAccessToken(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
