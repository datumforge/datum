package datumtemplates

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
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

	templateID := viper.GetString("template.get.id")

	var s []byte

	writer := tables.NewTableWriter(templateCmd.OutOrStdout(), "ID", "Name", "Description", "JSON")

	if templateID != "" {
		template, err := cli.Client.GetTemplate(ctx, templateID, cli.Interceptor)
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
