package datumclient

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/datumforge/datum/pkg/sessions"
)

// WithAuthorizationAndSession adds the authorization header and session to the client request
func WithAuthorizationAndSession(accessToken string, session string) clientv2.RequestInterceptor {
	return func(
		ctx context.Context,
		req *http.Request,
		gqlInfo *clientv2.GQLRequestInfo,
		res interface{},
		next clientv2.RequestInterceptorFunc,
	) error {
		// setting authorization header if its not already set
		h := req.Header.Get("Authorization")
		if h == "" {
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		}

		// add session cookie
		if strings.Contains(req.Host, "localhost") {
			req.AddCookie(sessions.NewDevSessionCookie(session))
		} else {
			req.AddCookie(sessions.NewSessionCookie(session))
		}

		return next(ctx, req, gqlInfo, res)
	}
}

// WithAuthorization adds the authorization header to the client request
func WithAuthorization(accessToken string) clientv2.RequestInterceptor {
	return func(
		ctx context.Context,
		req *http.Request,
		gqlInfo *clientv2.GQLRequestInfo,
		res interface{},
		next clientv2.RequestInterceptorFunc,
	) error {
		// setting authorization header if its not already set
		h := req.Header.Get("Authorization")
		if h == "" {
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
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
