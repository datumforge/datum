package datumuser

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base user command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "user",
	Short: "the subcommands for working with the datum user",
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
	case *datumclient.GetAllUsers:
		var nodes []*datumclient.GetAllUsers_Users_Edges_Node

		for _, i := range v.Users.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetUserByID:
		e = v.User
	case *datumclient.CreateUser:
		e = v.CreateUser.User
	case *datumclient.UpdateUser:
		e = v.UpdateUser.User
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.User

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.User
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
func tableOutput(out []datumclient.User) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Email", "FirstName", "LastName", "DisplayName", "AuthProvider")
	for _, i := range out {
		// this doesn't visually show you the json in the table but leaving it in for now
		writer.AddRow(i.ID, i.Email, *i.FirstName, *i.LastName, i.DisplayName, i.AuthProvider)
	}

	writer.Render()
}
