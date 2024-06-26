package datumfeature

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base feature command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "feature",
	Short: "the subcommands for working with datum features",
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
	case *datumclient.GetFeatures:
		var nodes []*datumclient.GetFeatures_Features_Edges_Node

		for _, i := range v.Features.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.CreateFeature:
		e = v.CreateFeature.Feature
	case *datumclient.UpdateFeature:
		e = v.UpdateFeature.Feature
	case *datumclient.DeleteFeature:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Feature

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Feature
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
func tableOutput(out []datumclient.Feature) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "DisplayName", "Enabled", "Description")

	for _, i := range out {
		desc := ""
		if i.Description != nil {
			desc = *i.Description
		}

		writer.AddRow(i.ID, i.Name, *i.DisplayName, i.Enabled, desc)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteFeature) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteFeature.DeletedID)

	writer.Render()
}
