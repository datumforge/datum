package datumwebhooks

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

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
	webhookCreateCmd.Flags().StringP("description", "d", "", "description of the webhook")
	webhookCreateCmd.Flags().StringP("url", "u", "", "the destination url the webhook is sent to")
	webhookCreateCmd.Flags().StringP("owner-id", "o", "", "owner of the webhook")
	webhookCreateCmd.Flags().BoolP("enabled", "e", true, "if the webhook is enabled")
}

func createwebhook(ctx context.Context) error {
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

	url := datum.Config.String("url")
	if url == "" {
		return datum.NewRequiredFieldMissingError("url")
	}

	enabled := datum.Config.Bool("enabled")
	ownerID := datum.Config.String("owner-id")
	description := datum.Config.String("description")

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
