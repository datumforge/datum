// Package datum is our cobra/viper cli implementation
package datum

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/TylerBrock/colorjson"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	appName          = "datum"
	defaultGraphHost = "http://localhost:17608/query"
)

var (
	cfgFile string
	host    string
	logger  *zap.SugaredLogger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "the datum cli",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/."+appName+".yaml)")
	viperBindFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().StringVar(&host, "host", defaultGraphHost, "graph api host url")
	viperBindFlag("datum.host", rootCmd.PersistentFlags().Lookup("host"))

	// Logging flags
	rootCmd.PersistentFlags().Bool("debug", false, "enable debug logging")
	viperBindFlag("logging.debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.PersistentFlags().Bool("pretty", false, "enable pretty (human readable) logging output")
	viperBindFlag("logging.pretty", rootCmd.PersistentFlags().Lookup("pretty"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".datum" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".datum")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetEnvPrefix("datum")
	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()

	setupLogging()

	if err == nil {
		logger.Infow("using config file", "file", viper.ConfigFileUsed())
	}
}

func setupLogging() {
	cfg := zap.NewProductionConfig()
	if viper.GetBool("logging.pretty") {
		cfg = zap.NewDevelopmentConfig()
	}

	if viper.GetBool("logging.debug") {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	logger = l.Sugar().With("app", appName)
	defer logger.Sync() //nolint:errcheck
}

// viperBindFlag provides a wrapper around the viper bindings that panics if an error occurs
func viperBindFlag(name string, flag *pflag.Flag) {
	err := viper.BindPFlag(name, flag)
	if err != nil {
		panic(err)
	}
}

func jsonPrint(s []byte) error {
	var obj map[string]interface{}

	err := json.Unmarshal(s, &obj)
	if err != nil {
		return err
	}

	f := colorjson.NewFormatter()
	f.Indent = 2

	o, err := f.Marshal(obj)
	if err != nil {
		return err
	}

	fmt.Println(string(o))

	return nil
}
