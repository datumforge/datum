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

	"github.com/datumforge/datum/internal/httpserve/route"
	"github.com/datumforge/datum/pkg/models"
)

// Switch to a different Datum organization
func Switch(c *Client, ctx context.Context, r models.SwitchOrganizationRequest, accessToken string) (*oauth2.Token, error) {
	method := http.MethodPost
	endpoint := "switch"

	u := fmt.Sprintf("%s%s/%s", c.Client.BaseURL, route.V1Version, endpoint)

	queryURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", accessToken))
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

	out := models.SwitchOrganizationReply{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, newAuthenticationError(resp.StatusCode, fmt.Sprintf("%v", out.Reply))
	}

	return getTokensFromCookiesFromResponse(resp), err
}
