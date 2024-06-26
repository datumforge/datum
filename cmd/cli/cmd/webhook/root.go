package datumwebhooks

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base webhook command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "webhook",
	Short: "the subcommands for working with datum webhooks",
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
	case *datumclient.GetAllWebhooks:
		var nodes []*datumclient.GetAllWebhooks_Webhooks_Edges_Node

		for _, i := range v.Webhooks.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.CreateWebhook:
		e = v.CreateWebhook.Webhook
	case *datumclient.UpdateWebhook:
		e = v.UpdateWebhook.Webhook
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Webhook

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Webhook
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
func tableOutput(out []datumclient.Webhook) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "Description", "Destination URL", "Enabled")
	for _, i := range out {
		// this doesn't visually show you the json in the table but leaving it in for now
		writer.AddRow(i.ID, i.Name, *i.Description, i.DestinationURL, i.Enabled)
	}

	writer.Render()
}
