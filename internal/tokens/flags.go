package tokens

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/datumforge/datum/internal/utils/viperconfig"
)

const (
	defaultJWKSRemoteTimeout = 5 * time.Second
	defaultAccessDuration    = 1 * time.Hour
	defaultRefreshDuration   = 2 * time.Hour
	defaultRefreshOverlap    = -15 * time.Minute
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

	err = viperconfig.BindConfigFlag(v, flags, "jwt.access-duration", "jwks-access-duration", defaultAccessDuration, "length of time the access token is valid", flags.Duration)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "jwt.refresh-duration", "jwks-refresh-duration", defaultRefreshDuration, "length of time the refresh token is valid", flags.Duration)
	if err != nil {
		return err
	}

	err = viperconfig.BindConfigFlag(v, flags, "jwt.refresh-overlap", "jwks-refresh-overlap", defaultRefreshOverlap, "overlap duration between refresh and access token expiration", flags.Duration)
	if err != nil {
		return err
	}

	return nil
}
