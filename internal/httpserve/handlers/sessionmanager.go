package handlers

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	githubOAuth2 "golang.org/x/oauth2/github"
	googleOAuth2 "golang.org/x/oauth2/google"

	"github.com/datumforge/datum/internal/providers/github"
	"github.com/datumforge/datum/internal/providers/google"
	"github.com/datumforge/datum/internal/sessions"
)

// OauthProviderConfig represents the configuration for OAuth providers such as Github and Google
type OauthProviderConfig struct {
	RedirectURL string `yaml:"redirectURL" split_words:"true" default:"http://localhost:3001/dashboard"`
	GithubConfig
	GoogleConfig
}

// GithubConfig represents the configuration settings for a Github Oauth Provider
type GithubConfig struct {
	ClientID       string   `yaml:"clientId" split_words:"true"`
	ClientSecret   string   `yaml:"clientSecret" split_words:"true"`
	ClientEndpoint string   `yaml:"clientEndpoint" split_words:"true" default:"http://localhost:17608"`
	Scopes         []string `yaml:"scopes" split_words:"true" default:"user:email,read:user"`
	RedirectURL    string   `yaml:"redirectURL" split_words:"true" default:"/v1/github/callback"`
	Orgs           []string `yaml:"orgs" split_words:"true"`
}

// GoogleConfig represents the configuration settings for a Google Oauth Provider
type GoogleConfig struct {
	ClientID       string   `yaml:"clientId" split_words:"true"`
	ClientSecret   string   `yaml:"clientSecret" split_words:"true"`
	ClientEndpoint string   `yaml:"clientEndpoint" split_words:"true" default:"http://localhost:17608"`
	RedirectURL    string   `yaml:"redirectURL" split_words:"true" default:"/v1/google/callback"`
	Scopes         []string `yaml:"scopes" split_words:"true" default:"email"`
}

const (
	githubProvider = "github"
	googleProvider = "google"
)

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

func (h *Handler) IsAuthenticated(req *http.Request) bool {
	if _, err := h.SM.Get(req, sessions.DefaultCookieName); err == nil {
		return true
	}

	return false
}

// GetGoogleLoginHandlers returns the google login and callback handlers
func (h *Handler) GetGoogleLoginHandlers() (http.Handler, http.Handler) {
	oauth2Config := &oauth2.Config{
		ClientID:     h.OauthProvider.GoogleConfig.ClientID,
		ClientSecret: h.OauthProvider.GoogleConfig.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s%s", h.OauthProvider.GoogleConfig.ClientEndpoint, h.OauthProvider.GoogleConfig.RedirectURL),
		Endpoint:     googleOAuth2.Endpoint,
		Scopes:       h.OauthProvider.GoogleConfig.Scopes,
	}

	// state param cookies require HTTPS by default; disable for localhost development
	stateConfig := sessions.DebugOnlyCookieConfig
	loginHandler := google.StateHandler(stateConfig, google.LoginHandler(oauth2Config, nil))
	callbackHandler := google.StateHandler(stateConfig, google.CallbackHandler(oauth2Config, h.issueGoogleSession(), nil))

	return loginHandler, callbackHandler
}

// issueGoogleSession issues a cookie session after successful Facebook login
func (h *Handler) issueGoogleSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		googleUser, err := google.UserFromContext(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session := h.SM.New(sessions.DefaultCookieName)
		sessionID := sessions.GenerateSessionID()

		setSessionMap := map[string]string{}
		setSessionMap[sessions.ExternalUserIDKey] = googleUser.Id
		setSessionMap[sessions.UsernameKey] = googleUser.Name
		setSessionMap[sessions.EmailKey] = googleUser.Email
		setSessionMap[sessions.UserTypeKey] = googleProvider
		session.Set(sessionID, setSessionMap)

		if err := session.Save(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, req, h.OauthProvider.RedirectURL, http.StatusFound)
	}

	return http.HandlerFunc(fn)
}

// GetGitHubLoginHandlers returns the github login and callback handlers
func (h *Handler) GetGitHubLoginHandlers() (http.Handler, http.Handler) {
	oauth2Config := &oauth2.Config{
		ClientID:     h.OauthProvider.GithubConfig.ClientID,
		ClientSecret: h.OauthProvider.GithubConfig.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s%s", h.OauthProvider.GithubConfig.ClientEndpoint, h.OauthProvider.GithubConfig.RedirectURL),
		Endpoint:     githubOAuth2.Endpoint,
		Scopes:       h.OauthProvider.GithubConfig.Scopes,
	}

	// state param cookies require HTTPS by default; disable for localhost development
	stateConfig := sessions.DebugOnlyCookieConfig
	loginHandler := github.StateHandler(stateConfig, github.LoginHandler(oauth2Config, nil))
	callbackHandler := github.StateHandler(stateConfig, github.CallbackHandler(oauth2Config, h.issueGitHubSession(), nil))

	return loginHandler, callbackHandler
}

// issueGitHubSession issues a cookie session after successful Facebook login
func (h *Handler) issueGitHubSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		githubUser, err := github.UserFromContext(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session := h.SM.New(sessions.DefaultCookieName)
		sessionID := sessions.GenerateSessionID()

		setSessionMap := map[string]string{}
		setSessionMap[sessions.ExternalUserIDKey] = fmt.Sprintf("%v", githubUser.ID)
		setSessionMap[sessions.UsernameKey] = *githubUser.Login
		setSessionMap[sessions.UserTypeKey] = githubProvider
		setSessionMap[sessions.EmailKey] = *githubUser.Email
		session.Set(sessionID, setSessionMap)

		if err := session.Save(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, req, h.OauthProvider.RedirectURL, http.StatusFound)
	}

	return http.HandlerFunc(fn)
}
