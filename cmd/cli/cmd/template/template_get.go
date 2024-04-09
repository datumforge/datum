package datumtemplates

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var templateGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details of existing datum templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		return templates(cmd.Context())
	},
}

func init() {
	templateCmd.AddCommand(templateGetCmd)

	templateGetCmd.Flags().BoolP("current", "c", false, "get template from currently logged in organization")
	datum.ViperBindFlag("template.get.current", templateGetCmd.Flags().Lookup("current"))

	templateGetCmd.Flags().StringP("id", "i", "", "get a specific template by ID")
	datum.ViperBindFlag("template.get.id", templateGetCmd.Flags().Lookup("id"))
}

func templates(ctx context.Context) error {
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	oID := viper.GetString("template.get.id")

	var s []byte

	current := viper.GetBool("template.get.current")

	writer := tables.NewTableWriter(templateCmd.OutOrStdout(), "ID", "Name", "Description", "JSON")

	if current {
		claims, err := tokens.ParseUnverifiedTokenClaims(cli.AccessToken)
		if err != nil {
			return err
		}

		oID = claims.ParseOrgID().String()
	}

	if oID != "" {
		template, err := cli.Client.GetTemplate(ctx, oID, cli.Interceptor)
		if err != nil {
			return err
		}

		if viper.GetString("output.format") == "json" {
			s, err := json.Marshal(template.Template)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}
		// this doesn't visually show you the json in the table but leaving it in for now
		writer.AddRow(template.Template.ID, template.Template.Name, *template.Template.Description, template.Template.Jsonconfig)
		writer.Render()

		return nil
	}

	templates, err := cli.Client.GetAllTemplates(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(templates.Templates)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		return datum.JSONPrint(s)
	}
	// table writer doesn't visually show details of the json (it shows as bytes) but leaving in for now
	for _, template := range templates.Templates.Edges {
		writer.AddRow(template.Node.ID, template.Node.Name, *template.Node.Description, template.Node.Jsonconfig)
	}

	writer.Render()

	return nil
}
