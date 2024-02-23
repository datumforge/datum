package auth

import (
	"context"
	"errors"
	"time"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/rout"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/pkg/auth"
)

// SessionSkipperFunc is the function that determines if the session check should be skipped
// due to the request being a PAT auth request
var SessionSkipperFunc = func(c echo.Context) bool {
	return c.Get(auth.GetContextName(auth.ContextAuthType)) == auth.PATAuthentication
}

// Authenticate is a middleware function that is used to authenticate requests - it is not applied to all routes so be cognizant of that
func Authenticate(conf AuthOptions) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// if skipper function returns true, skip this middleware
			if conf.Skipper(c) {
				return next(c)
			}

			// execute any before functions
			if conf.BeforeFunc != nil {
				conf.BeforeFunc(c)
			}

			validator, err := conf.Validator()
			if err != nil {
				return err
			}

			// Create a reauthenticator function to handle refresh tokens if they are provided.
			reauthenticate := Reauthenticate(conf, validator)

			// Get access token from the request, if not available then attempt to refresh
			// using the refresh token cookie.
			accessToken, err := auth.GetAccessToken(c)
			if err != nil {
				switch {
				case errors.Is(err, ErrNoAuthorization):
					if accessToken, err = reauthenticate(c); err != nil {
						return rout.HTTPErrorResponse(err)
					}
				default:
					return rout.HTTPErrorResponse(err)
				}
			}

			// Verify the access token is authorized for use with datum and extract claims.
			authType := auth.JWTAuthentication

			claims, err := validator.Verify(accessToken)
			if err != nil {
				// if its not a JWT, check to see if its a PAT
				if conf.DBClient == nil {
					return rout.HTTPErrorResponse(err)
				}

				claims, err = checkToken(c.Request().Context(), conf, accessToken)
				if err != nil {
					return rout.HTTPErrorResponse(err)
				}

				authType = auth.PATAuthentication
			}

			// Add claims to context for use in downstream processing and continue handlers
			c.Set(auth.GetContextName(auth.ContextUserClaims), claims)

			// Set auth type in context
			c.Set(auth.GetContextName(auth.ContextAuthType), authType)

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
		refreshToken, err := auth.GetRefreshToken(c)
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
		auth.SetAuthCookies(c.Response().Writer, reply.AccessToken, reply.RefreshToken)

		return reply.AccessToken, nil
	}
}

// checkToken checks the bearer authorization token against the database to see if the provided
// token is an active personal access token. If the token is valid, the claims are returned
func checkToken(ctx context.Context, conf AuthOptions, token string) (*tokens.Claims, error) {
	// allow check to bypass privacy rules
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	pat, err := conf.DBClient.PersonalAccessToken.Query().Where(personalaccesstoken.Token(token)).Only(ctx)
	if err != nil {
		return nil, err
	}

	if pat.ExpiresAt.Before(time.Now()) {
		return nil, rout.ErrExpiredCredentials
	}

	claims := &tokens.Claims{
		UserID: pat.OwnerID,
	}

	return claims, nil
}
