package datumentitlementplanfeatures

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base feature command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "plan-feature",
	Short: "the subcommands for working with datum entitlement plan features",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the output in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the planFeatures in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the output and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetEntitlementPlanFeatures:
		var nodes []*datumclient.GetEntitlementPlanFeatures_EntitlementPlanFeatures_Edges_Node

		for _, i := range v.EntitlementPlanFeatures.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.CreateEntitlementPlanFeature:
		e = v.CreateEntitlementPlanFeature.EntitlementPlanFeature
	case *datumclient.UpdateEntitlementPlanFeature:
		e = v.UpdateEntitlementPlanFeature.EntitlementPlanFeature
	case *datumclient.DeleteEntitlementPlanFeature:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.EntitlementPlanFeature

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.EntitlementPlanFeature
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
func tableOutput(out []datumclient.EntitlementPlanFeature) {
	headers := []string{"ID", "PlanName", "FeatureName"}

	// check if the planFeatures is empty and print the headers only
	if len(out) == 0 {
		writer := tables.NewTableWriter(cmd.OutOrStdout(), headers...)
		writer.Render()

		return
	}

	// get the metadata keys from the first planFeature and add them to the headers
	for k := range out[0].Metadata {
		headers = append(headers, k)
	}

	writer := tables.NewTableWriter(cmd.OutOrStdout(), headers...)

	for _, f := range out {
		items := []interface{}{f.ID, f.Plan.Name, f.Feature.Name}

		// add the metadata values to the items
		for _, v := range f.Metadata {
			items = append(items, v)
		}

		writer.AddRow(items...)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteEntitlementPlanFeature) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteEntitlementPlanFeature.DeletedID)

	writer.Render()
}
