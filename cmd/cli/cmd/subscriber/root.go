package datumsubscribers

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base subscribers command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "subscriber",
	Short: "the subcommands for working with the subscribers of a datum organization",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the out in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the out in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the out and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetSubscribers:
		var nodes []*datumclient.GetSubscribers_Subscribers_Edges_Node

		for _, i := range v.Subscribers.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetAllSubscribers:
		var nodes []*datumclient.GetAllSubscribers_Subscribers_Edges_Node

		for _, i := range v.Subscribers.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetSubscriberByEmail:
		e = v.Subscriber
	case *datumclient.CreateBulkSubscriber:
		e = v.CreateBulkSubscriber.Subscribers
	case *datumclient.CreateBulkCSVSubscriber:
		e = v.CreateBulkCSVSubscriber.Subscribers
	case *datumclient.CreateSubscriber:
		e = v.CreateSubscriber.Subscriber
	case *datumclient.DeleteSubscriber:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Subscriber

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Subscriber
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
func tableOutput(out []datumclient.Subscriber) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Email", "Verified", "Active")
	for _, i := range out {
		writer.AddRow(i.ID, i.Email, i.VerifiedEmail, i.Active)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted email in a table format
func deletedTableOutput(e *datumclient.DeleteSubscriber) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteSubscriber.Email)

	writer.Render()
}
