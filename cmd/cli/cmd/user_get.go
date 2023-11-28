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

var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details of existing datum user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return users(cmd.Context())
	},
}

func init() {
	userCmd.AddCommand(userGetCmd)

	// TODO: implement get self once we know the sub -> user ID
	// orgGetCmd.Flags().BoolP("current-user", "c", false, "get current user info")
	// viperBindFlag("user.get.current", orgGetCmd.Flags().Lookup("current-user"))

	userGetCmd.Flags().StringP("id", "i", "", "user id to query")
	viperBindFlag("user.get.id", userGetCmd.Flags().Lookup("id"))
}

func users(ctx context.Context) error {
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

	// filter options
	userID := viper.GetString("user.get.id")
	// self := viper.GetBool("user.get.current")

	var s []byte

	// if a user ID is provided, filter on that user, otherwise get all
	if userID == "" {
		users, err := c.GetAllUsers(ctx, i)
		if err != nil {
			return err
		}

		s, err = json.Marshal(users)
		if err != nil {
			return err
		}
	} else {
		user, err := c.GetUserByID(ctx, userID, i)
		if err != nil {
			return err
		}

		s, err = json.Marshal(user)
		if err != nil {
			return err
		}
	}

	return jsonPrint(s)
}
