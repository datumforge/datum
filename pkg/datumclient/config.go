package datumclient

import (
	"net/http"
	"net/url"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/datumforge/datum/pkg/httpsling"
)

// Config is the configuration for the Datum API client
type Config struct {
	BaseURL         *url.URL `json:"base_url" yaml:"base_url" default:"http://localhost:17608"`
	GraphQLPath     string   `json:"graphqlpath" default:"/query"`
	HTTPSling       *httpsling.Config
	Interceptors    []clientv2.RequestInterceptor
	HTTPSlingClient *httpsling.Client
	Credentials     Credentials
	Clientv2Options clientv2.Options
	Token           string
	TokenRefresh    string
}

// graphRequestPath returns the full URL path for the GraphQL endpoint
func graphRequestPath(config Config) string {
	baseurl := config.BaseURL.String()
	return baseurl + config.GraphQLPath
}

var DefaultClientConfig = Config{
	BaseURL: &url.URL{
		Scheme: "http",
		Host:   "localhost:17608",
	},
	HTTPSling: &httpsling.Config{
		Headers: &http.Header{
			"Accept":          []string{httpsling.ContentTypeJSONUTF8},
			"Accept-Language": []string{"en-US,en"},
			"Accept-Encoding": []string{"gzip, deflate, br"},
			"Content-Type":    []string{httpsling.ContentTypeJSONUTF8},
		},
	},
	GraphQLPath:     "/query",
	Interceptors:    []clientv2.RequestInterceptor{},
	Clientv2Options: clientv2.Options{ParseDataAlongWithErrors: false},
}
