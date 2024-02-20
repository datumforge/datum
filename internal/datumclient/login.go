package datumclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/route"
	"github.com/datumforge/datum/internal/sessions"
)

// Login creates a login request to the Datum API
func Login(c *Client, ctx context.Context, login handlers.LoginRequest) (*oauth2.Token, error) {
	method := http.MethodPost
	endpoint := "login"

	u := fmt.Sprintf("%s%s/%s", c.Client.BaseURL, route.V1Version, endpoint)

	queryURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(login)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(b))

	resp, err := c.Client.Client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	out := handlers.LoginReply{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, newAuthenticationError(resp.StatusCode, out.Message)
	}

	return getTokensFromCookiesFromResponse(resp), nil
}

// getTokensFromCookies returns the access and refresh tokens from the http cookies
func getTokensFromCookies(cookies []*http.Cookie) (token *oauth2.Token) {
	token = &oauth2.Token{}

	for _, c := range cookies {
		if c.Name == auth.AccessTokenCookie {
			token.AccessToken = c.Value
		}

		if c.Name == auth.RefreshTokenCookie {
			token.RefreshToken = c.Value
		}
	}

	return token
}

// getTokensFromCookiesFromResponse parses the HTTP Response for cookies and returns the access and refresh tokens
func getTokensFromCookiesFromResponse(resp *http.Response) (token *oauth2.Token) {
	// parse cookies
	cookies := resp.Cookies()

	return getTokensFromCookies(cookies)
}

// getTokensFromCookies parses the HTTP Request for cookies and returns the session and access and refresh tokens
func getTokensFromCookieRequest(r *http.Request, isDev bool) (token *oauth2.Token, session string) {
	// parse cookies
	cookies := r.Cookies()
	cookieName := sessions.DefaultCookieName

	// Use the dev cookie when running on localhost
	if isDev {
		cookieName = sessions.DevCookieName
	}

	for _, c := range cookies {
		if c.Name == cookieName {
			session = c.Value
		}
	}

	return getTokensFromCookies(cookies), session
}

// GetSessionFromCookieJar parses the cookie jar for the session cookie
func GetSessionFromCookieJar(c *Client) (sessionID string, err error) {
	u, err := url.Parse(c.Client.BaseURL)
	if err != nil {
		return "", err
	}

	cookies := c.Client.Client.Jar.Cookies(u)
	cookieName := sessions.DefaultCookieName

	// Use the dev cookie when running on localhost
	if strings.Contains(u.Host, "localhost") {
		cookieName = sessions.DevCookieName
	}

	for _, c := range cookies {
		if c.Name == cookieName {
			return c.Value, nil
		}
	}

	return "", nil
}
