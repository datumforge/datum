package datumtemplates

import (
	"context"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/pkg/datumclient"
)

var templateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum template",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createTemplate(cmd.Context())
	},
}

func init() {
	templateCmd.AddCommand(templateCreateCmd)

	templateCreateCmd.Flags().StringP("name", "n", "", "name of the template")
	templateCreateCmd.Flags().StringP("description", "d", "", "description of the template")
	templateCreateCmd.Flags().StringP("json-config", "j", "", "json payload for the template")
	templateCreateCmd.Flags().StringP("ui-schema", "u", "", "ui schema for the template")
	templateCreateCmd.Flags().StringP("type", "t", "DOCUMENT", "type of the template")
}

func createTemplate(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := datum.Config.String("name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("template name")
	}

	description := datum.Config.String("description")
	jsonconfig := datum.Config.String("json-config")
	templateType := datum.Config.String("type")
	uischema := datum.Config.String("ui-schema")

	var data []byte

	if data, err = os.ReadFile(jsonconfig); err != nil {
		return err
	}

	input := datumclient.CreateTemplateInput{
		Name:       name,
		Jsonconfig: data,
	}

	if description != "" {
		input.Description = &description
	}

	if uischema != "" {
		var data []byte

		if data, err = os.ReadFile(uischema); err != nil {
			return err
		}

		input.Uischema = data
	}

	if templateType != "" {
		input.TemplateType = enums.ToDocumentType(templateType)
	}

	o, err := client.CreateTemplate(ctx, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
