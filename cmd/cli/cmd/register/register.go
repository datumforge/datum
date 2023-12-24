package register

import (
	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

// registerCmd represents the base register user command when called without any subcommands
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "The subcommands for working with datum registration",
}

func init() {
	datum.RootCmd.AddCommand(registerCmd)
}
