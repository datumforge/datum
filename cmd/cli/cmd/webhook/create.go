package datumwebhooks

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createcmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum webhook",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createcmd)

	createcmd.Flags().StringP("name", "n", "", "name of the webhook")
	createcmd.Flags().StringP("description", "d", "", "description of the webhook")
	createcmd.Flags().StringP("url", "u", "", "the destination url the webhook is sent to")
	createcmd.Flags().BoolP("enabled", "e", true, "if the webhook is enabled")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateWebhookInput, err error) {
	input.Name = datum.Config.String("name")
	if input.Name == "" {
		return input, datum.NewRequiredFieldMissingError("name")
	}

	input.DestinationURL = datum.Config.String("url")
	if input.DestinationURL == "" {
		return input, datum.NewRequiredFieldMissingError("url")
	}

	enabled := datum.Config.Bool("enabled")
	if !enabled {
		input.Enabled = &enabled
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	return input, nil
}

// create a new webhook
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateWebhook(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
