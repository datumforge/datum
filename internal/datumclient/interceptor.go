package datumclient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/datumforge/datum/internal/sessions"
)

// WithAuthorization adds the authorization header and session to the client request
func WithAuthorization(accessToken string, session string) clientv2.RequestInterceptor {
	return func(
		ctx context.Context,
		req *http.Request,
		gqlInfo *clientv2.GQLRequestInfo,
		w http.ResponseWriter,
		res interface{},
		next clientv2.RequestInterceptorFunc,
	) error {
		// setting authorization header if its not already set
		h := req.Header.Get("Authorization")
		if h == "" {
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		}

		session := sessions.New("datum")
		session.Set()

		// add session cookie
		req.AddCookie(sessions.n

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
