package datumusersetting

import (
	"encoding/json"
	"strings"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base user setting command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "user-setting",
	Short: "the subcommands for working with the datum user settings",
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
	case *datumclient.GetAllUserSettings:
		var nodes []*datumclient.GetAllUserSettings_UserSettings_Edges_Node

		for _, i := range v.UserSettings.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetUserSettingByID:
		e = v.UserSetting
	case *datumclient.UpdateUserSetting:
		e = v.UpdateUserSetting.UserSetting
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.UserSetting

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.UserSetting
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
func tableOutput(out []datumclient.UserSetting) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "DefaultOrgName", "DefaultOrgID", "2FA", "Status", "EmailConfirmed", "Tags")

	for _, i := range out {
		defaultOrgName := ""
		defaultOrgID := ""

		if i.DefaultOrg != nil {
			defaultOrgName = i.DefaultOrg.Name
			defaultOrgID = i.DefaultOrg.ID
		}

		writer.AddRow(i.ID, defaultOrgName, defaultOrgID, *i.IsTfaEnabled, i.Status, i.EmailConfirmed, strings.Join(i.Tags, ", "))
	}

	writer.Render()
}
