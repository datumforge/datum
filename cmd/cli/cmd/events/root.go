package datumevents

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base event command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "event",
	Short: "the subcommands for working with datum events",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the output in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the output in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the output and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetAllEvents:
		var nodes []*datumclient.GetAllEvents_Events_Edges_Node

		for _, i := range v.Events.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetEventByID:
		e = v.Event
	case *datumclient.CreateEvent:
		e = v.CreateEvent.Event
	case *datumclient.UpdateEvent:
		e = v.UpdateEvent.Event
	case *datumclient.DeleteEvent:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Event

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Event
		err = json.Unmarshal(s, &in)
		cobra.CheckErr(err)

		list = append(list, in)
	}

	tableOutput(list)

	return nil
}

// jsonOutput prints the output in a JSON format
func jsonOutput(out any) error {
	s, err := json.Marshal(out)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}

// tableOutput prints the output in a table format
func tableOutput(out []datumclient.Event) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "EventType", "EventMetadata", "CorrelationID")

	for _, i := range out {
		writer.AddRow(i.ID, i.EventType, i.Metadata, i.CorrelationID)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted plan in a table format
func deletedTableOutput(e *datumclient.DeleteEvent) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteEvent.DeletedID)

	writer.Render()
}
