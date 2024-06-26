package datumentitlementplan

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base plan command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "plan",
	Short: "the subcommands for working with datum entitlement plans",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the plans in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the plans in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the plans and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetEntitlementPlans:
		var nodes []*datumclient.GetEntitlementPlans_EntitlementPlans_Edges_Node

		for _, i := range v.EntitlementPlans.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetEntitlementPlanByID:
		e = v.EntitlementPlan
	case *datumclient.CreateEntitlementPlan:
		e = v.CreateEntitlementPlan.EntitlementPlan
	case *datumclient.UpdateEntitlementPlan:
		e = v.UpdateEntitlementPlan.EntitlementPlan
	case *datumclient.DeleteEntitlementPlan:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var planList []datumclient.EntitlementPlan

	err = json.Unmarshal(s, &planList)
	if err != nil {
		var plan datumclient.EntitlementPlan
		err = json.Unmarshal(s, &plan)
		cobra.CheckErr(err)

		planList = append(planList, plan)
	}

	tableOutput(planList)

	return nil
}

// jsonOutput prints the output in a JSON format
func jsonOutput(out any) error {
	s, err := json.Marshal(out)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}

// tableOutput prints the plans in a table format
func tableOutput(plans []datumclient.EntitlementPlan) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "Display Name", "Description", "Version")

	for _, p := range plans {
		writer.AddRow(p.ID, p.Name, *p.DisplayName, *p.Description, p.Version)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted plan in a table format
func deletedTableOutput(e *datumclient.DeleteEntitlementPlan) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteEntitlementPlan.DeletedID)

	writer.Render()
}
