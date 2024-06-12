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
	Short: "Create a new datum event",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createevent(cmd.Context())
	},
}

func init() {
	eventCmd.AddCommand(eventCreateCmd)

	eventCreateCmd.Flags().StringP("type", "t", "", "type of the event")
	eventCreateCmd.Flags().StringP("metadata", "m", "", "metadata for the event")
	eventCreateCmd.Flags().StringSliceP("user-ids", "u", []string{}, "user id associated with the event")
	eventCreateCmd.Flags().StringP("event-json", "j", "", "json payload for the template")
}

func createevent(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)

	defer datum.StoreSessionCookies(client)

	var s []byte

	eventType := datum.Config.String("type")
	if eventType == "" {
		return datum.NewRequiredFieldMissingError("type")
	}

	userid := datum.Config.Strings("user-ids")
	eventjson := datum.Config.String("event-json")
	metadata := datum.Config.String("metadata")

	input := datumclient.CreateEventInput{
		EventType: eventType,
	}

	if userid != nil {
		input.UserIDs = userid
	}

	if eventjson != "" {
		var data []byte

		if data, err = os.ReadFile(eventjson); err != nil {
			return err
		}

		parsedMessage, err := datum.ParseBytes(data)
		cobra.CheckErr(err)

		input.Metadata = parsedMessage
	}

	if metadata != "" {
		if err := json.Unmarshal([]byte(metadata), &input.Metadata); err != nil {
			return err
		}
	}

	if userid != nil {
		input.UserIDs = userid
	}

	u, err := client.CreateEvent(ctx, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(u)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
