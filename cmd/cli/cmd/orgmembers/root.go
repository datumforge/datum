package datumorgmembers

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base orgMembers command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "org-members",
	Short: "the subcommands for working with the members of a datum organization",
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
	case *datumclient.GetOrgMembersByOrgID:
		var nodes []*datumclient.GetOrgMembersByOrgID_OrgMemberships_Edges_Node

		for _, i := range v.OrgMemberships.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.AddUserToOrgWithRole:
		e = v.CreateOrgMembership.OrgMembership
	case *datumclient.UpdateUserRoleInOrg:
		e = v.UpdateOrgMembership.OrgMembership
	case *datumclient.RemoveUserFromOrg:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.OrgMembership

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.OrgMembership
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
func tableOutput(out []datumclient.OrgMembership) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "UserID", "DisplayName", "FirstName", "LastName", "Email", "Role")
	for _, i := range out {
		writer.AddRow(i.UserID, i.User.DisplayName, *i.User.FirstName, *i.User.LastName, i.User.Email, i.Role)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.RemoveUserFromOrg) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteOrgMembership.DeletedID)

	writer.Render()
}
