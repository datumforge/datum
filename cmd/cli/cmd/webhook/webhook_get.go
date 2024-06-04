package datumwebhooks

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var webhookGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details of existing datum webhooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return webhooks(cmd.Context())
	},
}

func init() {
	webhookCmd.AddCommand(webhookGetCmd)

	webhookGetCmd.Flags().StringP("id", "i", "", "get a specific webhook by ID")
}

func webhooks(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	oID := datum.Config.String("id")

	var s []byte

	writer := tables.NewTableWriter(webhookCmd.OutOrStdout(), "ID", "Name", "Description", "Destination URL", "Enabled")

	if oID != "" {
		webhook, err := client.GetWebhookByID(ctx, oID)
		if err != nil {
			return err
		}

		if datum.OutputFormat == datum.JSONOutput {
			s, err := json.Marshal(webhook.Webhook)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer.AddRow(webhook.Webhook.ID, webhook.Webhook.Name, *webhook.Webhook.Description, webhook.Webhook.DestinationURL, webhook.Webhook.Enabled)
		writer.Render()

		return nil
	}

	webhooks, err := client.GetAllWebhooks(ctx)
	if err != nil {
		return err
	}

	s, err = json.Marshal(webhooks.Webhooks)
	if err != nil {
		return err
	}

	if datum.OutputFormat == datum.JSONOutput {
		return datum.JSONPrint(s)
	}

	for _, webhook := range webhooks.Webhooks.Edges {
		writer.AddRow(webhook.Node.ID, webhook.Node.Name, *webhook.Node.Description, webhook.Node.DestinationURL, webhook.Node.Enabled)
	}

	writer.Render()

	return nil
}
