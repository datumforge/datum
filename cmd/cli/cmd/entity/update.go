package datumentity

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/enums"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum entity",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "entity id to update")

	// command line flags for the update command
	updateCmd.Flags().StringP("name", "n", "", "name of the entity")
	updateCmd.Flags().StringP("display-name", "s", "", "human friendly name of the entity")
	updateCmd.Flags().StringP("type", "t", "", "type of the entity")
	updateCmd.Flags().StringP("description", "d", "", "description of the entity")
	updateCmd.Flags().StringSliceP("contacts", "c", []string{}, "contacts to associate with the entity")
}

// updateValidation validates the required fields for the command
func updateValidation() (id string, input datumclient.UpdateEntityInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("entity id")
	}

	// validation of required fields for the update command
	name := datum.Config.String("name")
	if name != "" {
		input.Name = &name
	}

	entityType := datum.Config.String("type")
	if entityType != "" {
		entityType := enums.EntityType(entityType)
		input.EntityType = &entityType
	}

	displayName := datum.Config.String("display-name")
	if displayName != "" {
		input.DisplayName = &displayName
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	if len(datum.Config.Strings("contacts")) > 0 {
		input.AddContactIDs = datum.Config.Strings("contacts")
	}

	return id, input, nil
}

// update an existing entity in the datum platform
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdateEntity(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
