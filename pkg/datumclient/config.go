package datumclient

import (
	"net/http"
	"net/url"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/datumforge/datum/pkg/httpsling"
)

// Config is the configuration for the Datum API client
type Config struct {
	// BaseURL is the base URL for the Datum API
	BaseURL *url.URL `json:"baseUrl" yaml:"base_url" default:"http://localhost:17608"`
	// GraphQLPath is the path to the GraphQL endpoint
	GraphQLPath string `json:"graphqlPath" default:"/query"`
	// HTTPSling is the configuration for the HTTPSling client
	HTTPSling *httpsling.Config
	// Interceptors are the request interceptors for the graph client
	Interceptors []clientv2.RequestInterceptor
	// Credentials are the credentials for the client
	Credentials Credentials
	// Clientv2Options are the options for the graph client
	Clientv2Options clientv2.Options
}

// graphRequestPath returns the full URL path for the GraphQL endpoint
func graphRequestPath(config Config) string {
	baseurl := config.BaseURL.String()

	return baseurl + config.GraphQLPath
}

// NewDefaultConfig returns a new default configuration for the Datum API client
func NewDefaultConfig() Config {
	return defaultClientConfig
}

var defaultClientConfig = Config{
	BaseURL: &url.URL{
		Scheme: "http",
		Host:   "localhost:17608",
	},
	GraphQLPath: "/query",
	HTTPSling: &httpsling.Config{
		Headers: &http.Header{
			"Accept":          []string{httpsling.ContentTypeJSONUTF8},
			"Accept-Language": []string{"en-US,en"},
			"Content-Type":    []string{httpsling.ContentTypeJSONUTF8},
		},
	},
	Interceptors:    []clientv2.RequestInterceptor{},
	Clientv2Options: clientv2.Options{ParseDataAlongWithErrors: false},
}
