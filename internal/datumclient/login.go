package datumclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
)

// Login creates a login request to the Datum API
func Login(c *Client, ctx context.Context, user handlers.User) (*oauth2.Token, error) {
	method := http.MethodPost
	endpoint := "login"

	u := fmt.Sprintf("%s%s", c.Client.BaseURL, endpoint)

	queryURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(b))

	resp, err := c.Client.Client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	out := handlers.Response{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, newAuthenticationError(resp.StatusCode, out.Message)
	}

	return getTokensFromCookies(resp), nil
}

// getTokensFromCookies parses the HTTP Response for cookies and returns the access and refresh tokens
func getTokensFromCookies(resp *http.Response) *oauth2.Token {
	token := &oauth2.Token{}

	// parse cookies
	cookies := resp.Cookies()

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
