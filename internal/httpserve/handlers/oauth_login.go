package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/oauth2"
	githubOAuth2 "golang.org/x/oauth2/github"
	googleOAuth2 "golang.org/x/oauth2/google"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/providers/github"
	"github.com/datumforge/datum/pkg/providers/google"
	oauth "github.com/datumforge/datum/pkg/providers/oauth2"
	"github.com/datumforge/datum/pkg/sessions"
)

// OauthProviderConfig represents the configuration for OAuth providers such as Github and Google
type OauthProviderConfig struct {
	// RedirectURL is the URL that the OAuth2 client will redirect to after authentication with datum
	RedirectURL string `json:"redirect_url" koanf:"redirect_url" default:"http://localhost:3001/api/auth/callback/datum"`
	// Github contains the configuration settings for the Github Oauth Provider
	Github github.ProviderConfig `json:"github" koanf:"github"`
	// Google contains the configuration settings for the Google Oauth Provider
	Google google.ProviderConfig `json:"google" koanf:"google"`
}

const (
	githubProvider = "github"
	googleProvider = "google"
)

func (h *Handler) getGoogleOauth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     h.OauthProvider.Google.ClientID,
		ClientSecret: h.OauthProvider.Google.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s%s", h.OauthProvider.Google.ClientEndpoint, h.OauthProvider.Google.RedirectURL),
		Endpoint:     googleOAuth2.Endpoint,
		Scopes:       h.OauthProvider.Google.Scopes,
	}
}

func (h *Handler) getGithubOauth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     h.OauthProvider.Github.ClientID,
		ClientSecret: h.OauthProvider.Github.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s%s", h.OauthProvider.Github.ClientEndpoint, h.OauthProvider.Github.RedirectURL),
		Endpoint:     githubOAuth2.Endpoint,
		Scopes:       h.OauthProvider.Github.Scopes,
	}
}

// RequireLogin redirects unauthenticated users to the login route
func (h *Handler) RequireLogin(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if !h.IsAuthenticated(req) {
			http.Redirect(w, req, "/", http.StatusFound)
			return
		}

		next.ServeHTTP(w, req)
	}

	return http.HandlerFunc(fn)
}

// IsAuthenticated checks the sessions to a valid session cookie
func (h *Handler) IsAuthenticated(req *http.Request) bool {
	if _, err := h.SessionConfig.SessionManager.Get(req, h.SessionConfig.CookieConfig.Name); err == nil {
		return true
	}

	return false
}

// GetGoogleLoginHandlers returns the google login and callback handlers
func (h *Handler) GetGoogleLoginHandlers() (http.Handler, http.Handler) {
	oauth2Config := h.getGoogleOauth2Config()

	loginHandler := google.StateHandler(*h.SessionConfig.CookieConfig, google.LoginHandler(oauth2Config, nil))
	callbackHandler := google.StateHandler(*h.SessionConfig.CookieConfig, google.CallbackHandler(oauth2Config, h.issueGoogleSession(), nil))

	return loginHandler, callbackHandler
}

// issueGoogleSession issues a cookie session after successful Facebook login
func (h *Handler) issueGoogleSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		redirectURI, err := h.getRedirectURI(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		googleUser, err := google.UserFromContext(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctxWithToken := token.NewContextWithOauthTooToken(ctx, googleUser.Email)

		// check if users exists and create if not
		user, err := h.CheckAndCreateUser(ctxWithToken, googleUser.Name, googleUser.Email, enums.Google)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create session with external data
		setSessionMap := map[string]string{}
		setSessionMap[sessions.ExternalUserIDKey] = googleUser.Id
		setSessionMap[sessions.UsernameKey] = googleUser.Name
		setSessionMap[sessions.EmailKey] = googleUser.Email
		setSessionMap[sessions.UserTypeKey] = googleProvider
		setSessionMap[sessions.UserIDKey] = user.ID

		if err := h.SessionConfig.SaveAndStoreSession(ctxWithToken, w, setSessionMap, user.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set context for remaining request based on logged in user
		userCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(user.ID, true))

		if err := h.addDefaultOrgToUserQuery(userCtx, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// this might get moved based on the UI auth flow
		// but works here for the cli login
		claims := createClaims(user)

		access, refresh, err := h.TM.CreateTokenPair(claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set cookies on request with the access and refresh token
		auth.SetAuthCookies(w, access, refresh)

		// remove cookie
		sessions.RemoveCookie(w, "redirect_to", *h.SessionConfig.CookieConfig)

		http.Redirect(w, req, redirectURI, http.StatusFound)
	}

	return http.HandlerFunc(fn)
}

// GetGitHubLoginHandlers returns the github login and callback handlers
func (h *Handler) GetGitHubLoginHandlers() (http.Handler, http.Handler) {
	oauth2Config := h.getGithubOauth2Config()

	loginHandler := github.StateHandler(*h.SessionConfig.CookieConfig, github.LoginHandler(oauth2Config, nil))
	callbackHandler := github.StateHandler(*h.SessionConfig.CookieConfig, github.CallbackHandler(oauth2Config, h.issueGitHubSession(), nil))

	return loginHandler, callbackHandler
}

// issueGitHubSession issues a cookie session after successful Facebook login
func (h *Handler) issueGitHubSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		redirectURI, err := h.getRedirectURI(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		githubUser, err := github.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// we need the email to keep going, if its not there error the request
		if githubUser.Email == nil {
			http.Error(w, ErrNoEmailFound.Error(), http.StatusBadRequest)
			return
		}

		ctxWithToken := token.NewContextWithOauthTooToken(ctx, *githubUser.Email)

		// check if users exists and create if not, updates last seen of existing user
		user, err := h.CheckAndCreateUser(ctxWithToken, *githubUser.Name, *githubUser.Email, enums.GitHub)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set context for remaining request based on logged in user
		userCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(user.ID, true))

		if err := h.addDefaultOrgToUserQuery(userCtx, user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// this might get moved based on the UI auth flow
		// but works here for the cli login
		claims := createClaims(user)

		access, refresh, err := h.TM.CreateTokenPair(claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		auth.SetAuthCookies(w, access, refresh)

		setSessionMap := map[string]string{}
		setSessionMap[sessions.ExternalUserIDKey] = fmt.Sprintf("%v", githubUser.ID)
		setSessionMap[sessions.UsernameKey] = *githubUser.Login
		setSessionMap[sessions.UserTypeKey] = githubProvider
		setSessionMap[sessions.EmailKey] = *githubUser.Email
		setSessionMap[sessions.UserIDKey] = user.ID

		if err := h.SessionConfig.SaveAndStoreSession(ctx, w, setSessionMap, user.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// setup viewer context
		viewerCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(user.ID, true))

		// remove cookie now that its in the context
		sessions.RemoveCookie(w, "redirect_to", *h.SessionConfig.CookieConfig)

		// redirect with context set
		http.Redirect(w, req.WithContext(viewerCtx), redirectURI, http.StatusFound)
	}

	return http.HandlerFunc(fn)
}

// parseName attempts to parse a full name into first and last names of the user
func parseName(name string) (c ent.CreateUserInput) {
	parts := strings.Split(name, " ")
	c.FirstName = parts[0]

	if len(parts) > 1 {
		c.LastName = strings.Join(parts[1:], " ")
	}

	return
}

// CheckAndCreateUser takes a user with an OauthTooToken set in the context and checks if the user is already created
// if the user already exists, update last seen
func (h *Handler) CheckAndCreateUser(ctx context.Context, name, email string, provider enums.AuthProvider) (*ent.User, error) {
	// check if users exists
	entUser, err := h.getUserByEmail(ctx, email, provider)
	if err != nil {
		// if the user is not found, create now
		if ent.IsNotFound(err) {
			isOAuthUser := true
			lastSeen := time.Now()

			// create new user input
			input := parseName(name)
			input.Email = email
			input.Oauth = &isOAuthUser
			input.AuthProvider = &provider
			input.LastSeen = &lastSeen

			entUser, err = h.createUser(ctx, input)
			if err != nil {
				h.Logger.Errorw("error creating new user", "error", err)

				return nil, err
			}

			// return newly created user
			return entUser, nil
		}

		return nil, err
	}

	// update last seen of user
	if err := h.updateUserLastSeen(ctx, entUser.ID); err != nil {
		h.Logger.Errorw("unable to update last seen", "error", err)

		return nil, err
	}

	return entUser, nil
}

// getRedirectURI checks headers for a request type, if not set, will default to the browser redirect url
func (h *Handler) getRedirectURI(req *http.Request) (string, error) {
	redirectURI := oauth.RedirectFromContext(req.Context())

	// use the default if it was not passed in
	if redirectURI == "" {
		redirectURI = h.OauthProvider.RedirectURL
	}

	return redirectURI, nil
}
