package datumevents

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var eventGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details of existing datum events",
	RunE: func(cmd *cobra.Command, args []string) error {
		return events(cmd.Context())
	},
}

func init() {
	eventCmd.AddCommand(eventGetCmd)

	eventGetCmd.Flags().StringP("id", "i", "", "get a specific event by ID")
	datum.ViperBindFlag("event.get.id", eventGetCmd.Flags().Lookup("id"))
}

func events(ctx context.Context) error {
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	eventID := viper.GetString("event.get.id")
	if eventID != "" {
		event, err := cli.Client.GetEventByID(ctx, eventID, cli.Interceptor)
		if err != nil {
			return err
		}

		if datum.OutputFormat == "json" {
			s, err = json.Marshal(event)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer := tables.NewTableWriter(eventCmd.OutOrStdout(), "ID", "EventType", "EventMetadata", "CorrelationID")
		writer.AddRow(event.Event.ID, event.Event.EventType, event.Event.Metadata, event.Event.CorrelationID)

		writer.Render()

		return nil
	}

	writer := tables.NewTableWriter(eventCmd.OutOrStdout(), "ID", "EventType", "EventMetadata", "CorrelationID")

	events, err := cli.Client.GetAllEvents(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(events.Events)
	if err != nil {
		return err
	}

	if datum.OutputFormat == "json" {
		return datum.JSONPrint(s)
	}

	for _, event := range events.Events.Edges {
		writer.AddRow(event.Node.ID, event.Node.EventType, event.Node.Metadata, event.Node.CorrelationID)
	}

	writer.Render()

	return nil
}
