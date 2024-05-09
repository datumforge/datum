package datumintegrations

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var integrationCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum integration",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createintegration(cmd.Context())
	},
}

func init() {
	integrationCmd.AddCommand(integrationCreateCmd)

	integrationCreateCmd.Flags().StringP("name", "n", "", "name of the integration")
	datum.ViperBindFlag("integration.create.name", integrationCreateCmd.Flags().Lookup("name"))

	integrationCreateCmd.Flags().StringP("description", "d", "", "description of the integration")
	datum.ViperBindFlag("integration.create.description", integrationCreateCmd.Flags().Lookup("description"))

	integrationCreateCmd.Flags().StringP("kind", "k", "", "the kind of integration")
	datum.ViperBindFlag("integration.create.kind", integrationCreateCmd.Flags().Lookup("kind"))

	integrationCreateCmd.Flags().StringP("owner-id", "o", "", "owner of the integration")
	datum.ViperBindFlag("integration.create.owner-id", integrationCreateCmd.Flags().Lookup("owner-id"))
}

func createintegration(ctx context.Context) error {
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := viper.GetString("integration.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("name")
	}

	kind := viper.GetString("integration.create.kind")
	if kind == "" {
		return datum.NewRequiredFieldMissingError("kind")
	}

	ownerID := viper.GetString("integration.create.owner-id")
	description := viper.GetString("integration.create.description")

	input := datumclient.CreateIntegrationInput{
		Name: name,
		Kind: &kind,
	}

	if description != "" {
		input.Description = &description
	}

	if ownerID != "" {
		input.OwnerID = ownerID
	}

	w, err := cli.Client.CreateIntegration(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(w)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
