package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

var (
	callbackTimeout = time.Second * 60
)

const (
	ServerPort = 3535
)

type authClient struct {
	oauthConfig  *oauth2.Config
	codeVerifier string
	state        string
}

func (c *authClient) handlePKCECallback(w http.ResponseWriter, r *http.Request) (*oauth2.Token, error) {
	if c.state == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, ErrEmptyClientOauthState
	}

	state := r.URL.Query().Get("state")
	if state != c.state {
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.Wrap(ErrMismatchedOauthState,
			fmt.Sprintf("expected '%s' - got '%s'", c.state, state))
	}

	code := r.URL.Query().Get("code")

	token, err := c.oauthConfig.Exchange(
		r.Context(),
		code,
		oauth2.VerifierOption(c.codeVerifier),
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}

	_, _ = w.Write([]byte("Success. You can now close this window.")) //nolint

	return token, nil
}

type getTokenResult struct {
	token *oauth2.Token
	err   error
}

// AuthPKCE starts a server and listens for an oauth2 callback and will
// return the API token to the caller
func AuthPKCE(oauthConfig *oauth2.Config, audience string) (*oauth2.Token, error) {
	const numStateBytes = 18

	state, err := randomBytesToBase64String(numStateBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate oauth state string - %w", err)
	}

	cv := oauth2.GenerateVerifier()

	c := &authClient{
		oauthConfig:  oauthConfig,
		state:        state,
		codeVerifier: cv,
	}

	// this client forces to use PKCE
	// code_challenge_method = S256 is set by S256ChallengeOption
	authCodeURL := c.oauthConfig.AuthCodeURL(
		string(state),
		oauth2.S256ChallengeOption(cv),
		oauth2.SetAuthURLParam("audience", audience),
	)

	tokenResults := make(chan getTokenResult, 1)

	ctx, cancel := context.WithTimeout(context.Background(), callbackTimeout)
	defer cancel()

	http.Handle("/oauth2/callback", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := c.handlePKCECallback(w, r)

		select {
		case tokenResults <- getTokenResult{token: token, err: err}:
		case <-ctx.Done():
		}
	}))

	server := &http.Server{
		Addr: "127.0.0.1:3535",
	}
	defer func() {
		_ = server.Shutdown(ctx)
	}()

	go func() {
		err := server.ListenAndServe()

		select {
		case tokenResults <- getTokenResult{err: err}:
		case <-ctx.Done():
		}
	}()

	if err := open.Start(authCodeURL); err != nil {
		fmt.Printf("Failed to open browser automatically, please visit %s to complete auth\n\n", authCodeURL)
	}

	var result getTokenResult

	select {
	case result = <-tokenResults:
	case <-ctx.Done():
		result.err = fmt.Errorf("timed-out waiting for user to complete oauth2 workflow - %w", ctx.Err())
	}

	if result.err != nil {
		return nil, errors.Wrap(result.err, ErrFailedToGetOauthToken.Error())
	}

	return result.token, nil
}

func randomBytesToBase64String(nBytes int) (string, error) {
	if nBytes%3 != 0 {
		return "", ErrBase64NumBytes
	}

	result := make([]byte, nBytes)

	_, err := rand.Read(result)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(result), nil
}
