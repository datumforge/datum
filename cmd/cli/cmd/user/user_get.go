package datumuser

import (
	"context"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/tokens"
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

	userGetCmd.Flags().BoolP("self", "s", false, "get current user info, requires authentication")
	datum.ViperBindFlag("user.get.self", userGetCmd.Flags().Lookup("self"))

	userGetCmd.Flags().StringP("id", "i", "", "user id to query")
	datum.ViperBindFlag("user.get.id", userGetCmd.Flags().Lookup("id"))
}

func users(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	// filter options
	userID := viper.GetString("user.get.id")
	self := viper.GetBool("user.get.self")

	var s []byte

	if self {
		claims, err := tokens.ParseUnverifiedTokenClaims(cli.AccessToken)
		if err != nil {
			return err
		}

		userID = claims.ParseUserID().String()
	}

	// if a user ID is provided, filter on that user, otherwise get all
	if userID == "" {
		users, err := cli.Client.GetAllUsers(ctx, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err = json.Marshal(users)
		if err != nil {
			return err
		}
	} else {
		user, err := cli.Client.GetUserByID(ctx, userID, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err = json.Marshal(user)
		if err != nil {
			return err
		}
	}

	return datum.JSONPrint(s)
}
