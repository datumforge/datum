package datumwebhooks

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
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
	datum.ViperBindFlag("webhook.get.id", webhookGetCmd.Flags().Lookup("id"))
}

func webhooks(ctx context.Context) error {
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	oID := viper.GetString("webhook.get.id")

	var s []byte

	writer := tables.NewTableWriter(webhookCmd.OutOrStdout(), "ID", "Name", "Description")

	if oID != "" {
		webhook, err := cli.Client.GetWebhookByID(ctx, oID, cli.Interceptor)
		if err != nil {
			return err
		}

		if viper.GetString("output.format") == "json" {
			s, err := json.Marshal(webhook.Webhook)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer.AddRow(webhook.Webhook.ID, webhook.Webhook.Name, *webhook.Webhook.Description)
		writer.Render()

		return nil
	}

	webhooks, err := cli.Client.GetAllWebhooks(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(webhooks.Webhooks)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		return datum.JSONPrint(s)
	}

	for _, webhook := range webhooks.Webhooks.Edges {
		writer.AddRow(webhook.Node.ID, webhook.Node.Name, *webhook.Node.Description)
	}

	writer.Render()

	return nil
}
