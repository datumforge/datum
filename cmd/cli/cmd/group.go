package datum

import "github.com/spf13/cobra"

// orgCmd represents the base org command when called without any subcommands
var groupCmd = &cobra.Command{
	Use:   "org",
	Short: "The subcommands for working with the datum organization",
}

func init() {
	rootCmd.AddCommand(groupCmd)
}
