package datumclient

import (
	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/datumforge/datum/pkg/httpsling"
)

type Config struct {
	BaseURL            string
	GraphQLPath        string
	HTTPSling          *httpsling.Config
	Interceptors       []clientv2.RequestInterceptor
	HTTPSlingClient    *httpsling.Client
	Credentials        Credentials
	Clientv2Options    clientv2.Options
	GraphQueryEndpoint string
	Token              string
	TokenRefresh       string
}
