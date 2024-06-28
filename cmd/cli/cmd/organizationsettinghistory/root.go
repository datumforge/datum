package datumorganizationsettinghistory

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base organizationSettingHistory command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "organization-setting-history",
	Short: "the subcommands for working with datum organizationSettingHistories",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the output in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the organizationSettingHistories in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the organizationSettingHistories and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetAllOrganizationSettingHistories:
		var nodes []*datumclient.GetAllOrganizationSettingHistories_OrganizationSettingHistories_Edges_Node

		for _, i := range v.OrganizationSettingHistories.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetOrganizationSettingHistories:
		var nodes []*datumclient.GetOrganizationSettingHistories_OrganizationSettingHistories_Edges_Node

		for _, i := range v.OrganizationSettingHistories.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.OrganizationSettingHistory

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.OrganizationSettingHistory
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
func tableOutput(out []datumclient.OrganizationSettingHistory) {
	// create a table writer
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Ref", "Operation", "UpdatedAt", "UpdatedBy")
	for _, i := range out {
		writer.AddRow(i.ID, *i.Ref, i.Operation, *i.UpdatedAt, *i.UpdatedBy)
	}

	writer.Render()
}
