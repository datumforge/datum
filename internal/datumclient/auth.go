package datumclient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

// WithAccessToken adds the authorization header to the client request
func WithAccessToken(accessToken string) clientv2.RequestInterceptor {
	return func(
		ctx context.Context,
		req *http.Request,
		gqlInfo *clientv2.GQLRequestInfo,
		res interface{},
		next clientv2.RequestInterceptorFunc,
	) error {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

		return next(ctx, req, gqlInfo, res)
	}
}
