package datumevents

import (
	"context"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var eventCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum event",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(eventCreateCmd)

	eventCreateCmd.Flags().StringP("type", "t", "", "type of the event")
	eventCreateCmd.Flags().StringP("metadata", "m", "", "metadata for the event")
	eventCreateCmd.Flags().StringSliceP("user-ids", "u", []string{}, "user id associated with the event")
	eventCreateCmd.Flags().StringP("event-json", "j", "", "json payload for the template")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateEventInput, err error) {
	input.EventType = datum.Config.String("type")
	if input.EventType == "" {
		return input, datum.NewRequiredFieldMissingError("type")
	}

	userIDs := datum.Config.Strings("user-ids")
	if userIDs != nil {
		input.UserIDs = userIDs
	}

	eventJSON := datum.Config.String("event-json")
	if eventJSON != "" {
		var data []byte

		if data, err = os.ReadFile(eventJSON); err != nil {
			cobra.CheckErr(err)
		}

		parsedMessage, err := datum.ParseBytes(data)
		cobra.CheckErr(err)

		input.Metadata = parsedMessage
	}

	metadata := datum.Config.String("metadata")
	if metadata != "" {
		err := json.Unmarshal([]byte(metadata), &input.Metadata)
		cobra.CheckErr(err)
	}

	return input, nil
}

// create a new datum event
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateEvent(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
