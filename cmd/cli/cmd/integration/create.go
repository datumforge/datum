package datumintegrations

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var integrationCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum integration",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(integrationCreateCmd)

	integrationCreateCmd.Flags().StringP("name", "n", "", "name of the integration")
	integrationCreateCmd.Flags().StringP("description", "d", "", "description of the integration")
	integrationCreateCmd.Flags().StringP("kind", "k", "", "the kind of integration")
	integrationCreateCmd.Flags().StringP("webhook-id", "w", "", "the webhook id to associate with the integration")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateIntegrationInput, err error) {
	input.Name = datum.Config.String("name")
	if input.Name == "" {
		return input, datum.NewRequiredFieldMissingError("name")
	}

	kind := datum.Config.String("kind")
	if kind == "" {
		return input, datum.NewRequiredFieldMissingError("kind")
	}

	input.Kind = &kind

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	webhookID := datum.Config.String("webhook-id")
	if webhookID != "" {
		input.WebhookIDs = append(input.WebhookIDs, webhookID)
	}

	return input, nil
}

// create an integration in the datum platform
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateIntegration(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
