package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/apitoken"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/pkg/auth"
	api "github.com/datumforge/datum/pkg/models"
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

			var au *auth.AuthenticatedUser

			claims, err := validator.Verify(accessToken)
			if err != nil {
				// if its not a JWT, check to see if its a PAT or API Token
				if conf.DBClient == nil {
					return rout.HTTPErrorResponse(rout.ErrInvalidCredentials)
				}

				au, err = checkToken(c.Request().Context(), conf, accessToken)
				if err != nil {
					return rout.HTTPErrorResponse(rout.ErrInvalidCredentials)
				}
			} else {
				// Add claims to context for use in downstream processing and continue handlers
				au, err = createAuthenticatedUserFromClaims(c.Request().Context(), conf.DBClient, claims, authType)
				if err != nil {
					return rout.HTTPErrorResponse(rout.ErrInvalidCredentials)
				}
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
		req := &api.RefreshRequest{RefreshToken: refreshToken}

		reply, err := conf.reauth.Refresh(c.Request().Context(), req)
		if err != nil {
			return "", err
		}

		// Set the new access and refresh cookies
		auth.SetAuthCookies(c.Response().Writer, reply.AccessToken, reply.RefreshToken, *conf.CookieConfig)

		return reply.AccessToken, nil
	}
}

// createAuthenticatedUserFromClaims creates an authenticated user from the claims provided
func createAuthenticatedUserFromClaims(ctx context.Context, dbClient *generated.Client, claims *tokens.Claims, authType auth.AuthenticationType) (*auth.AuthenticatedUser, error) {
	// get the user ID from the claims
	user, err := dbClient.User.Get(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	// all the query to get the organization, need to bypass the authz filter to get the org
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	org, err := dbClient.Organization.Get(ctx, claims.OrgID)
	if err != nil {
		return nil, err
	}

	return &auth.AuthenticatedUser{
		SubjectID:          user.ID,
		SubjectName:        getSubjectName(user),
		OrganizationID:     org.ID,
		OrganizationName:   org.Name,
		OrganizationIDs:    []string{org.ID},
		AuthenticationType: authType,
	}, nil
}

// checkToken checks the bearer authorization token against the database to see if the provided
// token is an active personal access token or api token.
// If the token is valid, the authenticated user is returned
func checkToken(ctx context.Context, conf AuthOptions, token string) (*auth.AuthenticatedUser, error) {
	// allow check to bypass privacy rules
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	// check if the token is a personal access token
	au, err := isValidPersonalAccessToken(ctx, conf.DBClient, token)
	if err == nil {
		return au, nil
	}

	// check if the token is an API token
	au, err = isValidAPIToken(ctx, conf.DBClient, token)
	if err == nil {
		return au, nil
	}

	return nil, err
}

// isValidPersonalAccessToken checks if the provided token is a valid personal access token and returns the authenticated user
func isValidPersonalAccessToken(ctx context.Context, dbClient *generated.Client, token string) (*auth.AuthenticatedUser, error) {
	pat, err := dbClient.PersonalAccessToken.Query().Where(personalaccesstoken.Token(token)).
		WithOwner().
		WithOrganizations().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if pat.ExpiresAt.Before(time.Now()) {
		return nil, rout.ErrExpiredCredentials
	}

	// gather the authorized organization IDs
	var orgIDs []string

	orgs := pat.Edges.Organizations

	if len(orgs) == 0 {
		// an access token must have at least one organization to be used
		return nil, rout.ErrInvalidCredentials
	}

	for _, org := range orgs {
		orgIDs = append(orgIDs, org.ID)
	}

	return &auth.AuthenticatedUser{
		SubjectID:          pat.OwnerID,
		SubjectName:        getSubjectName(pat.Edges.Owner),
		OrganizationIDs:    orgIDs,
		AuthenticationType: auth.PATAuthentication,
	}, nil
}

func isValidAPIToken(ctx context.Context, dbClient *generated.Client, token string) (*auth.AuthenticatedUser, error) {
	t, err := dbClient.APIToken.Query().Where(apitoken.Token(token)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if !t.ExpiresAt.IsZero() && t.ExpiresAt.Before(time.Now()) {
		return nil, rout.ErrExpiredCredentials
	}

	return &auth.AuthenticatedUser{
		SubjectID:          t.ID,
		SubjectName:        fmt.Sprintf("service: %s", t.Name),
		OrganizationID:     t.OwnerID,
		OrganizationIDs:    []string{t.OwnerID},
		AuthenticationType: auth.APITokenAuthentication,
	}, nil
}

// getSubjectName returns the subject name for the user
func getSubjectName(user *generated.User) string {
	subjectName := user.FirstName + " " + user.LastName
	if subjectName == "" {
		subjectName = user.DisplayName
	}

	return subjectName
}
