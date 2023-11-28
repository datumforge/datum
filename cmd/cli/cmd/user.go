package datum

import "github.com/spf13/cobra"

// userCmd represents the base user command when called without any subcommands
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "The subcommands for working with the datum user",
}

func init() {
	rootCmd.AddCommand(userCmd)
}
