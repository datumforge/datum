package datumorgsetting

import (
	"encoding/json"
	"strings"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base org setting command when called without any subcommands
var cmd = &cobra.Command{
	Use:     "organization-setting",
	Aliases: []string{"org-setting"},
	Short:   "the subcommands for working with the datum organization settings",
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
	case *datumclient.GetOrganizationSettings:
		var nodes []*datumclient.GetOrganizationSettings_OrganizationSettings_Edges_Node

		for _, i := range v.OrganizationSettings.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetAllOrganizationSettings:
		var nodes []*datumclient.GetAllOrganizationSettings_OrganizationSettings_Edges_Node

		for _, i := range v.OrganizationSettings.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetOrganizationSettingByID:
		e = v.OrganizationSetting
	case *datumclient.UpdateOrganizationSetting:
		e = v.UpdateOrganizationSetting.OrganizationSetting
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.OrganizationSetting

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.OrganizationSetting
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
func tableOutput(out []datumclient.OrganizationSetting) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(),
		"ID",
		"OrganizationName",
		"BillingContact",
		"BillingAddress",
		"BillingEmail",
		"BillingPhone",
		"GeoLocation",
		"TaxIdentifier",
		"Tags",
		"Domains",
	)
	for _, i := range out {
		writer.AddRow(i.ID,
			i.Organization.Name,
			*i.BillingContact,
			*i.BillingAddress,
			*i.BillingEmail,
			*i.BillingPhone,
			*i.GeoLocation,
			*i.TaxIdentifier,
			strings.Join(i.Tags, ", "),
			strings.Join(i.Domains, ", "))
	}

	writer.Render()
}
