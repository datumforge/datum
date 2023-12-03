package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/datumforge/datum/cmd/cli/vars"
	"golang.org/x/oauth2"
)

// Exchanger holds the config for the token exchanger
type Exchanger struct {
	url       string
	grantType string
	tokenType string
}

// ExchangeResponse is the structure of the response from a token exchange
type ExchangeResponse struct {
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	IssuedTokenType string `json:"issued_token_type"`
	TokenType       string `json:"token_type"`
}

// NewExchanger returns a token exchanger
func NewExchanger(u, g, t string) *Exchanger {
	return &Exchanger{
		url:       u,
		grantType: g,
		tokenType: t,
	}
}

// Exchange attempts to exchange the given token for another via the token exchange
func (e *Exchanger) Exchange(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	form := url.Values{}
	form.Add("grant_type", e.grantType)
	form.Add("code", token.AccessToken)
	form.Add("redirect_uri", vars.RedirectURL)
	form.Add("client_id", vars.OauthClientID)
	form.Add("refresh_token", token.RefreshToken)
	form.Add("response_type", e.tokenType)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, e.url, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, _ := io.ReadAll(res.Body)
	fmt.Printf("res: %v\n", string(b))

	exResp := ExchangeResponse{}
	if err := json.NewDecoder(res.Body).Decode(&exResp); err != nil {
		return nil, err
	}

	fmt.Printf("Response: %v\n", exResp)

	return &oauth2.Token{
		AccessToken: exResp.AccessToken,
		TokenType:   exResp.TokenType,
		Expiry:      time.Now().Add(time.Duration(exResp.ExpiresIn)),
	}, nil
}
