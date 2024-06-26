package datumapitokens

import (
	"encoding/json"
	"strings"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base cmd command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "token",
	Short: "the subcommands for working with api tokens",
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
	case *datumclient.GetAllAPITokens:
		var nodes []*datumclient.GetAllAPITokens_APITokens_Edges_Node

		for _, i := range v.APITokens.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.CreateAPIToken:
		e = v.CreateAPIToken.APIToken
	case *datumclient.UpdateAPIToken:
		e = v.UpdateAPIToken.APIToken
	case *datumclient.DeleteAPIToken:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.APIToken

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.APIToken
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
func tableOutput(out []datumclient.APIToken) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "Token", "Scopes", "LastUsedAt", "ExpiresAt")

	for _, i := range out {
		lastUsed := "never"
		if i.LastUsedAt != nil {
			lastUsed = i.LastUsedAt.String()
		}

		writer.AddRow(i.ID, i.Name, i.Token, strings.Join(i.Scopes, ", "), lastUsed, i.ExpiresAt)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeleteAPIToken) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeleteAPIToken.DeletedID)

	writer.Render()
}
