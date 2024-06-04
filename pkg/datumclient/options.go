package datumclient

import (
	"net/url"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/datumforge/datum/pkg/httpsling"
)

// ClientOption allows us to configure the APIv1 client when it is created
type ClientOption func(c *APIv1) error

// WithClient sets the client for the APIv1 client
func WithClient(client *httpsling.Client) ClientOption {
	return func(c *APIv1) error {
		c.Config.HTTPSlingClient = client
		return nil
	}
}

// WithCredentials sets the credentials for the APIv1 client
func WithCredentials(creds Credentials) ClientOption {
	return func(c *APIv1) error {
		c.Config.Credentials = creds
		return nil
	}
}

// WithHTTPSlingConfig sets the config for the APIv1 client
func WithHTTPSlingConfig(config *httpsling.Config) ClientOption {
	return func(c *APIv1) error {
		c.Config.HTTPSling = config
		return nil
	}
}

// WithInterceptors sets the interceptors for the APIv1 client
func WithInterceptors(interceptors clientv2.RequestInterceptor) ClientOption {
	return func(c *APIv1) error {
		c.Config.Interceptors = []clientv2.RequestInterceptor{interceptors}
		return nil
	}
}

// WithClientv2Option sets the clientv2 options for the APIv1 client
func WithClientv2Option(option clientv2.Options) ClientOption {
	return func(c *APIv1) error {
		c.Config.Clientv2Options = option
		return nil
	}
}

// WithGraphQueryEndpoint sets the graph query endpoint for the APIv1 client
func WithGraphQueryEndpoint(endpoint string) ClientOption {
	return func(c *APIv1) error {
		c.Config.GraphQLPath = endpoint
		return nil
	}
}

// WithBaseURL sets the base URL for the APIv1 client
func WithBaseURL(baseURL *url.URL) ClientOption {
	return func(c *APIv1) error {
		c.Config.BaseURL = baseURL
		return nil
	}
}

// WithToken sets the token for the APIv1 client
func WithToken(token string) ClientOption {
	return func(c *APIv1) error {
		c.Config.Token = token
		return nil
	}
}

// WithTokenRefresh sets the token refresh for the APIv1 client
func WithTokenRefresh(tokenRefresh string) ClientOption {
	return func(c *APIv1) error {
		c.Config.TokenRefresh = tokenRefresh
		return nil
	}
}
