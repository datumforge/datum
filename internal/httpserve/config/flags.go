package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/datumforge/datum/internal/utils/viperconfig"
)

// RegisterServerFlags registers the flags for the server configuration
func RegisterServerFlags(v *viper.Viper, flags *pflag.FlagSet) error {
	err := viperconfig.BindConfigFlag(v, flags, "server.debug", "debug", false, "enable server debug", flags.Bool)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "server.listen", "listen", DefaultListenAddr, "address to listen on", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "server.https", "https", false, "enable serving from https", flags.Bool)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "server.ssl-cert", "ssl-cert", "", "ssl cert file location", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "server.ssl-key", "ssl-key", "", "ssl key file location", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "server.auto-cert", "auto-cert", false, "automatically generate tls cert", flags.Bool)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "server.cert-host", "cert-host", "example.com", "host to use to generate tls cert", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "server.shutdown-grace-period", "shutdown-grace-period", 0, "server shutdown grace periodt", flags.Duration)
	if err != nil {
		return err
	}

	return nil
}
