package fga

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/datumforge/datum/internal/utils/viperconfig"
)

const (
	defaultFGAScheme = "https"
)

// RegisterFGAFlags registers the flags for the openFGA configuration
func RegisterFGAFlags(v *viper.Viper, flags *pflag.FlagSet) error {
	err := viperconfig.BindConfigFlag(v, flags, "fga.host", "fga-host", "", "fga host without the scheme (e.g. api.fga.example instead of https://api.fga.example)", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "fga.scheme", "fga-scheme", defaultFGAScheme, "fga scheme (http vs. https)", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "fga.store.id", "fga-store-id", "", "fga scheme Store ID", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "fga.model.id", "fga-model-id", "", "fga authorization model ID", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "fga.model.create", "fga-model-create", false, "force create a fga authorization model, this should be used when a model exists, but transitioning to a new model", flags.Bool)
	if err != nil {
		return err
	}

	return nil
}
