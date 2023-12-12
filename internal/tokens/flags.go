package tokens

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/datumforge/datum/internal/utils/viperconfig"
)

const (
	defaultJWKSRemoteTimeout = 5 * time.Second
)

// RegisterAuthFlags registers the flags for the authentication configuration
func RegisterAuthFlags(v *viper.Viper, flags *pflag.FlagSet) error {
	// Auth Flags
	err := viperconfig.BindConfigFlag(v, flags, "auth", "auth", true, "enable authentication checks", flags.Bool)
	if err != nil {
		return err
	}

	// OIDC Flags
	err = viperconfig.BindConfigFlag(v, flags, "jwt.audience", "jwt-audience", "", "expected audience on datum JWT", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "jwt.issuer", "jwt-issuer", "", "expected issuer of datum JWT", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "jwt.cookie-domain", "jwt-cookie-domain", "datum.net", "cookie domain for datum JWT", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "jwt.kid", "jwt-kid", "", "kid for the JWT keys", flags.String)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "jwks.remote-timeout", "jwks-remote-timeout", defaultJWKSRemoteTimeout, "timeout for remote JWKS fetching", flags.Duration)
	if err != nil {
		return err
	}

	return nil
}
