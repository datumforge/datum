package datumwebhooks

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var webhookCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum webhook",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createwebhook(cmd.Context())
	},
}

func init() {
	webhookCmd.AddCommand(webhookCreateCmd)

	webhookCreateCmd.Flags().StringP("name", "n", "", "name of the webhook")
	datum.ViperBindFlag("webhook.create.name", webhookCreateCmd.Flags().Lookup("name"))

	webhookCreateCmd.Flags().StringP("description", "d", "", "description of the webhook")
	datum.ViperBindFlag("webhook.create.description", webhookCreateCmd.Flags().Lookup("description"))

	webhookCreateCmd.Flags().StringP("url", "u", "", "the destination url the webhook is sent to")
	datum.ViperBindFlag("webhook.create.url", webhookCreateCmd.Flags().Lookup("url"))

	webhookCreateCmd.Flags().StringP("owner-id", "o", "", "owner of the webhook")
	datum.ViperBindFlag("webhook.create.owner-id", webhookCreateCmd.Flags().Lookup("owner-id"))

	webhookCreateCmd.Flags().BoolP("enabled", "e", true, "if the webhook is enabled")
	datum.ViperBindFlag("webhook.create.enabled", webhookCreateCmd.Flags().Lookup("enabled"))
}

func createwebhook(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := viper.GetString("webhook.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("name")
	}

	url := viper.GetString("webhook.create.url")
	if url == "" {
		return datum.NewRequiredFieldMissingError("url")
	}

	enabled := viper.GetBool("webhook.create.enabled")
	ownerID := viper.GetString("webhook.create.owner-id")
	description := viper.GetString("webhook.create.description")

	input := datumclient.CreateWebhookInput{
		Name:           name,
		DestinationURL: url,
	}

	if !enabled {
		input.Enabled = &enabled
	}

	if description != "" {
		input.Description = &description
	}

	if ownerID != "" {
		input.OwnerID = &ownerID
	}

	w, err := client.CreateWebhook(ctx, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(w)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
