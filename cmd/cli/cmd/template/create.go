package datumtemplates

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/enums"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum template",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "name of the template")
	createCmd.Flags().StringP("description", "d", "", "description of the template")
	createCmd.Flags().StringP("json-config", "j", "", "json payload for the template")
	createCmd.Flags().StringP("ui-schema", "u", "", "ui schema for the template")
	createCmd.Flags().StringP("type", "t", "DOCUMENT", "type of the template")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateTemplateInput, err error) {
	input.Name = datum.Config.String("name")
	if input.Name == "" {
		return input, datum.NewRequiredFieldMissingError("template name")
	}

	jsonConfig := datum.Config.String("json-config")

	data, err := os.ReadFile(jsonConfig)
	cobra.CheckErr(err)

	input.Jsonconfig = data

	description := datum.Config.String("description")
	if description != "" {
		input.Description = &description
	}

	uiSchema := datum.Config.String("ui-schema")
	if uiSchema != "" {
		data, err = os.ReadFile(uiSchema)
		cobra.CheckErr(err)

		input.Uischema = data
	}

	templateType := datum.Config.String("type")
	if templateType != "" {
		input.TemplateType = enums.ToDocumentType(templateType)
	}

	return input, nil
}

// create a new template in datum
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateTemplate(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
