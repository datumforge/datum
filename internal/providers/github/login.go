package github

import (
	"net/http"
	"net/url"

	"github.com/google/go-github/v58/github"
	"golang.org/x/oauth2"

	oauth2Login "github.com/datumforge/datum/internal/providers/oauth2"
	"github.com/datumforge/datum/internal/sessions"
)

// StateHandler checks for a state cookie, if found, adds to context; if missing, a
// random generated value is added to the context and to a (short-lived) state cookie
// issued to the requester - this implements OAuth 2 RFC 6749 10.12 CSRF Protection
func StateHandler(config sessions.CookieConfig, success http.Handler) http.Handler {
	return oauth2Login.StateHandler(config, success)
}

// LoginHandler handles Github login requests by reading the state value from
// the ctx and redirecting requests to the AuthURL with that state value
func LoginHandler(config *oauth2.Config, failure http.Handler) http.Handler {
	return oauth2Login.LoginHandler(config, failure)
}

// CallbackHandler adds the GitHub access token and User to the ctx
func CallbackHandler(config *oauth2.Config, success, failure http.Handler) http.Handler {
	success = githubHandler(config, false, success, failure)

	return oauth2Login.CallbackHandler(config, success, failure)
}

// EnterpriseCallbackHandler handles GitHub Enterprise redirection URI requests
// and adds the GitHub access token and User to the ctx
func EnterpriseCallbackHandler(config *oauth2.Config, success, failure http.Handler) http.Handler {
	success = githubHandler(config, true, success, failure)

	return oauth2Login.CallbackHandler(config, success, failure)
}

// githubHandler gets the OAuth2 Token from the ctx to fetch the corresponding GitHub
// User and add them to the context
func githubHandler(config *oauth2.Config, isEnterprise bool, success, failure http.Handler) http.Handler {
	if failure == nil {
		failure = DefaultFailureHandler
	}

	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		token, err := oauth2Login.TokenFromContext(ctx)

		if err != nil {
			ctx = WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		httpClient := config.Client(ctx, token)

		var githubClient *github.Client

		if isEnterprise {
			githubClient, err = enterpriseGithubClientFromAuthURL(config.Endpoint.AuthURL, httpClient)

			if err != nil {
				ctx = WithError(ctx, ErrCreatingGithubClient)
				failure.ServeHTTP(w, req.WithContext(ctx))

				return
			}
		} else {
			githubClient = github.NewClient(httpClient)
		}

		user, resp, err := githubClient.Users.Get(ctx, "")
		err = validateResponse(user, resp, err)

		if err != nil {
			ctx = WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		ctx = WithUser(ctx, user)
		success.ServeHTTP(w, req.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// validateResponse returns an error if we get something unexpected
func validateResponse(user *github.User, resp *github.Response, err error) error {
	if err != nil || resp.StatusCode != http.StatusOK {
		return ErrUnableToGetGithubUser
	}

	if user == nil || user.ID == nil {
		return ErrUnableToGetGithubUser
	}

	return nil
}

// enterpriseGithubClientFromAuthURL returns a GitHub client that targets a GHE instance
func enterpriseGithubClientFromAuthURL(authURL string, httpClient *http.Client) (*github.Client, error) {
	client := github.NewClient(httpClient)

	baseURL, err := url.Parse(authURL)
	if err != nil {
		return nil, ErrFailedConstructingEndpointURL
	}

	baseURL.Path = "/api/v3/"
	client.BaseURL = baseURL
	client.UploadURL = baseURL

	return client, nil
}
