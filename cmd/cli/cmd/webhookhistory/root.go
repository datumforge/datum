package datumwebhookhistory

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base webhookHistory command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "webhook-history",
	Short: "the subcommands for working with datum webhookHistories",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the output in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the webhookHistories in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the webhookHistories and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetAllWebhookHistories:
		var nodes []*datumclient.GetAllWebhookHistories_WebhookHistories_Edges_Node

		for _, i := range v.WebhookHistories.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetWebhookHistories:
		var nodes []*datumclient.GetWebhookHistories_WebhookHistories_Edges_Node

		for _, i := range v.WebhookHistories.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.WebhookHistory

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.WebhookHistory
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
func tableOutput(out []datumclient.WebhookHistory) {
	// create a table writer
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Ref", "Operation", "UpdatedAt", "UpdatedBy")
	for _, i := range out {
		writer.AddRow(i.ID, *i.Ref, i.Operation, *i.UpdatedAt, *i.UpdatedBy)
	}

	writer.Render()
}
