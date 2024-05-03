package auth

import (
	"context"
	"errors"
	"time"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/apitoken"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/tokens"
)

// SessionSkipperFunc is the function that determines if the session check should be skipped
// due to the request being a PAT or API Token auth request
var SessionSkipperFunc = func(c echo.Context) bool {
	return auth.GetAuthTypeFromEchoContext(c) != auth.JWTAuthentication
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
				// if its not a JWT, check to see if its a PAT or API Token
				if conf.DBClient == nil {
					return rout.HTTPErrorResponse(rout.ErrInvalidCredentials)
				}

				claims, authType, err = checkToken(c.Request().Context(), conf, accessToken)
				if err != nil {
					return rout.HTTPErrorResponse(rout.ErrInvalidCredentials)
				}
			}

			// Add claims to context for use in downstream processing and continue handlers
			au, err := createAuthenticatedUser(c.Request().Context(), conf.DBClient, claims, authType)
			if err != nil {
				return rout.HTTPErrorResponse(rout.ErrInvalidCredentials)
			}

			auth.SetAuthenticatedUserContext(c, au)

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

func createAuthenticatedUser(ctx context.Context, dbClient *generated.Client, claims *tokens.Claims, authType auth.AuthenticationType) (*auth.AuthenticatedUser, error) {
	// get the user ID from the claims
	mappingID := claims.UserID
	mappingOrgID := claims.OrgID

	user, err := dbClient.User.Query().Where(user.MappingID(mappingID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	// all the query to get the organization, need to bypass the authz filter to get the org
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	org, err := dbClient.Organization.Query().Where(organization.MappingID(mappingOrgID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return &auth.AuthenticatedUser{
		SubjectID:          user.ID,
		OrganizationID:     org.ID,
		AuthenticationType: authType,
	}, nil
}

// checkToken checks the bearer authorization token against the database to see if the provided
// token is an active personal access token. If the token is valid, the claims are returned
func checkToken(ctx context.Context, conf AuthOptions, token string) (*tokens.Claims, auth.AuthenticationType, error) {
	// allow check to bypass privacy rules
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	// check if the token is a personal access token
	claims, err := isValidPersonalAccessToken(ctx, conf.DBClient, token)
	if err == nil {
		return claims, auth.PATAuthentication, nil
	}

	// check if the token is an API token
	claims, err = isValidAPIToken(ctx, conf.DBClient, token)
	if err == nil {
		return claims, auth.APITokenAuthentication, nil
	}

	return nil, "", err
}

func isValidPersonalAccessToken(ctx context.Context, dbClient *generated.Client, token string) (*tokens.Claims, error) {
	pat, err := dbClient.PersonalAccessToken.Query().Where(personalaccesstoken.Token(token)).Only(ctx)
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

func isValidAPIToken(ctx context.Context, dbClient *generated.Client, token string) (*tokens.Claims, error) {
	t, err := dbClient.APIToken.Query().Where(apitoken.Token(token)).Only(ctx)
	if err != nil {
		return nil, err
	}

	if !t.ExpiresAt.IsZero() && t.ExpiresAt.Before(time.Now()) {
		return nil, rout.ErrExpiredCredentials
	}

	claims := &tokens.Claims{
		OrgID: t.OwnerID,
	}

	return claims, nil
}
