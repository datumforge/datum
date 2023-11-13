package datum

import "github.com/spf13/cobra"

// orgCmd represents the base org command when called without any subcommands
var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "The subcommands for working with the datum organiation",
}

func init() {
	rootCmd.AddCommand(orgCmd)
}
