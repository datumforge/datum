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

var patDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum personal access token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deletePat(cmd.Context())
	},
}

func init() {
	patCmd.AddCommand(patDeleteCmd)

	patDeleteCmd.Flags().StringP("id", "i", "", "pat id to delete")
	viperBindFlag("pat.delete.id", patDeleteCmd.Flags().Lookup("id"))
}

func deletePat(ctx context.Context) error {
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

	oID := viper.GetString("pat.delete.id")
	if oID == "" {
		return ErrTokenIDRequired
	}

	o, err := c.DeletePersonalAccessToken(ctx, oID, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return jsonPrint(s)
}
