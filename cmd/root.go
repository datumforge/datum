package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/datumforge/datum/pkg/utils/sentry"
)

const appName = "datum"

var (
	logger *zap.SugaredLogger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "A datum repo for graph apis",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().Bool("pretty", false, "enable pretty (human readable) logging output")
	viperBindFlag("pretty", rootCmd.PersistentFlags().Lookup("pretty"))

	rootCmd.PersistentFlags().Bool("debug", false, "debug logging output")
	viperBindFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// initConfig reads in flags set for server startup
// all other configuration is done by the server with koanf
// refer to the README.md for more information
func initConfig() {
	err := viper.ReadInConfig()

	logger = sentry.NewLogger()

	if err == nil {
		logger.Infow("using config file", "file", viper.ConfigFileUsed())
	}
}

// viperBindFlag provides a wrapper around the viper bindings that panics if an error occurs
func viperBindFlag(name string, flag *pflag.Flag) {
	err := viper.BindPFlag(name, flag)
	if err != nil {
		panic(err)
	}
}
