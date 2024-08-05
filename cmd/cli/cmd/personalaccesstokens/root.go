package datumtokens

import (
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

// cmd represents the base cmd command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "pat",
	Short: "the subcommands for working with personal access tokens",
}

func init() {
	datum.RootCmd.AddCommand(cmd)
}

// consoleOutput prints the output in the console
func consoleOutput(e any) error {
	// check if the output format is JSON and print the pat in JSON format
	if datum.OutputFormat == datum.JSONOutput {
		return jsonOutput(e)
	}

	// check the type of the pat and print them in a table format
	switch v := e.(type) {
	case *datumclient.GetAllPersonalAccessTokens:
		var nodes []*datumclient.GetAllPersonalAccessTokens_PersonalAccessTokens_Edges_Node

		for _, i := range v.PersonalAccessTokens.Edges {
			nodes = append(nodes, i.Node)
		}

		e = nodes
	case *datumclient.GetPersonalAccessTokenByID:
		e = v.PersonalAccessToken
	case *datumclient.CreatePersonalAccessToken:
		e = v.CreatePersonalAccessToken.PersonalAccessToken
	case *datumclient.UpdatePersonalAccessToken:
		e = v.UpdatePersonalAccessToken.PersonalAccessToken
	case *datumclient.DeletePersonalAccessToken:
		deletedTableOutput(v)
		return nil
	}

	s, err := json.Marshal(e)
	cobra.CheckErr(err)

	var list []datumclient.PersonalAccessToken

	err = json.Unmarshal(s, &list)
	if err != nil {
		var in datumclient.PersonalAccessToken
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
func tableOutput(out []datumclient.PersonalAccessToken) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "ID", "Name", "Token", "LastUsedAt", "ExpiresAt")

	for _, i := range out {
		lastUsed := "never"
		if i.LastUsedAt != nil {
			lastUsed = i.LastUsedAt.String()
		}

		expiresAt := "never"
		if i.ExpiresAt != nil {
			expiresAt = i.ExpiresAt.String()
		}

		writer.AddRow(i.ID, i.Name, i.Token, lastUsed, expiresAt)
	}

	writer.Render()
}

// deleteTableOutput prints the deleted id in a table format
func deletedTableOutput(e *datumclient.DeletePersonalAccessToken) {
	writer := tables.NewTableWriter(cmd.OutOrStdout(), "DeletedID")

	writer.AddRow(e.DeletePersonalAccessToken.DeletedID)

	writer.Render()
}
