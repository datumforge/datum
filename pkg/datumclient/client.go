package datumclient

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/httpsling"
	api "github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/sessions"
	"golang.org/x/oauth2"
)

// DatumClient is the interface that wraps the Datum API client methods
type DatumClient struct {
	DatumRestClient
	DatumGraphClient
}

// DatumRestClient is the interface that wraps the Datum API REST client methods
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

// New creates a new API v1 client that implements the Datum Client interface
func New(config Config, opts ...ClientOption) (*DatumClient, error) {
	// configure rest client
	c, err := NewRestClient(config, opts...)
	if err != nil {
		return nil, err
	}

	api := c.(*APIv1)

	token, err := api.Config.Credentials.AccessToken()
	if err == nil {
		auth := Authorization{
			BearerToken: token,
		}

		config.Interceptors = append(config.Interceptors, auth.WithAuthorization())
	}

	graphClient := NewClient(api.Config.HTTPSlingClient.HTTPClient, graphRequestPath(config), &config.Clientv2Options, config.Interceptors...)

	return &DatumClient{
		c,
		graphClient,
	}, nil
}

// New creates a new API v1 client that implements the Datum Client interface
func NewRestClient(config Config, opts ...ClientOption) (DatumRestClient, error) {
	c := &APIv1{}

	c.Config.HTTPSlingClient = httpsling.Create(c.Config.HTTPSling)

	// Apply our options
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// APIv1 implements the DatumClient interface and provides methods to interact with the Datum API
type APIv1 struct {
	// Config is the configuration for the APIv1 client
	Config Config
}

// AccessToken returns the access token cached on the client or an error if it is not
// available. This method is primarily used for testing but can be used to fetch the
// access token for debugging or inspection if necessary.
func (c *APIv1) AccessToken() (_ string, err error) {
	var cookies []*http.Cookie
	if cookies, err = c.Cookies(); err != nil {
		return "", err
	}

	for _, cookie := range cookies {
		if cookie.Name == "access_token" {
			return cookie.Value, nil
		}
	}

	return "", err
}

// RefreshToken returns the refresh token cached on the client or an error if it is not
// available. This method is primarily used for testing but can be used to fetch the
// refresh token for debugging or inspection if necessary.
func (c *APIv1) RefreshToken() (_ string, err error) {
	var cookies []*http.Cookie
	if cookies, err = c.Cookies(); err != nil {
		return "", err
	}

	for _, cookie := range cookies {
		if cookie.Name == "refresh_token" {
			return cookie.Value, nil
		}
	}

	return "", err
}

// SetAuthTokens is a helper function to set the access and refresh tokens on the
// client cookie jar.
func (c *APIv1) SetAuthTokens(access, refresh string) error {
	if c.Config.HTTPSling.CookieJar == nil {
		return errors.New("client does not have a cookie jar, cannot set cookies")
	}

	// The URL for the cookies
	u := c.Config.BaseURL.ResolveReference(&url.URL{Path: "/"})

	// Set the cookies on the client
	cookies := make([]*http.Cookie, 0, 2)
	if access != "" {
		cookies = append(cookies, &http.Cookie{
			Name:     "access_token",
			Value:    access,
			Expires:  time.Now().Add(10 * time.Minute),
			HttpOnly: true,
			Secure:   true,
		})
	}

	if refresh != "" {
		cookies = append(cookies, &http.Cookie{
			Name:    "refresh_token",
			Value:   refresh,
			Expires: time.Now().Add(10 * time.Minute),
			Secure:  true,
		})
	}
	c.Config.HTTPSling.CookieJar.SetCookies(u, cookies)
	return nil
}

// ClearAuthTokens clears the access and refresh tokens on the client Jar.
func (c *APIv1) ClearAuthTokens() {
	if cookies, err := c.Cookies(); err == nil {
		// Expire the access and refresh cookies.
		for _, cookie := range cookies {
			switch cookie.Name {
			case "access_token":
				cookie.MaxAge = -1
			case "refresh_token":
				cookie.MaxAge = -1
			}
		}
	}
}

// Returns the cookies set from the previous request(s) on the client Jar.
func (c *APIv1) Cookies() (_ []*http.Cookie, err error) {
	if c.Config.HTTPSling.CookieJar == nil {
		return nil, err
	}

	cookies := c.Config.HTTPSling.CookieJar.Cookies(c.Config.BaseURL)
	return cookies, nil
}

// Ensure the APIv1 implements the DatumClient interface
var _ DatumRestClient = &APIv1{}

// Register a new user with the Datum API
func (s *APIv1) Register(ctx context.Context, in *api.RegisterRequest) (out *api.RegisterReply, err error) {
	req := s.Config.HTTPSlingClient.NewRequestBuilder(http.MethodPost, "/v1/register")
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
	req := s.Config.HTTPSlingClient.NewRequestBuilder(http.MethodPost, "/v1/login")
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
	req := s.Config.HTTPSlingClient.NewRequestBuilder(http.MethodPost, "/v1/refresh")
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
	req := s.Config.HTTPSlingClient.NewRequestBuilder(http.MethodPost, "/v1/switch")
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
	req := s.Config.HTTPSlingClient.NewRequestBuilder(http.MethodPost, "/v1/verify")
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
	req := s.Config.HTTPSlingClient.NewRequestBuilder(http.MethodPost, "/v1/resend")
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
	req := s.Config.HTTPSlingClient.NewRequestBuilder(http.MethodPost, "/v1/forgot-password")
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
	req := s.Config.HTTPSlingClient.NewRequestBuilder(http.MethodPost, "/v1/password-reset")
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
	req := s.Config.HTTPSlingClient.NewRequestBuilder(http.MethodPost, "/v1/invite")
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

func refreshToken(ctx context.Context, refresh string) (*api.RefreshReply, error) {
	c, err := New(DefaultClientConfig)
	if err != nil {
		return nil, err
	}

	req := api.RefreshRequest{
		RefreshToken: refresh,
	}

	return c.Refresh(ctx, &req)
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

// getTokensFromCookieRequest parses the HTTP Request for cookies and returns the session and access and refresh tokens
// this is used for the oauth login flow
func getTokensFromCookieRequest(r *http.Request) (token *oauth2.Token, session string) {
	// parse cookies
	cookies := r.Cookies()

	// get session from query string
	session = r.URL.Query().Get("session")

	return getTokensFromCookies(cookies), session
}

// GetSessionFromCookieJar parses the cookie jar for the session cookie
func (s *APIv1) GetSessionFromCookieJar() (sessionID string, err error) {
	u, err := url.Parse(s.Config.BaseURL.String())
	if err != nil {
		return "", err
	}

	cookies := s.Config.HTTPSling.CookieJar.Cookies(u)
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
