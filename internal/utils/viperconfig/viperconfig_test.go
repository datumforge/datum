package viperconfig_test

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	utils "github.com/datumforge/datum/internal/utils/viperconfig"
)

const (
	viperPath  = "test.path"
	cmdLineArg = "test-arg"
	help       = "test help"
)

func TestBindConfigFlagStringWithArg(t *testing.T) {
	t.Parallel()

	v := viper.New()
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	defaultValue := "test"

	err := utils.BindConfigFlag(
		v, flags, viperPath, cmdLineArg, defaultValue,
		help, flags.String)

	require.NoError(t, err, "Unexpected error")

	// Check that the flags are registered
	require.NoError(t, flags.Parse([]string{"--" + cmdLineArg + "=foo"}))
	require.Equal(t, "foo", v.GetString(viperPath))
}

func TestBindConfigFlagStringWithDefaultArg(t *testing.T) {
	t.Parallel()

	v := viper.New()
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	defaultValue := "test"

	err := utils.BindConfigFlag(
		v, flags, viperPath, cmdLineArg, defaultValue,
		help, flags.String)

	require.NoError(t, err, "Unexpected error")

	// Check that the flags are registered
	require.NoError(t, flags.Parse([]string{}))
	require.Equal(t, defaultValue, v.GetString(viperPath))
}

func TestBindConfigFlagIntWithArg(t *testing.T) {
	t.Parallel()

	v := viper.New()
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	defaultValue := 123

	err := utils.BindConfigFlag(
		v, flags, viperPath, cmdLineArg, defaultValue,
		help, flags.Int)

	require.NoError(t, err, "Unexpected error")

	// Check that the flags are registered
	require.NoError(t, flags.Parse([]string{"--" + cmdLineArg + "=456"}))
	require.Equal(t, 456, v.GetInt(viperPath))
}

func TestBindConfigFlagIntWithDefaultArg(t *testing.T) {
	t.Parallel()

	v := viper.New()
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	defaultValue := 123

	err := utils.BindConfigFlag(
		v, flags, viperPath, cmdLineArg, defaultValue,
		help, flags.Int)

	require.NoError(t, err, "Unexpected error")

	// Check that the flags are registered
	require.NoError(t, flags.Parse([]string{}))
	require.Equal(t, 123, v.GetInt(viperPath))
}
