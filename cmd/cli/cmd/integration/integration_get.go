package datumintegrations

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
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
}

func integrations(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	oID := datum.Config.String("id")

	var s []byte

	writer := tables.NewTableWriter(integrationCmd.OutOrStdout(), "OwnerID", "Name", "Description", "kind", "Webhook ID", "Webhook URL")

	if oID != "" {
		integration, err := client.GetIntegrationByID(ctx, oID)
		cobra.CheckErr(err)

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

	integrations, err := client.GetAllIntegrations(ctx)
	cobra.CheckErr(err)

	s, err = json.Marshal(integrations.Integrations)
	cobra.CheckErr(err)

	if datum.OutputFormat == datum.JSONOutput {
		return datum.JSONPrint(s)
	}

	for _, integration := range integrations.Integrations.Edges {
		writer.AddRow(integration.Node.OwnerID, integration.Node.Name, *integration.Node.Description, *integration.Node.Kind)
	}

	writer.Render()

	return nil
}
