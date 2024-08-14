package datumentity

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
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
	updateCmd.Flags().StringSliceP("contacts", "c", []string{}, "contact IDs to associate with the entity")
	updateCmd.Flags().StringSlice("domains", []string{}, "domains associated with the entity")
	updateCmd.Flags().String("note", "", "add note about the entity")
	updateCmd.Flags().String("status", "", "status of the entity")
}

// updateValidation validates the required fields for the command
func updateValidation(ctx context.Context) (id string, input datumclient.UpdateEntityInput, err error) {
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
		id, err := getEntityTypeID(ctx, entityType)
		cobra.CheckErr(err)

		input.EntityTypeID = &id
	}

	displayName := datum.Config.String("display-name")
	if displayName != "" {
		input.DisplayName = &displayName
	}

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	contacts := datum.Config.Strings("contacts")
	if len(contacts) > 0 {
		input.AddContactIDs = contacts
	}

	domains := datum.Config.Strings("domains")
	if len(domains) > 0 {
		input.AppendDomains = domains
	}

	note := datum.Config.String("note")
	if note != "" {
		input.Note = &datumclient.CreateNoteInput{
			Text: note,
		}
	}

	status := datum.Config.String("status")
	if status != "" {
		input.Status = &status
	}

	return id, input, nil
}

// update an existing entity in the datum platform
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation(ctx)
	cobra.CheckErr(err)

	o, err := client.UpdateEntity(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
