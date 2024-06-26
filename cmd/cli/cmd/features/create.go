package datumfeature

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum feature",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "name of the feature")
	createCmd.Flags().String("display-name", "", "human friendly name of the feature")
	createCmd.Flags().StringP("description", "d", "", "description of the feature")
	createCmd.Flags().BoolP("enabled", "e", true, "enabled status of the feature")
	createCmd.Flags().StringSlice("tags", []string{}, "tags associated with the plan")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateFeatureInput, err error) {
	name := datum.Config.String("name")
	if name == "" {
		return input, datum.NewRequiredFieldMissingError("feature name")
	}

	enabled := datum.Config.Bool("enabled")

	input = datumclient.CreateFeatureInput{
		Name:    name,
		Enabled: &enabled,
	}

	displayName := datum.Config.String("display-name")
	if displayName != "" {
		input.DisplayName = &displayName
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	tags := datum.Config.Strings("tags")
	if len(tags) > 0 {
		input.Tags = tags
	}

	return input, nil
}

// create a new feature in the datum platform
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	f, err := client.CreateFeature(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(f)
}
