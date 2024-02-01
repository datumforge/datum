package handlers

import (
	"fmt"
	"net/http"

	echo "github.com/datumforge/echox"
	"golang.org/x/oauth2"
	githubOAuth2 "golang.org/x/oauth2/github"
	googleOAuth2 "golang.org/x/oauth2/google"

	"github.com/datumforge/datum/internal/providers/github"
	"github.com/datumforge/datum/internal/providers/google"
	"github.com/datumforge/datum/internal/sessions"
)

var sessionName = "meow"

//func (h *Handler) GetSession(req *http.Request) (*sessions.Session, error) {
//	return h.SM.Get(req, sessionName)
//}
//
//func (h *Handler) IssueSession() *sessions.Session {
//	return h.issueSession()
//}
//
//func (h *Handler) issueSession() *sessions.Session {
//	return h.SM.New(sessionName)
//}
//
//func (h *Handler) DestroySession(w http.ResponseWriter) {
//	h.destroySession(w)
//}
//
//func (h *Handler) destroySession(w http.ResponseWriter) {
//	h.SM.Destroy(w, sessionName)
//}

// requireLogin redirects unauthenticated users to the login route.
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
	if _, err := h.SM.Get(req, sessionName); err == nil {
		return true
	}

	return false
}

var DebugOnlyCookieConfig = sessions.CookieConfig{
	Name:     "temporary-cookie",
	Path:     "/",
	MaxAge:   600, // 10 min
	HTTPOnly: true,
	Secure:   false, // allows cookies to be send over HTTP
	SameSite: http.SameSiteLaxMode,
}

func (h *Handler) GetGoogleLoginHandlers(clientID, clientSecret, callbackUrl string) (http.Handler, http.Handler) {
	oauth2Config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackUrl,
		Endpoint:     googleOAuth2.Endpoint,
		Scopes:       []string{"email"},
	}

	// state param cookies require HTTPS by default; disable for localhost development
	stateConfig := DebugOnlyCookieConfig
	loginHandler := google.StateHandler(stateConfig, google.LoginHandler(oauth2Config, nil))
	callbackHandler := google.StateHandler(stateConfig, google.CallbackHandler(oauth2Config, h.issueGoogleSession(), nil))
	return loginHandler, callbackHandler
}

// issueSession issues a cookie session after successful Facebook login
func (h *Handler) issueGoogleSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		googleUser, err := google.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// implement a success handler to issue some form of session
		session := h.issueSession()
		session.Set("userid", googleUser.Id)
		session.Set("username", googleUser.Name)
		session.Set("useremail", googleUser.Email)
		session.Set("usertype", "google")
		session.Save(w)
		http.Redirect(w, req, "/profile", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

func (h *Handler) GetGitHubLoginHandlers(clientID, clientSecret, callbackUrl string) (http.Handler, http.Handler) {
	// 1. Register Login and Callback handlers
	oauth2Config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackUrl,
		Endpoint:     githubOAuth2.Endpoint,
		Scopes:       []string{"email"},
	}

	// state param cookies require HTTPS by default; disable for localhost development
	stateConfig := DebugOnlyCookieConfig
	loginHandler := github.StateHandler(stateConfig, github.LoginHandler(oauth2Config, nil))
	callbackHandler := github.StateHandler(stateConfig, github.CallbackHandler(oauth2Config, h.issueGitHubSession(), nil))
	return loginHandler, callbackHandler
}

// issueSession issues a cookie session after successful Facebook login
func (h *Handler) issueGitHubSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		githubUser, err := github.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 2. Implement a success handler to issue some form of session
		session := h.issueSession()
		session.Set("userid", fmt.Sprintf("%v", *githubUser.ID))
		session.Set("username", fmt.Sprintf("%v", *githubUser.Login))
		session.Set("useremail", fmt.Sprintf("%v", *githubUser.Login))
		if nil != githubUser.Email {
			session.Set("useremail", fmt.Sprintf("%v", *githubUser.Email))
		}
		session.Set("usertype", "github")
		session.Save(w)
		http.Redirect(w, req, "/profile", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

func (h *Handler) meow(ctx echo.Context) error {

	// set sessions in response
	sc := sessions.NewSessionConfig(h.SM, h.RedisClient, h.Logger)

	userID := "meow"
	if err := sc.SaveAndStoreSession(ctx, userID); err != nil {
		h.Logger.Errorw("unable to save session", "error", err)

		return err
	}

	return nil

}

// logoutHandler destroys the session on POSTs and redirects to home.
func (h *Handler) LogoutHandler(w http.ResponseWriter, req *http.Request) {
	// if req.Method == "POST" {
	h.destroySession(w)
	// }
	http.Redirect(w, req, "/", http.StatusFound)
}
