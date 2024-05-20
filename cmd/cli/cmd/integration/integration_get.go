package datumintegrations

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var integrationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details of existing datum integrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return integrations(cmd.Context())
	},
}

func init() {
	integrationCmd.AddCommand(integrationGetCmd)

	integrationGetCmd.Flags().StringP("id", "i", "", "get a specific integration by ID")
	datum.ViperBindFlag("integration.get.id", integrationGetCmd.Flags().Lookup("id"))
}

func integrations(ctx context.Context) error {
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	oID := viper.GetString("integration.get.id")

	var s []byte

	writer := tables.NewTableWriter(integrationCmd.OutOrStdout(), "OwnerID", "Name", "Description", "kind", "Webhook ID", "Webhook URL")

	if oID != "" {
		integration, err := cli.Client.GetIntegrationByID(ctx, oID, cli.Interceptor)
		if err != nil {
			return err
		}

		if datum.OutputFormat == datum.JSONOutput {
			s, err := json.Marshal(integration.Integration)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer.AddRow(integration.Integration.OwnerID, integration.Integration.Name, *integration.Integration.Description, integration.Integration.Kind)
		writer.Render()

		return nil
	}

	integrations, err := cli.Client.GetAllIntegrations(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(integrations.Integrations)
	if err != nil {
		return err
	}

	if datum.OutputFormat == datum.JSONOutput {
		return datum.JSONPrint(s)
	}

	for _, integration := range integrations.Integrations.Edges {
		writer.AddRow(integration.Node.OwnerID, integration.Node.Name, *integration.Node.Description, *integration.Node.Kind)
	}

	writer.Render()

	return nil
}
