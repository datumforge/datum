package datumclient

import (
	"net/http"
)

// ClientOption allows us to configure the APIv1 client when it is created.
type ClientOption func(c *APIv1) error

func WithClient(client *http.Client) ClientOption {
	return func(c *APIv1) error {
		c.client = client
		return nil
	}
}

func WithCredentials(creds Credentials) ClientOption {
	return func(c *APIv1) error {
		c.creds = creds
		return nil
	}
}
