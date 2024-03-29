package datumclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	echo "github.com/datumforge/echox"
	"golang.org/x/oauth2"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/route"
)

// Refresh the access + refresh token pair to the Datum API
func Refresh(c *Client, ctx context.Context, r handlers.RefreshRequest) (*oauth2.Token, error) {
	method := http.MethodPost
	endpoint := "refresh"

	u := fmt.Sprintf("%s%s/%s", c.Client.BaseURL, route.V1Version, endpoint)

	queryURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// Set Headers
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(b))

	resp, err := c.Client.Client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	out := handlers.RefreshReply{}

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, newAuthenticationError(resp.StatusCode, out.Message)
	}

	return getTokensFromCookiesFromResponse(resp), nil
}
