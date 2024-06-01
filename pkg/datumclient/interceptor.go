package datumclient

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/datumforge/datum/pkg/httpsling"
	"github.com/datumforge/datum/pkg/sessions"
)

// Authorization contains the bearer token and optional session cookie
type Authorization struct {
	// BearerToken is the bearer token to be used in the authorization header
	// this can be the access token, api token, or personal access token
	BearerToken string
	// Session is the session cookie to be used in the request
	// this is required for requests using the access token
	Session string
}

// WithAuthorizationAndSession adds the authorization header and session to the client request
func (a Authorization) WithAuthorization() clientv2.RequestInterceptor {
	return func(
		ctx context.Context,
		req *http.Request,
		gqlInfo *clientv2.GQLRequestInfo,
		res interface{},
		next clientv2.RequestInterceptorFunc,
	) error {
		// setting authorization header if its not already set
		h := req.Header.Get(httpsling.HeaderAuthorization)
		if h == "" {
			auth := httpsling.BearerAuth{
				Token: a.BearerToken,
			}

			auth.Apply(req)
		}

		// add session cookie
		if a.Session != "" {
			if strings.Contains(req.Host, "localhost") {
				req.AddCookie(sessions.NewDevSessionCookie(a.Session))
			} else {
				req.AddCookie(sessions.NewSessionCookie(a.Session))
			}
		}

		return next(ctx, req, gqlInfo, res)
	}
}

// WithLoggingInterceptor adds a http debug logging interceptor
func WithLoggingInterceptor() clientv2.RequestInterceptor {
	return func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		fmt.Println("Request header sent:", req.Header)
		fmt.Println("Request body sent:", req.Body)

		return next(ctx, req, gqlInfo, res)
	}
}

// WithEmptyInterceptor adds an empty interceptor
func WithEmptyInterceptor() clientv2.RequestInterceptor {
	return func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		return next(ctx, req, gqlInfo, res)
	}
}
