package datumentitytype

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base entity type command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "entity-type",
	Short: "the subcommands for working with datum entity types",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the output in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the entity types in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the entity types and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetAllEntityTypes:
		var nodes []*datumclient.GetAllEntityTypes_EntityTypes_Edges_Node

		for _, i := range v.EntityTypes.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetEntityTypes:
		var nodes []*datumclient.GetEntityTypes_EntityTypes_Edges_Node

		for _, i := range v.EntityTypes.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetEntityTypeByID:
		e = v.EntityType
	case *datumclient.CreateEntityType:
		e = v.CreateEntityType.EntityType
	case *datumclient.UpdateEntityType:
		e = v.UpdateEntityType.EntityType
	case *datumclient.DeleteEntityType:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.EntityType

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.EntityType
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
func tableOutput(out []datumclient.EntityType) {
	// create a table writer
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name")
	for _, i := range out {
		writer.AddRow(i.ID, i.Name)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteEntityType) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteEntityType.DeletedID)

	writer.Render()
}
