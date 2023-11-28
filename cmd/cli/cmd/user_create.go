package datum

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Yamashou/gqlgenc/clientv2"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/datumforge/datum/internal/datumclient"
)

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createUser(cmd.Context())
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)

	userCreateCmd.Flags().StringP("email", "e", "", "email of the user")
	viperBindFlag("user.create.email", userCreateCmd.Flags().Lookup("email"))

	userCreateCmd.Flags().StringP("first-name", "f", "", "first name of the user")
	viperBindFlag("user.create.first-name", userCreateCmd.Flags().Lookup("first-name"))

	userCreateCmd.Flags().StringP("last-name", "l", "", "last name of the user")
	viperBindFlag("user.create.last-name", userCreateCmd.Flags().Lookup("last-name"))

	userCreateCmd.Flags().StringP("display-name", "d", "", "first name of the user")
	viperBindFlag("user.create.display-name", userCreateCmd.Flags().Lookup("display-name"))
}

func createUser(ctx context.Context) error {
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

	email := viper.GetString("user.create.email")
	if email == "" {
		return ErrUserEmailRequired
	}

	firstName := viper.GetString("user.create.first-name")
	if firstName == "" {
		return ErrUserFirstNameRequired
	}

	lastName := viper.GetString("user.create.last-name")
	if lastName == "" {
		return ErrUserLastNameRequired
	}

	displayName := viper.GetString("user.create.display-name")
	if displayName == "" {
		// set a default display name if not set
		displayName = strings.ToLower(fmt.Sprintf("%s.%s", firstName, lastName))
	}

	input := datumclient.CreateUserInput{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	u, err := c.CreateUser(ctx, input, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(u)
	if err != nil {
		return err
	}

	return jsonPrint(s)
}
