package datuminvite

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base invite command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "invite",
	Short: "the subcommands for working with the invitations of a datum organization",
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
	case *datumclient.GetAllInvites:
		var nodes []*datumclient.GetAllInvites_Invites_Edges_Node

		for _, i := range v.Invites.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.CreateInvite:
		e = v.CreateInvite.Invite
	case *datumclient.DeleteInvite:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Invite

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Invite
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
func tableOutput(out []datumclient.Invite) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Recipient", "Role", "Status")
	for _, i := range out {
		writer.AddRow(i.ID, i.Recipient, i.Role, i.Status)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteInvite) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteInvite.DeletedID)

	writer.Render()
}
