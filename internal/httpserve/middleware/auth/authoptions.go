package auth

import (
	"context"
	"regexp"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwk"

	"github.com/datumforge/datum/internal/tokens"
)

const (
	Authorization             = "Authorization"
	DefaultKeysURL            = "http://localhost:17608/.well-known/jwks.json"
	DefaultAudience           = "http://localhost:17608"
	DefaultIssuer             = "http://localhost:17608"
	DefaultMinRefreshInterval = 5 * time.Minute
	DefaultCookieDomain       = "localhost:17608"
	AccessTokenCookie         = "access_token"
	RefreshTokenCookie        = "refresh_token"
)

// used to extract the access token from the header
var (
	bearer = regexp.MustCompile(`^\s*[Bb]earer\s+([a-zA-Z0-9_\-\.]+)\s*$`)
)

// AuthOption allows users to optionally supply configuration to the Authorization middleware.
type AuthOption func(opts *AuthOptions)

// AuthOptions is constructed from variadic AuthOption arguments with reasonable defaults.
type AuthOptions struct {
	// KeysURL endpoint to the JWKS public keys on the datum server
	KeysURL string
	// Audience to verify on tokens
	Audience string
	// Issuer to verify on tokens
	Issuer string
	// MinRefreshInterval to cache the JWKS public keys
	MinRefreshInterval time.Duration
	// CookieDomain to use for auth cookies
	CookieDomain string
	// Context to control the lifecycle of the background fetch routine
	Context context.Context

	//  validator constructed by the auth options (can be directly supplied by the user).
	validator tokens.Validator
	// reauth constructed by the auth options (can be directly supplied by the user).
	reauth Reauthenticator
}

// Reauthenticator generates new access and refresh pair given a valid refresh token.
type Reauthenticator interface {
	Refresh(context.Context, *RefreshRequest) (*LoginReply, error)
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
	OrgID        string `json:"org_id,omitempty"`
}

type LoginReply struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	LastLogin    string `json:"last_login,omitempty"`
}

// NewAuthOptions creates an AuthOptions object with reasonable defaults and any user
// supplied input from the AuthOption variadic arguments.
func NewAuthOptions(opts ...AuthOption) (conf AuthOptions) {
	conf = AuthOptions{
		KeysURL:            DefaultKeysURL,
		Audience:           DefaultAudience,
		Issuer:             DefaultIssuer,
		MinRefreshInterval: DefaultMinRefreshInterval,
	}

	for _, opt := range opts {
		opt(&conf)
	}

	// Create a context if one has not been supplied by the user.
	if conf.Context == nil && conf.validator == nil {
		conf.Context = context.Background()
	}

	return conf
}

// Validator returns the user supplied validator or constructs a new JWKS Cache
// Validator from the supplied options. If the options are invalid or the validator
// cannot be created an error is returned
func (conf *AuthOptions) Validator() (tokens.Validator, error) {
	if conf.validator == nil {
		cache := jwk.NewCache(conf.Context)

		err := cache.Register(conf.KeysURL, jwk.WithMinRefreshInterval(conf.MinRefreshInterval))
		if err != nil {
			return nil, ErrShitWentBad
		}

		conf.validator, err = tokens.NewCachedJWKSValidator(conf.Context, cache, conf.KeysURL, conf.Audience, conf.Issuer)
		if err != nil {
			return nil, err
		}
	}

	return conf.validator, nil
}

// WithAuthOptions allows the user to update the default auth options with an auth
// options struct to set many options values at once. Zero values are ignored, so if
// using this option, the defaults will still be preserved if not set on the input.
func WithAuthOptions(opts AuthOptions) AuthOption {
	return func(conf *AuthOptions) {
		if opts.KeysURL != "" {
			conf.KeysURL = opts.KeysURL
		}

		if opts.Audience != "" {
			conf.Audience = opts.Audience
		}

		if opts.Issuer != "" {
			conf.Issuer = opts.Issuer
		}

		if opts.MinRefreshInterval != 0 {
			conf.MinRefreshInterval = opts.MinRefreshInterval
		}

		if opts.Context != nil {
			conf.Context = opts.Context
		}
	}
}

// WithJWKSEndpoint allows the user to specify an alternative endpoint to fetch the JWKS
// public keys from. This is useful for testing or for different environments.
func WithJWKSEndpoint(url string) AuthOption {
	return func(opts *AuthOptions) {
		opts.KeysURL = url
	}
}

// WithAudience allows the user to specify an alternative audience.
func WithAudience(audience string) AuthOption {
	return func(opts *AuthOptions) {
		opts.Audience = audience
	}
}

// WithIssuer allows the user to specify an alternative issuer.
func WithIssuer(issuer string) AuthOption {
	return func(opts *AuthOptions) {
		opts.Issuer = issuer
	}
}

// WithMinRefreshInterval allows the user to specify an alternative minimum duration
// between cache refreshes to control refresh behavior for the JWKS public keys.
func WithMinRefreshInterval(interval time.Duration) AuthOption {
	return func(opts *AuthOptions) {
		opts.MinRefreshInterval = interval
	}
}

// WithContext allows the user to specify an external, cancelable context to control
// the background refresh behavior of the JWKS cache.
func WithContext(ctx context.Context) AuthOption {
	return func(opts *AuthOptions) {
		opts.Context = ctx
	}
}

// WithValidator allows the user to specify an alternative validator to the auth
// middleware. This is particularly useful for testing authentication.
func WithValidator(validator tokens.Validator) AuthOption {
	return func(opts *AuthOptions) {
		opts.validator = validator
	}
}

// WithReauthenticator allows the user to specify a reauthenticator to the auth
// middleware.
func WithReauthenticator(reauth Reauthenticator) AuthOption {
	return func(opts *AuthOptions) {
		opts.reauth = reauth
	}
}
