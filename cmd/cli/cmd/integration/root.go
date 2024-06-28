package datumintegrations

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base integration command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "integration",
	Short: "the subcommands for working with datum integrations",
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
	case *datumclient.GetAllIntegrations:
		var nodes []*datumclient.GetAllIntegrations_Integrations_Edges_Node

		for _, i := range v.Integrations.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetIntegrationByID:
		e = v.Integration
	case *datumclient.CreateIntegration:
		e = v.CreateIntegration.Integration
	case *datumclient.UpdateIntegration:
		e = v.UpdateIntegration.Integration
	case *datumclient.DeleteIntegration:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Integration

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Integration
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
func tableOutput(out []datumclient.Integration) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "Description", "Kind")
	for _, i := range out {
		writer.AddRow(i.ID, i.Name, *i.Description, i.Kind)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteIntegration) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteIntegration.DeletedID)

	writer.Render()
}
