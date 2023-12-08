package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/tokens"
)

func TestDefaultAuthOptions(t *testing.T) {
	// Should be able to create a default auth options with no extra input.
	conf := auth.NewAuthOptions()
	require.NotZero(t, conf, "a zero valued configuration was returned")
	require.Equal(t, auth.DefaultKeysURL, conf.KeysURL)
	require.Equal(t, auth.DefaultAudience, conf.Audience)
	require.Equal(t, auth.DefaultIssuer, conf.Issuer)
	require.Equal(t, auth.DefaultMinRefreshInterval, conf.MinRefreshInterval)
	require.NotNil(t, conf.Context, "no context was created")
}

func TestAuthOptions(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	conf := auth.NewAuthOptions(
		auth.WithJWKSEndpoint("http://localhost:8088/.well-known/jwks.json"),
		auth.WithAudience("http://localhost:3000"),
		auth.WithIssuer("http://localhost:8088"),
		auth.WithMinRefreshInterval(67*time.Minute),
		auth.WithContext(ctx),
	)

	cancel()
	require.NotZero(t, conf, "a zero valued configuration was returned")
	require.Equal(t, "http://localhost:8088/.well-known/jwks.json", conf.KeysURL)
	require.Equal(t, "http://localhost:3000", conf.Audience)
	require.Equal(t, "http://localhost:8088", conf.Issuer)
	require.Equal(t, 67*time.Minute, conf.MinRefreshInterval)
	require.ErrorIs(t, conf.Context.Err(), context.Canceled)
}

func TestAuthOptionsOverride(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	opts := auth.AuthOptions{
		KeysURL:            "http://localhost:8088/.well-known/jwks.json",
		Audience:           "http://localhost:3000",
		Issuer:             "http://localhost:8088",
		MinRefreshInterval: 42 * time.Minute,
		Context:            ctx,
	}

	conf := auth.NewAuthOptions(auth.WithAuthOptions(opts))
	require.NotSame(t, opts, conf, "expected a new configuration object to be created")
	require.Equal(t, opts, conf, "expected the opts to override the configuration defaults")

	// Ensure the context is the same on the configuration
	cancel()
	require.ErrorIs(t, conf.Context.Err(), context.Canceled)
}

func TestAuthOptionsValidator(t *testing.T) {
	validator := &tokens.MockValidator{}
	conf := auth.NewAuthOptions(auth.WithValidator(validator))
	require.NotZero(t, conf, "a zero valued configuration was returned")

	actual, err := conf.Validator()
	require.NoError(t, err, "could not create default validator")
	require.Same(t, validator, actual, "conf did not return the same validator")
}
