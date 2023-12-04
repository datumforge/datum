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

var patCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum personal access token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createPat(cmd.Context())
	},
}

func init() {
	patCmd.AddCommand(patCreateCmd)

	patCreateCmd.Flags().StringP("name", "n", "", "name of the personal access token")
	viperBindFlag("pat.create.name", patCreateCmd.Flags().Lookup("name"))

	patCreateCmd.Flags().StringP("description", "d", "", "description of the pat")
	viperBindFlag("pat.create.description", patCreateCmd.Flags().Lookup("description"))

	patCreateCmd.Flags().StringP("owner-id", "o", "", "the owner of the personal access token")
	viperBindFlag("pat.create.owner-id", patCreateCmd.Flags().Lookup("owner-id"))
}

func createPat(ctx context.Context) error {
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

	name := viper.GetString("pat.create.name")
	if name == "" {
		return ErrTokenNameRequired
	}

	owner := viper.GetString("pat.create.owner-id")
	if owner == "" {
		return ErrUserIDRequired
	}

	description := viper.GetString("pat.create.description")

	input := datumclient.CreatePersonalAccessTokenInput{
		Name:    name,
		OwnerID: owner,
	}

	if description != "" {
		input.Description = &description
	}

	o, err := c.CreatePersonalAccessToken(ctx, input, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return jsonPrint(s)
}
