package datumentity

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum entity",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	// command line flags for the create command
	createCmd.Flags().StringP("name", "n", "", "name of the entity")
	createCmd.Flags().StringP("display-name", "s", "", "human friendly name of the entity")
	createCmd.Flags().StringP("type", "t", "", "type of the entity")
	createCmd.Flags().StringP("description", "d", "", "description of the entity")
}

// createValidation validates the required fields for the command
func createValidation(ctx context.Context) (input datumclient.CreateEntityInput, err error) {
	// validation of required fields for the create command
	input.Name = datum.Config.String("name")
	if input.Name == "" {
		return input, datum.NewRequiredFieldMissingError("entity name")
	}

	entityType := datum.Config.String("type")
	if entityType != "" {
		// get the entity type id
		id, err := getEntityTypeID(ctx, entityType)
		cobra.CheckErr(err)

		fmt.Println("Entity Type ID: ", id)

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

	return input, nil
}

// create a new datum entity
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation(ctx)
	cobra.CheckErr(err)

	o, err := client.CreateEntity(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}

func getEntityTypeID(ctx context.Context, name string) (string, error) {
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)

	where := &datumclient.EntityTypeWhereInput{
		Name: &name,
	}

	o, err := client.GetEntityTypes(ctx, where)
	cobra.CheckErr(err)

	if len(o.EntityTypes.Edges) == 0 || len(o.EntityTypes.Edges) > 1 {
		return "", fmt.Errorf("%w: entity type '%s' not found", datum.ErrNotFound, name)
	}

	return o.EntityTypes.Edges[0].Node.ID, nil
}
