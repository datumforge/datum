package datumtemplates

import (
	"context"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/tokens"
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
	datum.ViperBindFlag("template.create.name", templateCreateCmd.Flags().Lookup("name"))

	templateCreateCmd.Flags().StringP("description", "d", "", "description of the template")
	datum.ViperBindFlag("template.create.description", templateCreateCmd.Flags().Lookup("description"))

	templateCreateCmd.Flags().StringP("org-id", "o", "", "organization id, leave empty to create within current org")
	datum.ViperBindFlag("template.create.org-id", templateCreateCmd.Flags().Lookup("org-id"))

	templateCreateCmd.Flags().StringP("jsonconfig", "j", "", "json payload for the template")
	datum.ViperBindFlag("template.create.jsonconfig", templateCreateCmd.Flags().Lookup("jsonconfig"))

	templateCreateCmd.Flags().StringP("uischema", "u", "", "uischema for the template")
	datum.ViperBindFlag("template.create.uischema", templateCreateCmd.Flags().Lookup("uischema"))

	templateCreateCmd.Flags().StringP("type", "t", "DOCUMENT", "type of the template")
	datum.ViperBindFlag("template.create.type", templateCreateCmd.Flags().Lookup("type"))
}

func createTemplate(ctx context.Context) error {
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := viper.GetString("template.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("template name")
	}

	description := viper.GetString("template.create.description")
	parentOrgID := viper.GetString("template.create.org-id")
	jsonconfig := viper.GetString("template.create.jsonconfig")
	templateType := viper.GetString("template.create.type")
	uischema := viper.GetString("template.create.uischema")

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

	if parentOrgID != "" {
		input.OwnerID = parentOrgID
	}

	if templateType != "" {
		input.TemplateType = enums.ToDocumentType(templateType)
	}

	if parentOrgID == "" {
		claims, err := tokens.ParseUnverifiedTokenClaims(cli.AccessToken)
		if err != nil {
			return err
		}

		oID := claims.ParseOrgID().String()

		input.OwnerID = oID
	}

	o, err := cli.Client.CreateTemplate(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
