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

type GithubConfig struct {
	ClientID     string
	ClientSecret string
	Orgs         []string
}

type GoogleConfig struct {
	ClientID     string
	ClientSecret string
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

func (h *Handler) IsAuthenticated(req *http.Request) bool {
	if _, err := h.SM.Get(req, sessions.DefaultCookieName); err == nil {
		return true
	}

	return false
}

// GetGoogleLoginHandlers returns the google login and callback handlers
func (h *Handler) GetGoogleLoginHandlers() (http.Handler, http.Handler) {
	ClientID := ""
	ClientSecret := ""

	var clientEndpoint string

	oauth2Config := &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  fmt.Sprintf("%s/auth/google/callback", clientEndpoint),
		Endpoint:     googleOAuth2.Endpoint,
		Scopes:       []string{"email"},
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
		setSessionMap["userID"] = googleUser.Id
		setSessionMap["username"] = googleUser.Name
		setSessionMap["useremail"] = googleUser.Email
		setSessionMap["usertype"] = "google"
		session.Set(sessionID, setSessionMap)

		if err := session.Save(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, req, "/profile", http.StatusFound)
	}

	return http.HandlerFunc(fn)
}

// GetGitHubLoginHandlers returns the github login and callback handlers
func (h *Handler) GetGitHubLoginHandlers() (http.Handler, http.Handler) {
	ClientID := ""
	ClientSecret := ""

	var clientEndpoint string

	oauth2Config := &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  fmt.Sprintf("%s/github/callback", clientEndpoint),
		Endpoint:     githubOAuth2.Endpoint,
		Scopes:       []string{"email"},
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
		setSessionMap["userID"] = fmt.Sprintf("%v", githubUser.ID)
		setSessionMap["username"] = *githubUser.Login
		setSessionMap["useremail"] = *githubUser.Email
		setSessionMap["usertype"] = "github"
		session.Set(sessionID, setSessionMap)

		if err := session.Save(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, req, "/profile", http.StatusFound)
	}

	return http.HandlerFunc(fn)
}
