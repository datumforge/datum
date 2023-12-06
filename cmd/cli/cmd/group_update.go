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

var groupUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupUpdateCmd)

	groupUpdateCmd.Flags().StringP("id", "i", "", "group id to update")
	viperBindFlag("group.update.id", groupUpdateCmd.Flags().Lookup("id"))

	groupUpdateCmd.Flags().StringP("name", "n", "", "name of the group")
	viperBindFlag("group.update.name", groupUpdateCmd.Flags().Lookup("name"))

	groupUpdateCmd.Flags().StringP("short-name", "s", "", "display name of the group")
	viperBindFlag("group.update.short-name", groupUpdateCmd.Flags().Lookup("short-name"))

	groupUpdateCmd.Flags().StringP("description", "d", "", "description of the group")
	viperBindFlag("group.update.description", groupUpdateCmd.Flags().Lookup("description"))
}

func updateGroup(ctx context.Context) error {
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

	oID := viper.GetString("group.update.id")
	if oID == "" {
		return ErrGroupIDRequired
	}

	name := viper.GetString("group.update.name")
	displayName := viper.GetString("group.update.short-name")
	description := viper.GetString("group.update.description")

	input := datumclient.UpdateGroupInput{}

	if name != "" {
		input.Name = &name
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := c.UpdateGroup(ctx, oID, input, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return jsonPrint(s)
}
