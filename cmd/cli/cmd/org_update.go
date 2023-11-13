package datum

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/datumforge/datum/internal/datumclient"
)

var orgUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateOrg(cmd.Context())
	},
}

func init() {
	orgCmd.AddCommand(orgUpdateCmd)

	orgUpdateCmd.Flags().StringP("id", "i", "", "org id to update")
	viperBindFlag("org.update.id", orgUpdateCmd.Flags().Lookup("id"))

	orgUpdateCmd.Flags().StringP("name", "n", "", "name of the organization")
	viperBindFlag("org.update.name", orgUpdateCmd.Flags().Lookup("name"))

	orgUpdateCmd.Flags().StringP("short-name", "s", "", "display name of the organization")
	viperBindFlag("org.update.short-name", orgUpdateCmd.Flags().Lookup("short-name"))

	orgUpdateCmd.Flags().StringP("description", "d", "", "description of the organization")
	viperBindFlag("org.update.description", orgUpdateCmd.Flags().Lookup("description"))
}

func updateOrg(ctx context.Context) error {
	// setup datum http client
	h := &http.Client{}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	i := func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		// TODO: Add Auth Headers
		return next(ctx, req, gqlInfo, res)
	}

	// new client with params
	c := datumclient.NewClient(h, host, opt, i)

	var s []byte

	oID := viper.GetString("org.update.id")
	if oID == "" {
		return ErrOrgIDRequired
	}

	name := viper.GetString("org.update.name")
	displayName := viper.GetString("org.update.short-name")
	description := viper.GetString("org.update.description")

	input := datumclient.UpdateOrganizationInput{}

	if name != "" {
		input.Name = &name
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := c.UpdateOrganization(ctx, oID, input, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return jsonPrint(s)
}
