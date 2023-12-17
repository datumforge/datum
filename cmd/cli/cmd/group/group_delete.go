package datumgroup

import (
	"context"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
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
	datum.ViperBindFlag("group.delete.id", groupDeleteCmd.Flags().Lookup("id"))
}

func deleteGroup(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	var s []byte

	gID := viper.GetString("group.delete.id")
	if gID == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	o, err := cli.Client.DeleteGroup(ctx, gID, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
