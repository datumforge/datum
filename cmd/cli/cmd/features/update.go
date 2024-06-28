package datumfeature

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum feature",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "feature id to update")
	updateCmd.Flags().String("display-name", "", "human friendly name of the feature")
	updateCmd.Flags().StringP("description", "d", "", "description of the feature")
	updateCmd.Flags().BoolP("enabled", "e", true, "enable or disable the feature")
	updateCmd.Flags().StringSlice("tags", []string{}, "tags associated with the feature")
}

// updateValidation validates the required fields for the command
func updateValidation() (id string, input datumclient.UpdateFeatureInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("feature id")
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	displayName := datum.Config.String("display-name")
	if displayName != "" {
		input.DisplayName = &displayName
	}

	enabled := datum.Config.Bool("enabled")
	if !enabled {
		input.Enabled = &enabled
	}

	tags := datum.Config.Strings("tags")
	if len(tags) > 0 {
		input.Tags = tags
	}

	return id, input, nil
}

// update an existing feature in the datum platform
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	f, err := client.UpdateFeature(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(f)
}
