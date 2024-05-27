package datumclient

import "github.com/datumforge/datum/pkg/httpsling"

// ClientOption allows us to configure the APIv1 client when it is created
type ClientOption func(c *APIv1) error

// WithClient sets the client for the APIv1 client
func WithClient(client *httpsling.Client) ClientOption {
	return func(c *APIv1) error {
		c.client = client
		return nil
	}
}

// WithCredentials sets the credentials for the APIv1 client
func WithCredentials(creds Credentials) ClientOption {
	return func(c *APIv1) error {
		c.creds = creds
		return nil
	}
}

// WithHTTPSlingConfig sets the config for the APIv1 client
func WithHTTPSlingConfig(config *httpsling.Config) ClientOption {
	return func(c *APIv1) error {
		c.config = config
		return nil
	}
}
