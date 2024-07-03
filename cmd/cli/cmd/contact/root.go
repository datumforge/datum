package datumcontact

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base contact command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "contact",
	Short: "the subcommands for working with datum contacts",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the output in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the contacts in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the contacts and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetAllContacts:
		var nodes []*datumclient.GetAllContacts_Contacts_Edges_Node

		for _, i := range v.Contacts.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetContacts:
		var nodes []*datumclient.GetContacts_Contacts_Edges_Node

		for _, i := range v.Contacts.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetContactByID:
		e = v.Contact
	case *datumclient.CreateContact:
		e = v.CreateContact.Contact
	case *datumclient.UpdateContact:
		e = v.UpdateContact.Contact
	case *datumclient.DeleteContact:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.Contact

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.Contact
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
func tableOutput(out []datumclient.Contact) {
	// create a table writer
	// TODO: add additional columns to the table writer
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "Status", "Company", "Title", "Email", "PhoneNumber", "Address")
	for _, i := range out {
		writer.AddRow(i.ID, i.FullName, i.Status.String(), *i.Company, *i.Title, *i.Email, *i.PhoneNumber, *i.Address)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteContact) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteContact.DeletedID)

	writer.Render()
}
