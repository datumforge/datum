package datumgroup

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base group command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "group",
	Short: "the subcommands for working with datum groups",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the output in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the groups in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the groups and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetAllGroups:
		var nodes []*datumclient.GetAllGroups_Groups_Edges_Node

		for _, i := range v.Groups.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetGroupByID:
		e = v.Group
	case *datumclient.CreateGroup:
		e = v.CreateGroup.Group
	case *datumclient.UpdateGroup:
		e = v.UpdateGroup.Group
	case *datumclient.DeleteGroup:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Group

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Group
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
func tableOutput(out []datumclient.Group) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "Description", "Visibility", "Organization", "Members")
	for _, i := range out {
		writer.AddRow(i.ID, i.DisplayName, *i.Description, i.Setting.Visibility, i.Owner.DisplayName, len(i.Members))
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteGroup) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteGroup.DeletedID)

	writer.Render()
}
