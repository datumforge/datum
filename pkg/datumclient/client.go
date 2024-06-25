package datumclient

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"

	"github.com/datumforge/datum/pkg/httpsling"
	api "github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/sessions"
)

const (
	// cookieExpiryMinutes is the duration for which the cookies are valid
	cookieExpiryMinutes = 10 * time.Minute
)

// DatumClient wraps the Datum API client methods to form a single client interface
type DatumClient struct {
	DatumRestClient
	DatumGraphClient
}

// A Reauthenticator generates new access and refresh pair given a valid refresh token
type Reauthenticator interface {
	Refresh(context.Context, *api.RefreshRequest) (*api.RefreshReply, error)
}

// NewWithDefaults creates a new API v1 client with default configuration
func NewWithDefaults(opts ...ClientOption) (*DatumClient, error) {
	config := NewDefaultConfig()

	return New(config, opts...)
}

// New creates a new API v1 client that implements the Datum Client interface
func New(config Config, opts ...ClientOption) (*DatumClient, error) {
	// configure rest client
	c, err := NewRestClient(config, opts...)
	if err != nil {
		return nil, err
	}

	api := c.(*APIv1)

	// create the graph client
	// use api.Config instead of config because some fields are updated in NewRestClient
	graphClient := NewClient(
		api.HTTPSlingClient.HTTPClient,
		graphRequestPath(api.Config),
		&api.Config.Clientv2Options,
		api.Config.Interceptors...,
	)

	return &DatumClient{
		c,
		graphClient,
	}, nil
}

// newHTTPClient creates a new HTTP sling client with the given configuration
func newHTTPClient(config Config) (*httpsling.Client, error) {
	// copy the values from the base config to the httpsling config
	if config.HTTPSling == nil {
		config.HTTPSling = &httpsling.Config{}
	}

	if config.HTTPSling.BaseURL == "" {
		config.HTTPSling.BaseURL = config.BaseURL.String()
	}

	client := httpsling.Create(config.HTTPSling)

	// set the default cookie jar
	if err := client.SetDefaultCookieJar(); err != nil {
		return nil, err
	}

	return client, nil
}

// APIv1 implements the DatumClient interface and provides methods to interact with the Datum API
type APIv1 struct {
	// Config is the configuration for the APIv1 client
	Config Config
	// HTTPSlingClient is the HTTP client for the APIv1 client
	HTTPSlingClient *httpsling.Client
}

// Config is the configuration for the APIv1 client
func (c *DatumClient) Config() Config {
	api := c.DatumRestClient.(*APIv1)

	return api.Config
}

// HTTPSlingClient is the http client for the APIv1 client
func (c *DatumClient) HTTPSlingClient() *httpsling.Client {
	api := c.DatumRestClient.(*APIv1)

	return api.HTTPSlingClient
}

// AccessToken returns the access token cached on the client or an error if it is not
// available. This method is primarily used for testing but can be used to fetch the
// access token for debugging or inspection if necessary.
func (c *DatumClient) AccessToken() (_ string, err error) {
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
func (c *DatumClient) RefreshToken() (_ string, err error) {
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
func (c *DatumClient) SetAuthTokens(access, refresh string) error {
	if c.HTTPSlingClient().HTTPClient.Jar == nil {
		return ErrNoCookieJarSet
	}

	// The URL for the cookies
	u := c.Config().BaseURL.ResolveReference(&url.URL{Path: "/"})

	// Set the cookies on the client
	cookies := make([]*http.Cookie, 0, 2) //nolint:mnd
	if access != "" {
		cookies = append(cookies, &http.Cookie{
			Name:     "access_token",
			Value:    access,
			Expires:  time.Now().Add(cookieExpiryMinutes),
			HttpOnly: true,
			Secure:   true,
		})
	}

	if refresh != "" {
		cookies = append(cookies, &http.Cookie{
			Name:    "refresh_token",
			Value:   refresh,
			Expires: time.Now().Add(cookieExpiryMinutes),
			Secure:  true,
		})
	}

	c.Config().HTTPSling.CookieJar.SetCookies(u, cookies)

	return nil
}

// ClearAuthTokens clears the access and refresh tokens on the client Jar.
func (c *DatumClient) ClearAuthTokens() {
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
func (c *DatumClient) Cookies() ([]*http.Cookie, error) {
	if c.HTTPSlingClient().HTTPClient.Jar == nil {
		return nil, ErrNoCookieJarSet
	}

	cookies := c.HTTPSlingClient().HTTPClient.Jar.Cookies(c.Config().BaseURL)

	return cookies, nil
}

// GetSessionFromCookieJar parses the cookie jar for the session cookie
func (c *DatumClient) GetSessionFromCookieJar() (sessionID string, err error) {
	cookies, err := c.Cookies()
	if err != nil {
		return "", err
	}

	cookieName := sessions.DefaultCookieName

	// Use the dev cookie when running on localhost
	if strings.Contains(c.Config().BaseURL.Host, "localhost") {
		cookieName = sessions.DevCookieName
	}

	for _, c := range cookies {
		if c.Name == cookieName {
			return c.Value, nil
		}
	}

	return "", nil
}

// ClearAuthTokens clears the access and refresh tokens on the client Jar.
func (c *DatumClient) GetAuthTokensFromCookieJar() *oauth2.Token {
	token := oauth2.Token{}

	if cookies, err := c.Cookies(); err == nil {
		fmt.Println("checking cookies")
		for _, cookie := range cookies {
			fmt.Println("cookie", cookie.Name, cookie.Value)
			switch cookie.Name {
			case "access_token":
				token.AccessToken = cookie.Value
			case "refresh_token":
				token.RefreshToken = cookie.Value
			}
		}
	}

	return &token
}
