package datumintegrations

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var integrationCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum integration",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createIntegration(cmd.Context())
	},
}

func init() {
	integrationCmd.AddCommand(integrationCreateCmd)

	integrationCreateCmd.Flags().StringP("name", "n", "", "name of the integration")
	integrationCreateCmd.Flags().StringP("description", "d", "", "description of the integration")
	integrationCreateCmd.Flags().StringP("kind", "k", "", "the kind of integration")
	integrationCreateCmd.Flags().StringP("owner-id", "o", "", "owner of the integration")
	integrationCreateCmd.Flags().StringP("webhook-id", "w", "", "the webhook id to associate with the integration")
}

func createIntegration(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := datum.Config.String("name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("name")
	}

	kind := datum.Config.String("kind")
	if kind == "" {
		return datum.NewRequiredFieldMissingError("kind")
	}

	ownerID := datum.Config.String("owner-id")
	description := datum.Config.String("description")

	input := datumclient.CreateIntegrationInput{
		Name: name,
		Kind: &kind,
	}

	if webhookID := datum.Config.String("webhook-id"); webhookID != "" {
		input.WebhookIDs = append(input.WebhookIDs, webhookID)
	}

	if description != "" {
		input.Description = &description
	}

	if ownerID != "" {
		input.OwnerID = &ownerID
	}

	w, err := client.CreateIntegration(ctx, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(w)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
