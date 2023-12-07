package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/tokens"
	api "github.com/datumforge/datum/internal/utils/responses"
	echo "github.com/datumforge/echox"
)

const (
	authorization             = "Authorization"
	ContextUserClaims         = "user_claims"
	ContextAccessToken        = "access_token"
	ContextRequestID          = "request_id"
	DefaultKeysURL            = "https://auth.datum.net/.well-known/jwks.json"
	DefaultAudience           = "https://datum.net"
	DefaultIssuer             = "https://auth.datum.net"
	DefaultMinRefreshInterval = 5 * time.Minute
	DefaultCookieDomain       = "datum.net"
	AccessTokenCookie         = "access_token"
	RefreshTokenCookie        = "refresh_token"
)

// used to extract the access token from the header
var (
	bearer = regexp.MustCompile(`^\s*[Bb]earer\s+([a-zA-Z0-9_\-\.]+)\s*$`)
	logger *zap.SugaredLogger
)

// AuthOption allows users to optionally supply configuration to the Authorization middleware.
type AuthOption func(opts *AuthOptions)

// AuthOptions is constructed from variadic AuthOption arguments with reasonable defaults.
type AuthOptions struct {
	KeysURL            string           // The URL endpoint to the JWKS public keys on the datum server
	Audience           string           // The audience to verify on tokens
	Issuer             string           // The issuer to verify on tokens
	MinRefreshInterval time.Duration    // Minimum amount of time the JWKS public keys are cached
	CookieDomain       string           // The domain to use for auth cookies
	Context            context.Context  // The context object to control the lifecycle of the background fetch routine
	validator          tokens.Validator // The validator constructed by the auth options (can be directly supplied by the user).
	reauth             Reauthenticator  // The refresher constructed by the auth options (can be directly supplied by the user).
}

// A Reauthenticator generates new access and refresh pair given a valid refresh token.
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

func Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		conf := NewAuthOptions()

		var validator tokens.Validator

		if validator, err = conf.Validator(); err != nil {
			return err
		}

		// Create a reauthenticator function to handle refresh tokens if they are provided.
		reauthenticate := Reauthenticate(conf, validator)

		return func(c echo.Context) error {
			var (
				err         error
				accessToken string
				claims      *tokens.Claims
			)

			// Get access token from the request, if not available then attempt to refresh
			// using the refresh token cookie.
			if accessToken, err = GetAccessToken(c); err != nil {
				switch {
				case errors.Is(err, ErrNoAuthorization):
					if accessToken, err = reauthenticate(c); err != nil {

						ErrorResponse(ErrAuthRequired)
						return err
					}
				default:
					ErrorResponse(ErrAuthRequired)
					return err
				}
			}

			// Verify the access token is authorized for use with datum and extract claims.
			if claims, err = validator.Verify(accessToken); err != nil {
				ErrorResponse(ErrAuthRequired)
				return err
			}

			// Add claims to context for use in downstream processing and continue handlers
			ctx := context.WithValue(c.Request().Context(), ContextUserClaims, claims)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
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
func (conf *AuthOptions) Validator() (_ tokens.Validator, err error) {
	if conf.validator == nil {
		cache := jwk.NewCache(conf.Context)
		if err := cache.Register(conf.KeysURL, jwk.WithMinRefreshInterval(conf.MinRefreshInterval)); err != nil {
			return nil, fmt.Errorf("shit went bad")
		}

		if conf.validator, err = tokens.NewCachedJWKSValidator(conf.Context, cache, conf.KeysURL, conf.Audience, conf.Issuer); err != nil {
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

// Reauthenticate is a middleware helper that can use refresh tokens in the echo context
// to obtain a new access token. If it is unable to obtain a new valid access token,
// then an error is returned and processing should stop.
func Reauthenticate(conf AuthOptions, validator tokens.Validator) func(c *echo.Context) (string, error) {
	// If no reauthenticator is available on the configuration, always return an error.
	if conf.reauth == nil {
		return func(c *echo.Context) (string, error) {
			return "", ErrRefreshDisabled
		}
	}

	// If the reauthenticator is available, return a function that utilizes it.
	return func(c *echo.Context) (_ string, err error) {
		// Get the refresh token from the cookies or the headers of the request.

		refreshToken := GetRefreshToken(c)

		if err != nil {
			return "", err
		}

		// Check to ensure the refresh token is still valid.
		if _, err = validator.Verify(refreshToken); err != nil {
			return "", err
		}

		// Reauthenticate using the refresh token.
		req := &RefreshRequest{RefreshToken: refreshToken}
		reply, err := conf.reauth.Refresh(c.Request.Context(), req)
		if err != nil {
			return "", err
		}

		// Set the new access and refresh cookies
		if err = SetAuthCookies(c, reply.AccessToken, reply.RefreshToken, conf.CookieDomain); err != nil {
			return "", err
		}

		return reply.AccessToken, nil
	}
}

// GetAccessToken retrieves the bearer token from the authorization header and parses it
// to return only the JWT access token component of the header. Alternatively, if the
// authorization header is not present, then the token is fetched from cookies. If the
// header is missing or the token is not available, an error is returned.
//
// NOTE: the authorization header takes precedence over access tokens in cookies.
func GetAccessToken(c echo.Context) (tks string, err error) {
	// Attempt to get the access token from the header.

	if origin := c.Request().Header.Get("Origin"); len(origin) == 0 {
		c.Request().Header.Set("Origin", "*")
	}

	if header := echo.Context.Request().Header.Get(authorization); header != "" {
		match := bearer.FindStringSubmatch(header)
		if len(match) == 2 {
			return match[1], nil
		}
		return "", ErrParseBearer
	}

	// Attempt to get the access token from cookies.
	var cookie string

	if cookie, err = c.SetCookie(AccessTokenCookie); err == nil {
		// If the error is nil, that means we were able to retrieve the access token cookie
		return cookie, nil
	}
	return "", ErrNoAuthorization
}

// GetRefreshToken retrieves the refresh token from the cookies in the request. If the
// cookie is not present or expired then an error is returned.
func GetRefreshToken(c echo.Context) (tks string, err error) {
	rftc := c.Cookie(RefreshTokenCookie)
	if tks, err = c.Cookie(RefreshTokenCookie); err != nil {
		return "", ErrNoRefreshToken
	}
	return tks, nil
}

// GetClaims fetches and parses datum claims from the echo context. Returns an
// error if no claims exist on the context; panics if the claims are not the correct
// type -- however the panic should be recovered by middleware.
func GetClaims(c *echo.Context) (*tokens.Claims, error) {
	claims, exists := c.Get(ContextUserClaims)
	if !exists {
		return nil, ErrNoClaims
	}
	return claims.(*tokens.Claims), nil
}

// ContextFromRequest creates a context from the echo request context, copying fields
// that may be required for forwarded requests. This method should be called by
// handlers which need to forward requests to other services and need to preserve data
// from the original request such as the user's credentials.
func ContextFromRequest(c *echo.Context) (ctx context.Context, err error) {
	var req *http.Request
	if req = c.Request; req == nil {
		return nil, ErrNoRequest
	}

	// Add access token to context (from either header or cookie using Authenticate middleware)
	ctx = req.Context()
	if token := c.GetString(ContextAccessToken); token != "" {
		ctx = api.ContextWithToken(ctx, token)
	}

	// Add request id to context
	if requestID := c.GetString(ContextRequestID); requestID != "" {
		ctx = api.ContextWithRequestID(ctx, requestID)
	} else if requestID := c.Request.Header.Get("X-Request-ID"); requestID != "" {
		ctx = api.ContextWithRequestID(ctx, requestID)
	}
	return ctx, nil
}

// SetAuthCookies is a helper function to set authentication cookies on a echo request.
// The access token cookie (access_token) is an http only cookie that expires when the
// access token expires. The refresh token cookie is not an http only cookie (it can be
// accessed by client-side scripts) and it expires when the refresh token expires. Both
// cookies require https and will not be set (silently) over http connections.
func SetAuthCookies(c *echo.Context, accessToken, refreshToken, domain string) (err error) {
	// Parse access token to get expiration time
	var accessExpires time.Time
	if accessExpires, err = tokens.ExpiresAt(accessToken); err != nil {
		return err
	}

	// Set the access token cookie: httpOnly is true; cannot be accessed by Javascript
	accessMaxAge := int((time.Until(accessExpires)).Seconds())
	c.SetCookie(AccessTokenCookie, accessToken, accessMaxAge, "/", domain, true, true)

	// Parse refresh token to get expiration time
	var refreshExpires time.Time
	if refreshExpires, err = tokens.ExpiresAt(refreshToken); err != nil {
		return err
	}

	// Set the refresh token cookie: httpOnly is false; can be accessed by Javascript
	refreshMaxAge := int((time.Until(refreshExpires)).Seconds())
	c.SetCookie(RefreshTokenCookie, refreshToken, refreshMaxAge, "/", domain, true, false)
	return nil
}

// ClearAuthCookies is a helper function to clear authentication cookies on a echo
// request to effectively logger out a user.
func ClearAuthCookies(c *echo.Context, domain string) {
	c.SetCookie(AccessTokenCookie, "", -1, "/", domain, true, true)
	c.SetCookie(RefreshTokenCookie, "", -1, "/", domain, true, false)
}
