package viperconfig

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// FlagInst is a function that creates a flag and returns a pointer
type FlagInst[V any] func(name string, value V, usage string) *V

// FlagInstShort is a function that creates a flag and returns a pointer
type FlagInstShort[V any] func(name, shorthand string, value V, usage string) *V

// BindConfigFlag is a helper function that binds a configuration value to a flag
// v: The viper.Viper object used to retrieve the configuration value
// flags: The pflag.FlagSet object used to retrieve the flag value
// viperPath: The path used to retrieve the configuration value from Viper
// cmdLineArg: The flag name used to check if the flag has been set and to retrieve its value
// help: The help text for the flag
// defaultValue: A default value used to determine the type of the flag (string, int, etc.)
// binder: A function that creates a flag and returns a pointer to the value
func BindConfigFlag[V any](
	v *viper.Viper,
	flags *pflag.FlagSet,
	viperPath string,
	cmdLineArg string,
	defaultValue V,
	help string,
	binder FlagInst[V],
) error {
	binder(cmdLineArg, defaultValue, help)
	return doViperBind[V](v, flags, viperPath, cmdLineArg, defaultValue)
}

// BindConfigFlagWithShort is a helper function that binds a configuration value to a flag
// v: The viper.Viper object used to retrieve the configuration value
// flags: The pflag.FlagSet object used to retrieve the flag value
// viperPath: The path used to retrieve the configuration value from Viper
// cmdLineArg: The flag name used to check if the flag has been set and to retrieve its value
// short: The short name for the flag
// help: The help text for the flag
// defaultValue: A default value used to determine the type of the flag (string, int, etc.)
// binder: A function that creates a flag and returns a pointer to the value
func BindConfigFlagWithShort[V any](
	v *viper.Viper,
	flags *pflag.FlagSet,
	viperPath string,
	cmdLineArg string,
	short string,
	defaultValue V,
	help string,
	binder FlagInstShort[V],
) error {
	binder(cmdLineArg, short, defaultValue, help)
	return doViperBind[V](v, flags, viperPath, cmdLineArg, defaultValue)
}

func doViperBind[V any](
	v *viper.Viper,
	flags *pflag.FlagSet,
	viperPath string,
	cmdLineArg string,
	defaultValue V,
) error {
	v.SetDefault(viperPath, defaultValue)

	if err := v.BindPFlag(viperPath, flags.Lookup(cmdLineArg)); err != nil {
		return fmt.Errorf("failed to bind flag %s to viper path %s: %w", cmdLineArg, viperPath, err)
	}

	return nil
}
