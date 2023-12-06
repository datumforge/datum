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

var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupDeleteCmd)

	groupDeleteCmd.Flags().StringP("id", "i", "", "group id to delete")
	viperBindFlag("group.delete.id", groupDeleteCmd.Flags().Lookup("id"))
}

func deleteGroup(ctx context.Context) error {
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

	oID := viper.GetString("group.delete.id")
	if oID == "" {
		return ErrGroupIDRequired
	}

	o, err := c.DeleteGroup(ctx, oID, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return jsonPrint(s)
}
