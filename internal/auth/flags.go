package auth

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/datumforge/datum/internal/utils/viperconfig"
)

const (
	defaultOIDCJWKSRemoteTimeout = 5 * time.Second
)

// RegisterAuthFlags registers the flags for the authentication configuration
func RegisterAuthFlags(v *viper.Viper, flags *pflag.FlagSet) error {
	// echo-jwt flags
	err := viperconfig.BindConfigFlag(v, flags, "jwt.secretkey", "jwt-secretkey", "", "secret key for echojwt config", flags.String)
	if err != nil {
		return err
	}

	// OIDC Flags
	err = viperconfig.BindConfigFlag(v, flags, "oidc.enabled", "oidc", true, "enable authentication checks", flags.Bool)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "oidc.audience", "oidc-audience", "", "expected audience on OIDC JWT", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "oidc.issuer", "oidc-issuer", "", "expected issuer of OIDC JWT", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "oidc.jwks.remote-timeout", "oidc-jwks-remote-timeout", defaultOIDCJWKSRemoteTimeout, "timeout for remote JWKS fetching", flags.Duration)
	if err != nil {
		return err
	}

	return nil
}
