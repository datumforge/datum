package datumclient

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

var (
	// callbackTimeout to timeout handler when validating an oauth login
	callbackTimeout = time.Second * 60
	// readTimeout for the handler to receive the oauth request
	readTimeout = time.Second * 10
)

type getTokenResult struct {
	token   *oauth2.Token
	session string
	err     error
}

// OauthLogin opens a browser window to authenticate with the provided oauth2 URL
func OauthLogin(u string, isDev bool) (*oauth2.Token, string, error) {
	serverAddr := "localhost:18000"
	callback := "/oauth/callback"

	redirectURL := fmt.Sprintf("http://%s%s", serverAddr, callback)

	urlWithRedirect := fmt.Sprintf("%s?redirect_uri=%s", u, url.QueryEscape(redirectURL))

	tokenResults := make(chan getTokenResult, 1)

	ctx, cancel := context.WithTimeout(context.Background(), callbackTimeout)
	defer cancel()

	serveMux := http.NewServeMux()
	serveMux.HandleFunc(callback, func(w http.ResponseWriter, r *http.Request) {
		token, session := getTokensFromCookieRequest(r)

		_, _ = w.Write([]byte("Success. You can now close this window.")) //nolint

		select {
		case tokenResults <- getTokenResult{token: token, session: session, err: nil}:
		case <-ctx.Done():
		}
	})

	server := &http.Server{
		Addr:              serverAddr,
		Handler:           serveMux,
		ReadHeaderTimeout: readTimeout,
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

	if err := open.Start(urlWithRedirect); err != nil {
		fmt.Printf("Failed to open browser automatically, please visit %s to complete auth\n\n", u)
	}

	var result getTokenResult

	select {
	case result = <-tokenResults:
	case <-ctx.Done():
		result.err = fmt.Errorf("timed-out waiting for user to complete oauth2 workflow - %w", ctx.Err())
	}

	if result.err != nil {
		return nil, "", ErrFailedToGetOauthToken
	}

	return result.token, result.session, nil
}
