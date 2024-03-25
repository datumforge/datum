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

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/route"
)

// Reset a user password
func Reset(c *Client, ctx context.Context, r handlers.ResetPasswordRequest) (*handlers.ResetPasswordReply, error) {
	method := http.MethodPost
	endpoint := "password-reset"

	u := fmt.Sprintf("%s%s/%s?token=%s", c.Client.BaseURL, route.V1Version, endpoint, r.Token)

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

	out := handlers.ResetPasswordReply{}

	if resp.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
			return nil, err
		}

		return nil, newRequestError(resp.StatusCode, out.Message)
	}

	out.Message = "success"

	return &out, err
}
