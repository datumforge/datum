package datumtokens

import (
	"context"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
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

	patCreateCmd.Flags().StringP("owner-id", "o", "", "the owner of the personal access token")
	datum.ViperBindFlag("pat.create.owner-id", patCreateCmd.Flags().Lookup("owner-id"))
}

func createPat(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	var s []byte

	name := viper.GetString("pat.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("token name")
	}

	owner := viper.GetString("pat.create.owner-id")
	if owner == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	description := viper.GetString("pat.create.description")

	input := datumclient.CreatePersonalAccessTokenInput{
		Name:    name,
		OwnerID: owner,
	}

	if description != "" {
		input.Description = &description
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
