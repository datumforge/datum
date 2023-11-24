package datum

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/Yamashou/gqlgenc/clientv2"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/datumforge/datum/internal/datumclient"
)

var orgDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteOrg(cmd.Context())
	},
}

func init() {
	orgCmd.AddCommand(orgDeleteCmd)

	orgDeleteCmd.Flags().StringP("id", "i", "", "org id to delete")
	viperBindFlag("org.delete.id", orgDeleteCmd.Flags().Lookup("id"))
}

func deleteOrg(ctx context.Context) error {
	// setup datum http client
	h := &http.Client{}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	token := os.Getenv("DATUM_ACCESS_TOKEN")

	i := datumclient.WithAccessToken(token)

	// new client with params
	c := datumclient.NewClient(h, host, opt, i)

	var s []byte

	oID := viper.GetString("org.delete.id")
	if oID == "" {
		return ErrOrgIDRequired
	}

	o, err := c.DeleteOrganization(ctx, oID, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return jsonPrint(s)
}
