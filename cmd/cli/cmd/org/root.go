package datumorg

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base org command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "org",
	Short: "the subcommands for working with the datum organization",
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
	case *datumclient.GetAllOrganizations:
		var nodes []*datumclient.GetAllOrganizations_Organizations_Edges_Node

		for _, i := range v.Organizations.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.CreateOrganization:
		e = v.CreateOrganization.Organization
	case *datumclient.UpdateOrganization:
		e = v.UpdateOrganization.Organization
	case *datumclient.DeleteOrganization:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Organization

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Organization
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
func tableOutput(out []datumclient.Organization) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "Description", "PersonalOrg", "Children", "Members")

	for _, i := range out {
		childrenLen := 0
		if i.Children != nil {
			childrenLen = len(i.Children.Edges)
		}

		writer.AddRow(i.ID, i.DisplayName, *i.Description, *i.PersonalOrg, childrenLen, len(i.Members))
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteOrganization) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteOrganization.DeletedID)

	writer.Render()
}
