package datumentitlement

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base entitlement command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "entitlement",
	Short: "the subcommands for working with datum entitlements",
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
	case *datumclient.GetEntitlements:
		var nodes []*datumclient.GetEntitlements_Entitlements_Edges_Node

		for _, i := range v.Entitlements.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetEntitlementByID:
		e = v.Entitlement
	case *datumclient.CreateEntitlement:
		e = v.CreateEntitlement.Entitlement
	case *datumclient.UpdateEntitlement:
		e = v.UpdateEntitlement.Entitlement
	case *datumclient.DeleteEntitlement:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var entitlementList []datumclient.Entitlement

	err = json.Unmarshal(s, &entitlementList)
	if err != nil {
		var entitlement datumclient.Entitlement
		err = json.Unmarshal(s, &entitlement)
		cobra.CheckErr(err)

		entitlementList = append(entitlementList, entitlement)
	}

	tableOutput(entitlementList)

	return nil
}

// jsonOutput prints the output in a JSON format
func jsonOutput(out any) error {
	s, err := json.Marshal(out)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}

// tableOutput prints the entitlements in a table format
func tableOutput(out []datumclient.Entitlement) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "PlanID", "ExpiresAt", "Expires", "Cancelled")

	for _, i := range out {
		planID := ""
		if i.Plan != nil {
			planID = i.Plan.ID
		}

		writer.AddRow(i.ID, planID, i.ExpiresAt, i.Expires, i.Cancelled)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteEntitlement) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteEntitlement.DeletedID)

	writer.Render()
}
