package datumtemplates

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base template command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "template",
	Short: "the subcommands for working with the datum templates",
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
	case *datumclient.GetAllTemplates:
		var nodes []*datumclient.GetAllTemplates_Templates_Edges_Node

		for _, i := range v.Templates.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetTemplateByID:
		e = v.Template
	case *datumclient.CreateTemplate:
		e = v.CreateTemplate.Template
	case *datumclient.UpdateTemplate:
		e = v.UpdateTemplate.Template
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Template

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Template
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
func tableOutput(out []datumclient.Template) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "Description", "JSON")
	for _, i := range out {
		// this doesn't visually show you the json in the table but leaving it in for now
		writer.AddRow(i.ID, i.Name, *i.Description, i.Jsonconfig)
	}

	writer.Render()
}
