package datumclient

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/datumforge/datum/pkg/httpsling"
	api "github.com/datumforge/datum/pkg/models"
)

type DatumClient struct {
	DatumRestClient
	DatumGraphClient
}

type DatumRestClient interface {
	Register(context.Context, *api.RegisterRequest) (*api.RegisterReply, error)
	Login(context.Context, *api.LoginRequest) (*api.LoginReply, error)
	Refresh(context.Context, *api.RefreshRequest) (*api.RefreshReply, error)
	Switch(context.Context, *api.SwitchOrganizationRequest) (*api.SwitchOrganizationReply, error)
	VerifyEmail(context.Context, *api.VerifyRequest) (*api.VerifyReply, error)
	ResendEmail(context.Context, *api.ResendRequest) (*api.ResendReply, error)
	ForgotPassword(context.Context, *api.ForgotPasswordRequest) (*api.ForgotPasswordReply, error)
	ResetPassword(context.Context, *api.ResetPasswordRequest) (*api.ResetPasswordReply, error)
	Invite(context.Context, *api.InviteRequest) (*api.InviteReply, error)
}

// A Reauthenticator generates new access and refresh pair given a valid refresh token
type Reauthenticator interface {
	Refresh(context.Context, *api.RefreshRequest) (*api.RefreshReply, error)
}

// TODO: we should provide a default graph query endpoint, but allow it to be overridden with options instead of this
type Config struct {
	BaseURL            string
	GraphQueryEndpoint string
}

// New creates a new API v1 client that implements the Datum Client interface
func New(config Config, opts ...ClientOption) (*DatumClient, error) {
	// configure rest client
	c, err := NewRestClient(config.BaseURL, opts...)
	if err != nil {
		return nil, err
	}

	api := c.(*APIv1)

	// configure graph client
	interceptors := []clientv2.RequestInterceptor{}

	token, err := api.creds.AccessToken()
	if err == nil {
		auth := Authorization{
			BearerToken: token,
		}

		interceptors = append(interceptors, auth.WithAuthorization())
	}

	// TODO: update the Options and interceptors as options to the client instead of hardcoding
	graphClient := NewClient(
		api.client.HTTPClient,                    // httpclient
		config.BaseURL+config.GraphQueryEndpoint, // full request path for graph
		&clientv2.Options{
			ParseDataAlongWithErrors: false,
		}, interceptors...,
	)

	return &DatumClient{
		c,
		graphClient,
	}, nil
}

// New creates a new API v1 client that implements the Datum Client interface
func NewRestClient(baseurl string, opts ...ClientOption) (DatumRestClient, error) {
	c := &APIv1{}

	ep, err := url.Parse(baseurl)
	if err != nil {
		return nil, err
	}

	c.baseurl = ep

	if c.client == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return nil, err
		}

		c.client = httpsling.Create(&httpsling.Config{
			BaseURL:    baseurl,
			Timeout:    5 * time.Second, // nolint: gomnd
			MaxRetries: 3,               // nolint: gomnd
		})

		c.client.HTTPClient.Jar = jar
	}

	c.client.SetDefaultUserAgent("Datum API Client/v1")
	c.client.SetDefaultHeader(httpsling.HeaderAccept, httpsling.ContentTypeJSON)
	c.client.SetDefaultHeader(httpsling.HeaderAcceptLanguage, "en-US,en")
	c.client.SetDefaultHeader(httpsling.HeaderAcceptEncoding, "gzip, deflate, br")
	c.client.SetDefaultHeader(httpsling.HeaderContentType, httpsling.ContentTypeJSONUTF8)

	// Apply our options
	for _, opt := range opts {
		if err = opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// APIv1 implements the DatumClient interface and provides methods to interact with the Datum API
type APIv1 struct {
	// baseurl is the base URL for requests
	baseurl *url.URL
	// config holds the httpsling configuration parameters to initialize the client
	config *httpsling.Config
	// creds is the default credentials used to authorize requests
	creds Credentials
	// client is the underlying HTTP client used to make requests provided by the httpsling library
	client *httpsling.Client
}

// Ensure the APIv1 implements the DatumClient interface
var _ DatumRestClient = &APIv1{}

// Register a new user with the Datum API
func (s *APIv1) Register(ctx context.Context, in *api.RegisterRequest) (out *api.RegisterReply, err error) {
	req := s.client.NewRequestBuilder(http.MethodPost, "/v1/register")
	req.Body(in)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	if err := resp.ScanJSON(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// Login to the Datum API
func (s *APIv1) Login(ctx context.Context, in *api.LoginRequest) (out *api.LoginReply, err error) {
	req := s.client.NewRequestBuilder(http.MethodPost, "/v1/login")
	req.Body(in)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	if err := resp.ScanJSON(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// Refresh a user's access token
func (s *APIv1) Refresh(ctx context.Context, in *api.RefreshRequest) (out *api.RefreshReply, err error) {
	req := s.client.NewRequestBuilder(http.MethodPost, "/v1/refresh")
	req.Body(in)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	if err := resp.ScanJSON(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// Switch the current organization context
func (s *APIv1) Switch(ctx context.Context, in *api.SwitchOrganizationRequest) (out *api.SwitchOrganizationReply, err error) {
	req := s.client.NewRequestBuilder(http.MethodPost, "/v1/switch")
	req.Body(in)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	if err := resp.ScanJSON(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// VerifyEmail verifies the email address of a user
func (s *APIv1) VerifyEmail(ctx context.Context, in *api.VerifyRequest) (out *api.VerifyReply, err error) {
	req := s.client.NewRequestBuilder(http.MethodPost, "/v1/verify")
	req.Body(in)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	if err := resp.ScanJSON(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// ResendEmail resends the verification email to the user
func (s *APIv1) ResendEmail(ctx context.Context, in *api.ResendRequest) (out *api.ResendReply, err error) {
	req := s.client.NewRequestBuilder(http.MethodPost, "/v1/resend")
	req.Body(in)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	if err := resp.ScanJSON(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// ForgotPassword sends a password reset email to the user
func (s *APIv1) ForgotPassword(ctx context.Context, in *api.ForgotPasswordRequest) (out *api.ForgotPasswordReply, err error) {
	req := s.client.NewRequestBuilder(http.MethodPost, "/v1/forgot-password")
	req.Body(in)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	if err := resp.ScanJSON(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// ResetPassword resets the user's password
func (s *APIv1) ResetPassword(ctx context.Context, in *api.ResetPasswordRequest) (out *api.ResetPasswordReply, err error) {
	req := s.client.NewRequestBuilder(http.MethodPost, "/v1/password-reset")
	req.Body(in)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	if err := resp.ScanJSON(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// Invite a user to an organization
func (s *APIv1) Invite(ctx context.Context, in *api.InviteRequest) (out *api.InviteReply, err error) {
	req := s.client.NewRequestBuilder(http.MethodPost, "/v1/invite")
	req.Body(in)

	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	if err := resp.ScanJSON(&out); err != nil {
		return nil, err
	}

	return out, nil
}
