package datumintegrations

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var integrationCreateWithWebhookCmd = &cobra.Command{
	Use:   "createw",
	Short: "Create a new datum integration with a webhook",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createIntegrationWithWebhook(cmd.Context())
	},
}

func init() {
	integrationCmd.AddCommand(integrationCreateWithWebhookCmd)

	integrationCreateWithWebhookCmd.Flags().StringP("name", "n", "", "name of the integration")
	datum.ViperBindFlag("integration.createw.name", integrationCreateWithWebhookCmd.Flags().Lookup("name"))

	integrationCreateWithWebhookCmd.Flags().StringP("description", "d", "", "description of the integration")
	datum.ViperBindFlag("integration.createw.description", integrationCreateWithWebhookCmd.Flags().Lookup("description"))

	integrationCreateWithWebhookCmd.Flags().StringP("kind", "k", "", "the kind of integration")
	datum.ViperBindFlag("integration.createw.kind", integrationCreateWithWebhookCmd.Flags().Lookup("kind"))

	integrationCreateWithWebhookCmd.Flags().StringP("owner-id", "o", "", "owner of the integration")
	datum.ViperBindFlag("integration.createw.owner-id", integrationCreateWithWebhookCmd.Flags().Lookup("owner-id"))

	integrationCreateWithWebhookCmd.Flags().StringP("url", "u", "", "the destination url the webhook is sent to")
	datum.ViperBindFlag("integration.createw.url", integrationCreateWithWebhookCmd.Flags().Lookup("url"))

	integrationCreateWithWebhookCmd.Flags().BoolP("enabled", "e", true, "if the webhook is enabled")
	datum.ViperBindFlag("integration.createw.enabled", integrationCreateWithWebhookCmd.Flags().Lookup("enabled"))
}

func createIntegrationWithWebhook(ctx context.Context) error {
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := viper.GetString("integration.createw.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("name")
	}

	kind := viper.GetString("integration.createw.kind")
	if kind == "" {
		return datum.NewRequiredFieldMissingError("kind")
	}

	ownerID := viper.GetString("integration.createw.owner-id")
	description := viper.GetString("integration.createw.description")

	webhookURL := viper.GetString("webhook.create.url")
	webhookEnabled := viper.GetBool("webhook.create.enabled")

	webhook := datumclient.CreateWebhookInput{
		Name:           name,
		DestinationURL: webhookURL,
		Enabled:        &webhookEnabled,
	}

	if description != "" {
		webhook.Description = &description
	}

	if ownerID != "" {
		webhook.OwnerID = &ownerID
	}

	w, err := cli.Client.CreateWebhook(ctx, webhook, cli.Interceptor)
	if err != nil {
		return err
	}

	input := datumclient.CreateIntegrationInput{
		Name:       name,
		Kind:       &kind,
		WebhookIDs: []string{w.CreateWebhook.Webhook.ID},
	}

	if ownerID != "" {
		input.OwnerID = ownerID
	}

	i, err := cli.Client.CreateIntegration(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(i)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
