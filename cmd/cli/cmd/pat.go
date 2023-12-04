package datum

import "github.com/spf13/cobra"

// patCmd represents the base patCmd command when called without any subcommands
var patCmd = &cobra.Command{
	Use:   "pat",
	Short: "The subcommands for working with personal access tokens",
}

func init() {
	rootCmd.AddCommand(patCmd)
}
