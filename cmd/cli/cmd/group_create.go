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

var groupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupCreateCmd)

	groupCreateCmd.Flags().StringP("name", "n", "", "name of the group")
	viperBindFlag("group.create.name", groupCreateCmd.Flags().Lookup("name"))

	groupCreateCmd.Flags().StringP("short-name", "s", "", "display name of the group")
	viperBindFlag("group.create.short-name", groupCreateCmd.Flags().Lookup("short-name"))

	groupCreateCmd.Flags().StringP("description", "d", "", "description of the group")
	viperBindFlag("group.create.description", groupCreateCmd.Flags().Lookup("description"))

	groupCreateCmd.Flags().StringP("owner-id", "o", "", "owner org id")
	viperBindFlag("group.create.owner-id", groupCreateCmd.Flags().Lookup("owner-id"))
}

func createGroup(ctx context.Context) error {
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

	name := viper.GetString("group.create.name")
	if name == "" {
		return ErrGroupNameRequired
	}

	owner := viper.GetString("group.create.owner-id")
	if owner == "" {
		return ErrOrgIDRequired
	}

	displayName := viper.GetString("group.create.short-name")
	description := viper.GetString("group.create.description")

	input := datumclient.CreateGroupInput{
		Name:    name,
		OwnerID: owner,
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := c.CreateGroup(ctx, input, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return jsonPrint(s)
}
