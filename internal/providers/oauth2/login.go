package oauth2

import (
	"net/http"

	"golang.org/x/oauth2"

	"github.com/datumforge/datum/internal/keygen"
	"github.com/datumforge/datum/internal/sessions"
)

// StateHandler checks for a state cookie, if found, adds to context; if missing, a
// random generated value is added to the context and to a (short-lived) state cookie
// issued to the requester - this implements OAuth 2 RFC 6749 10.12 CSRF Protection
func StateHandler(config sessions.CookieConfig, success http.Handler) http.Handler {
	funk := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		cookie, err := req.Cookie(config.Name)

		if err == nil {
			ctx = WithState(ctx, cookie.Value)
		} else {
			val := keygen.GenerateRandomString(32) // nolint: gomnd
			http.SetCookie(w, sessions.NewCookie(config.Name, val, &config))
			ctx = WithState(ctx, val)
		}

		success.ServeHTTP(w, req.WithContext(ctx))
	}

	return http.HandlerFunc(funk)
}

// LoginHandler reads the state value from the context and redirects requests to the AuthURL with that state value
func LoginHandler(config *oauth2.Config, failure http.Handler) http.Handler {
	if failure == nil {
		failure = DefaultFailureHandler
	}

	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		state, err := StateFromContext(ctx)

		if err != nil {
			ctx = WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		authURL := config.AuthCodeURL(state)
		http.Redirect(w, req, authURL, http.StatusFound)
	}

	return http.HandlerFunc(fn)
}

// CallbackHandler parses the auth code + state and compares it to the state value from the context
func CallbackHandler(config *oauth2.Config, success, failure http.Handler) http.Handler {
	if failure == nil {
		failure = DefaultFailureHandler
	}

	funk := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		authCode, state, err := parseCallback(req)

		if err != nil {
			ctx = WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		ownerState, err := StateFromContext(ctx)

		if err != nil {
			ctx = WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		if state != ownerState || state == "" {
			ctx = WithError(ctx, ErrInvalidState)
			failure.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		token, err := config.Exchange(ctx, authCode)

		if err != nil {
			ctx = WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		ctx = WithToken(ctx, token)
		success.ServeHTTP(w, req.WithContext(ctx))
	}

	return http.HandlerFunc(funk)
}

// parseCallback parses code and state parameters from the http.Request and returns them
func parseCallback(req *http.Request) (authCode, state string, err error) {
	err = req.ParseForm()

	if err != nil {
		return "", "", err
	}

	authCode = req.Form.Get("code")

	state = req.Form.Get("state")

	if authCode == "" || state == "" {
		return "", "", ErrMissingCodeOrState
	}

	return authCode, state, nil
}
