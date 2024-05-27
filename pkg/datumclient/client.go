package datumclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"

	"github.com/datumforge/datum/pkg/httpsling"
	api "github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/rout"
)

type DatumClient interface {
	Register(context.Context, *api.RegisterRequest) (*api.RegisterReply, error)
	Login(context.Context, *api.LoginRequest) (*api.LoginReply, error)
	Refresh(context.Context, *api.RefreshRequest) (*api.RefreshReply, error)
	Switch(context.Context, *api.SwitchOrganizationRequest) (*api.SwitchOrganizationReply, error)
	VerifyEmail(context.Context, *api.VerifyRequest) (*api.VerifyReply, error)
	ResendEmail(context.Context, *api.ResendRequest) (*api.ResendReply, error)
	ForgotPassword(context.Context, *api.ForgotPasswordRequest) (*api.ForgotPasswordReply, error)
	ResetPassword(context.Context, *api.ResetPasswordRequest) (*api.ResetPasswordReply, error)
	Invite(context.Context, *api.InviteRequest) (*api.InviteReply, error)
	WaitForReady(context.Context) error
}

// New creates a new API v1 client that implements the Datum Client interface
func New(endpoint string, opts ...ClientOption) (_ DatumClient, err error) {
	c := &APIv1{}

	if c.client == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return nil, err
		}

		c.client = httpsling.Create(&httpsling.Config{
			BaseURL:    endpoint,
			Timeout:    5 * time.Second, // nolint: gomnd
			MaxRetries: 3,               // nolint: gomnd
		})

		c.client.HTTPClient.Jar = jar
	}

	// Apply our options
	for _, opt := range opts {
		if err = opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// APIv1 implements the DatumClient interface and also wraps the httpsling client
type APIv1 struct {
	endpoint *url.URL // the base url for all requests
	config   *httpsling.Config
	creds    Credentials // default credentials used to authorize requests
	client   *httpsling.Client
	// graphClient DatumGraphClient
}

// Ensure the APIv1 implements the DatumClient interface
var _ DatumClient = &APIv1{}

// Register a new user with the Datum API
func (s *APIv1) Register(ctx context.Context, in *api.RegisterRequest) (out *api.RegisterReply, err error) {
	var req *http.Request

	if req, err = s.NewRequest(ctx, http.MethodPost, "/v1/register", in, nil); err != nil {
		return nil, err
	}

	resp, err := s.Do(req, &out, true)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return out, nil
}

// Login to the Datum API
func (s *APIv1) Login(ctx context.Context, in *api.LoginRequest) (out *api.LoginReply, err error) {
	var req *http.Request

	if req, err = s.NewRequest(ctx, http.MethodPost, "/v1/login", in, nil); err != nil {
		return nil, err
	}

	resp, err := s.Do(req, &out, true)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return out, nil
}

// Refresh a user's access token
func (s *APIv1) Refresh(ctx context.Context, in *api.RefreshRequest) (out *api.RefreshReply, err error) {
	var req *http.Request

	if req, err = s.NewRequest(ctx, http.MethodPost, "/v1/refresh", in, nil); err != nil {
		return nil, err
	}

	resp, err := s.Do(req, &out, true)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return out, nil
}

// Switch the current organization context
func (s *APIv1) Switch(ctx context.Context, in *api.SwitchOrganizationRequest) (out *api.SwitchOrganizationReply, err error) {
	var req *http.Request

	if req, err = s.NewRequest(ctx, http.MethodPost, "/v1/switch", in, nil); err != nil {
		return nil, err
	}

	resp, err := s.Do(req, &out, true)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return out, nil
}

// VerifyEmail verifies the email address of a user
func (s *APIv1) VerifyEmail(ctx context.Context, in *api.VerifyRequest) (out *api.VerifyReply, err error) {
	var req *http.Request

	if req, err = s.NewRequest(ctx, http.MethodPost, "/v1/verify", in, nil); err != nil {
		return nil, err
	}

	resp, err := s.Do(req, &out, true)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return out, nil
}

// ResendEmail resends the verification email to the user
func (s *APIv1) ResendEmail(ctx context.Context, in *api.ResendRequest) (out *api.ResendReply, err error) {
	var req *http.Request

	if req, err = s.NewRequest(ctx, http.MethodPost, "/v1/resend", in, nil); err != nil {
		return nil, err
	}

	resp, err := s.Do(req, &out, true)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return out, nil
}

// ForgotPassword sends a password reset email to the user
func (s *APIv1) ForgotPassword(ctx context.Context, in *api.ForgotPasswordRequest) (out *api.ForgotPasswordReply, err error) {
	var req *http.Request

	if req, err = s.NewRequest(ctx, http.MethodPost, "/v1/forgot-password", in, nil); err != nil {
		return nil, err
	}

	resp, err := s.Do(req, &out, true)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return out, nil
}

// ResetPassword resets the user's password
func (s *APIv1) ResetPassword(ctx context.Context, in *api.ResetPasswordRequest) (out *api.ResetPasswordReply, err error) {
	var req *http.Request

	if req, err = s.NewRequest(ctx, http.MethodPost, "/v1/password-reset", in, nil); err != nil {
		return nil, err
	}

	resp, err := s.Do(req, &out, true)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return out, nil
}

// Invite a user to an organization
func (s *APIv1) Invite(ctx context.Context, in *api.InviteRequest) (out *api.InviteReply, err error) {
	var req *http.Request

	if req, err = s.NewRequest(ctx, http.MethodPost, "/v1/invite", in, nil); err != nil {
		return nil, err
	}

	resp, err := s.Do(req, &out, true)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return out, nil
}

// WaitForReady waits for the Datum API to be ready
func (s *APIv1) WaitForReady(ctx context.Context) (err error) {
	if _, hasDeadline := ctx.Deadline(); !hasDeadline {
		var cancel context.CancelFunc

		ctx, cancel = context.WithTimeout(ctx, 5*time.Minute) // nolint: gomnd

		defer cancel()
	}

	var req *http.Request

	if req, err = s.NewRequest(ctx, http.MethodGet, "/v1/status", nil, nil); err != nil {
		return err
	}

	checkReady := func() (err error) {
		var rep *http.Response

		if rep, err = s.client.HTTPClient.Do(req); err != nil {
			return err
		}

		defer rep.Body.Close()

		if rep.StatusCode < 200 || rep.StatusCode >= 300 {
			return &rout.StatusError{StatusCode: rep.StatusCode, Reply: rout.Reply{Success: false, Error: http.StatusText(rep.StatusCode)}}
		}

		return nil
	}

	ticker := backoff.NewExponentialBackOff()

	for {
		if err = checkReady(); err == nil {
			return nil
		}

		// Log the error warning that we're still waiting to connect to Datum
		log.Warn().Err(err).Msg("waiting to connect to Datum")

		wait := time.After(ticker.NextBackOff())

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-wait:
		}
	}
}

const (
	userAgent    = "Datum API Client/v1"
	accept       = "application/json"
	acceptLang   = "en-US,en"
	acceptEncode = "gzip, deflate, br"
	contentType  = "application/json; charset=utf-8"
)

// NewRequest creates a new http request with the specified method, path, and data
func (s *APIv1) NewRequest(ctx context.Context, method, path string, data interface{}, params *url.Values) (req *http.Request, err error) {
	// Resolve the URL reference from the path
	url := s.endpoint.ResolveReference(&url.URL{Path: path})
	if params != nil && len(*params) > 0 {
		url.RawQuery = params.Encode()
	}

	var body io.ReadWriter

	switch {
	case data == nil:
		body = nil
	default:
		body = &bytes.Buffer{}
		if err = json.NewEncoder(body).Encode(data); err != nil {
			return nil, err
		}
	}

	if req, err = http.NewRequestWithContext(ctx, method, url.String(), body); err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Accept", accept)
	req.Header.Add("Accept-Language", acceptLang)
	req.Header.Add("Accept-Encoding", acceptEncode)
	req.Header.Add("Content-Type", contentType)

	var requestID string
	if requestID, _ = RequestIDFromContext(ctx); requestID == "" {
		requestID = ulid.Make().String()
	}

	req.Header.Add("X-Request-ID", requestID)

	var (
		ok    bool
		creds Credentials
	)

	if creds, ok = CredsFromContext(ctx); !ok {
		creds = s.creds
	}

	if creds != nil {
		var token string

		if token, err = creds.AccessToken(); err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+token)
	}

	if s.client.HTTPClient.Jar != nil {
		cookies := s.client.HTTPClient.Jar.Cookies(url)
		for _, cookie := range cookies {
			if cookie.Name == "csrf_token" {
				req.Header.Add("X-CSRF-TOKEN", cookie.Value)
			}
		}
	}

	return req, nil
}

// Do executes an http request against the server, performs error checking, and deserializes the response data into the specified struct
func (s *APIv1) Do(req *http.Request, data interface{}, checkStatus bool) (rep *http.Response, err error) {
	if rep, err = s.client.HTTPClient.Do(req); err != nil {
		return rep, err
	}

	defer rep.Body.Close()

	if checkStatus {
		if rep.StatusCode < 200 || rep.StatusCode >= 300 {
			serr := &rout.StatusError{
				StatusCode: rep.StatusCode,
			}

			if err = json.NewDecoder(rep.Body).Decode(&serr.Reply); err == nil {
				return rep, serr
			}

			serr.Reply = rout.Reply{Success: false}

			return rep, serr
		}
	}

	if data != nil && rep.StatusCode >= 200 && rep.StatusCode < 300 && rep.StatusCode != http.StatusNoContent {
		ct := rep.Header.Get("Content-Type")
		if ct != "" {
			mt, _, err := mime.ParseMediaType(ct)
			if err != nil {
				return nil, err
			}

			if mt != accept {
				return nil, err
			}
		}

		if err = json.NewDecoder(rep.Body).Decode(data); err != nil {
			return nil, err
		}
	}

	return rep, nil
}
