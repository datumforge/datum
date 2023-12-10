package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/tokens"
)

// ContextUserClaims is the context key for the user claims
var ContextUserClaims = &ContextKey{"user_claims"}

// ContextAccessToken is the context key for the access token
var ContextAccessToken = &ContextKey{"access_token"}

// ContextRequestID is the context key for the request ID
var ContextRequestID = &ContextKey{"request_id"}

// ContextKey is the key name for the additional context
type ContextKey struct {
	name string
}

func Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := NewAuthOptions()

			validator, err := conf.Validator()
			if err != nil {
				return err
			}

			// Create a reauthenticator function to handle refresh tokens if they are provided.
			reauthenticate := Reauthenticate(conf, validator)

			// Get access token from the request, if not available then attempt to refresh
			// using the refresh token cookie.
			accessToken, err := GetAccessToken(c)
			if err != nil {
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
			claims, err := validator.Verify(accessToken)
			if err != nil {
				ErrorResponse(ErrAuthRequired)
				return err
			}

			// Add claims to context for use in downstream processing and continue handlers
			c.Set(ContextUserClaims.name, claims)

			return next(c)
		}
	}
}

// Reauthenticate is a middleware helper that can use refresh tokens in the echo context
// to obtain a new access token. If it is unable to obtain a new valid access token,
// then an error is returned and processing should stop.
func Reauthenticate(conf AuthOptions, validator tokens.Validator) func(c echo.Context) (string, error) {
	// If no reauthenticator is available on the configuration, always return an error.
	if conf.reauth == nil {
		return func(c echo.Context) (string, error) {
			return "", ErrRefreshDisabled
		}
	}

	// If the reauthenticator is available, return a function that utilizes it.
	return func(c echo.Context) (string, error) {
		// Get the refresh token from the cookies or the headers of the request.
		refreshToken, err := GetRefreshToken(c)
		if err != nil {
			return "", err
		}

		// Check to ensure the refresh token is still valid.
		if _, err = validator.Verify(refreshToken); err != nil {
			return "", err
		}

		// Reauthenticate using the refresh token.
		req := &RefreshRequest{RefreshToken: refreshToken}

		reply, err := conf.reauth.Refresh(c.Request().Context(), req)
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
func GetAccessToken(c echo.Context) (string, error) {
	// Attempt to get the access token from the header.
	if h := c.Request().Header.Get(Authorization); h != "" {
		match := bearer.FindStringSubmatch(h)
		if len(match) == 2 { //nolint:gomnd
			return match[1], nil
		}

		return "", ErrParseBearer
	}

	// Attempt to get the access token from cookies.
	if cookie, err := c.Cookie(AccessTokenCookie); err == nil {
		// If the error is nil, that means we were able to retrieve the access token cookie
		if CookieExpired(cookie) {
			return "", ErrNoAuthorization
		}

		return cookie.Value, nil
	}

	return "", ErrNoAuthorization
}

// GetRefreshToken retrieves the refresh token from the cookies in the request. If the
// cookie is not present or expired then an error is returned.
func GetRefreshToken(c echo.Context) (string, error) {
	cookie, err := c.Cookie(RefreshTokenCookie)
	if err != nil {
		return "", ErrNoRefreshToken
	}

	// ensure cookie is not expired
	if CookieExpired(cookie) {
		return "", ErrNoRefreshToken
	}

	return cookie.Value, nil
}

// GetClaims fetches and parses datum claims from the echo context. Returns an
// error if no claims exist on the context
func GetClaims(c echo.Context) (*tokens.Claims, error) {
	claims, ok := c.Get(ContextUserClaims.name).(*tokens.Claims)
	if !ok {
		return nil, ErrNoClaims
	}

	return claims, nil
}

// AuthContextFromRequest creates a context from the echo request context, copying fields
// that may be required for forwarded requests. This method should be called by
// handlers which need to forward requests to other services and need to preserve data
// from the original request such as the user's credentials.
func AuthContextFromRequest(c echo.Context) (*context.Context, error) {
	req := c.Request()
	if req == nil {
		return nil, ErrNoRequest
	}

	// Add access token to context (from either header or cookie using Authenticate middleware)
	ctx := req.Context()
	if token := c.Get(ContextAccessToken.name); token != "" {
		ctx = context.WithValue(ctx, ContextAccessToken, token)
	}

	// Add request id to context
	if requestID := c.Get(ContextRequestID.name); requestID != "" {
		ctx = context.WithValue(ctx, ContextRequestID, requestID)
	} else if requestID := c.Request().Header.Get("X-Request-ID"); requestID != "" {
		ctx = context.WithValue(ctx, ContextRequestID, requestID)
	}

	return &ctx, nil
}

// SetAuthCookies is a helper function to set authentication cookies on a echo request.
// The access token cookie (access_token) is an http only cookie that expires when the
// access token expires. The refresh token cookie is not an http only cookie (it can be
// accessed by client-side scripts) and it expires when the refresh token expires. Both
// cookies require https and will not be set (silently) over http connections.
func SetAuthCookies(c echo.Context, accessToken, refreshToken, domain string) error {
	// Parse access token to get expiration time
	accessExpires, err := tokens.ExpiresAt(accessToken)
	if err != nil {
		return err
	}

	// Set the access token cookie: httpOnly is true; cannot be accessed by Javascript
	accessMaxAge := int((time.Until(accessExpires)).Seconds())
	cookie := &http.Cookie{
		Name:     AccessTokenCookie,
		Value:    accessToken,
		MaxAge:   accessMaxAge,
		Domain:   domain,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)

	// Parse refresh token to get expiration time
	refreshExpires, err := tokens.ExpiresAt(refreshToken)
	if err != nil {
		return err
	}

	// Set the refresh token cookie: httpOnly is false; can be accessed by Javascript
	refreshMaxAge := int((time.Until(refreshExpires)).Seconds())
	cookie = &http.Cookie{
		Name:     RefreshTokenCookie,
		Value:    refreshToken,
		MaxAge:   refreshMaxAge,
		Domain:   domain,
		Path:     "/",
		HttpOnly: false,
		Secure:   true,
	}

	c.SetCookie(cookie)

	return err
}

// ClearAuthCookies is a helper function to clear authentication cookies on a echo
// request to effectively logger out a user.
func ClearAuthCookies(c echo.Context, domain string) {
	cookie := &http.Cookie{
		Name:     AccessTokenCookie,
		Value:    "",
		MaxAge:   -1,
		Domain:   domain,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)

	cookie = &http.Cookie{
		Name:     RefreshTokenCookie,
		Value:    "",
		MaxAge:   -1,
		Domain:   domain,
		Path:     "/",
		HttpOnly: false,
		Secure:   true,
	}

	c.SetCookie(cookie)
}

// CookieExpired checks to see if a cookie is expired
func CookieExpired(cookie *http.Cookie) bool {
	// ensure cookie is not expired
	if !cookie.Expires.IsZero() && cookie.Expires.Before(time.Now()) {
		return true
	}

	// negative max age means to expire immediately
	if cookie.MaxAge < 0 {
		return true
	}

	return false
}
