package register

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/httpserve/handlers"
)

var registerUserCmd = &cobra.Command{
	Use:   "create",
	Short: "Register a new datum user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return registerUser(cmd.Context())
	},
}

func init() {
	registerCmd.AddCommand(registerUserCmd)

	registerUserCmd.Flags().StringP("email", "e", "", "email of the user")
	datum.ViperBindFlag("register.create.email", registerUserCmd.Flags().Lookup("email"))

	registerUserCmd.Flags().StringP("password", "p", "", "password of the user")
	datum.ViperBindFlag("register.create.password", registerUserCmd.Flags().Lookup("password"))

	registerUserCmd.Flags().StringP("first-name", "f", "", "first name of the user")
	datum.ViperBindFlag("register.create.first-name", registerUserCmd.Flags().Lookup("first-name"))

	registerUserCmd.Flags().StringP("last-name", "l", "", "last name of the user")
	datum.ViperBindFlag("register.create.last-name", registerUserCmd.Flags().Lookup("last-name"))
}

func registerUser(ctx context.Context) error {
	var s []byte

	email := viper.GetString("register.create.email")
	if email == "" {
		return datum.NewRequiredFieldMissingError("email")
	}

	firstName := viper.GetString("register.create.first-name")
	if firstName == "" {
		return datum.NewRequiredFieldMissingError("first name")
	}

	lastName := viper.GetString("register.create.last-name")
	if lastName == "" {
		return datum.NewRequiredFieldMissingError("last name")
	}

	password := viper.GetString("register.create.password")
	if password == "" {
		return datum.NewRequiredFieldMissingError("password")
	}

	register := handlers.RegisterRequest{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	}

	// setup datum http client
	h := &http.Client{}

	// set options
	opt := &clientv2.Options{}

	// new client with params
	c := datumclient.NewClient(h, datum.DatumHost, opt, nil)

	// this allows the use of the graph client to be used for the REST endpoints
	dc := c.(*datumclient.Client)

	registration, err := datumclient.Register(dc, ctx, register)
	if err != nil {
		return err
	}

	s, err = json.Marshal(registration)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
