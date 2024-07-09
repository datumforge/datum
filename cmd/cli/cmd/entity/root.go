package datumentity

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base entity command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "entity",
	Short: "the subcommands for working with datum entities",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the output in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the entities in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the entities and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetAllEntities:
		var nodes []*datumclient.GetAllEntities_Entities_Edges_Node

		for _, i := range v.Entities.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetEntities:
		var nodes []*datumclient.GetEntities_Entities_Edges_Node

		for _, i := range v.Entities.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetEntityByID:
		e = v.Entity
	case *datumclient.CreateEntity:
		e = v.CreateEntity.Entity
	case *datumclient.UpdateEntity:
		e = v.UpdateEntity.Entity
	case *datumclient.DeleteEntity:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Entity

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Entity
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
func tableOutput(out []datumclient.Entity) {
	// create a table writer
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "DisplayName", "Description", "EntityType")
	for _, i := range out {
		entityTypeName := ""
		if i.EntityType != nil {
			entityTypeName = i.EntityType.Name
		}

		writer.AddRow(i.ID, i.Name, i.DisplayName, *i.Description, entityTypeName)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteEntity) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteEntity.DeletedID)

	writer.Render()
}
