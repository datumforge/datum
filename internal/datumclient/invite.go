package datumclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/route"
)

// OrgInvite a new user within Datum org
func OrgInvite(c *Client, ctx context.Context, r handlers.Invite) (*handlers.InviteReply, error) {
	method := http.MethodPost
	endpoint := "invite"

	u := fmt.Sprintf("%s%s/%s", c.Client.BaseURL, route.V1Version, endpoint)

	queryURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

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

	out := handlers.InviteReply{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, newRequestError(resp.StatusCode, out.Message)
	}

	return &out, err
}
