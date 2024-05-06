package datumevents

import (
	"context"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	datum.ViperBindFlag("event.create.type", eventCreateCmd.Flags().Lookup("type"))

	eventCreateCmd.Flags().StringP("metadata", "m", "", "metadata for the event")
	datum.ViperBindFlag("event.create.metadata", eventCreateCmd.Flags().Lookup("metadata"))

	eventCreateCmd.Flags().StringP("userid", "u", "", "user id associated with the event")
	datum.ViperBindFlag("event.create.userid", eventCreateCmd.Flags().Lookup("userid"))

	eventCreateCmd.Flags().StringP("eventjson", "j", "", "json payload for the template")
	datum.ViperBindFlag("event.create.metadata", eventCreateCmd.Flags().Lookup("eventjson"))
}

func createevent(ctx context.Context) error {
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	eventType := viper.GetString("event.create.type")
	if eventType == "" {
		return datum.NewRequiredFieldMissingError("type")
	}

	userid := viper.GetStringSlice("event.create.userid")
	eventjson := viper.GetString("event.create.eventjson")
	metadata := viper.GetString("event.create.metadata")

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
		if err != nil {
			return err
		}

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

	u, err := cli.Client.CreateEvent(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(u)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
