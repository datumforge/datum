package datumclient

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestGetTokensFromCookies(t *testing.T) {
	tests := []struct {
		name      string
		cookies   []*http.Cookie
		wantToken *oauth2.Token
	}{
		{
			name: "happy path",
			cookies: []*http.Cookie{
				{Name: "access_token", Value: "access_token"},
				{Name: "refresh_token", Value: "refresh_token"},
			},
			wantToken: &oauth2.Token{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			},
		},
		{
			name: "missing access token",
			cookies: []*http.Cookie{
				{Name: "refresh_token", Value: "refresh_token"},
			},
			wantToken: &oauth2.Token{
				AccessToken:  "",
				RefreshToken: "refresh_token",
			},
		},
		{
			name: "missing refresh token",
			cookies: []*http.Cookie{
				{Name: "access_token", Value: "access_token"},
			},
			wantToken: &oauth2.Token{
				AccessToken:  "access_token",
				RefreshToken: "",
			},
		},
		{
			name:    "missing both tokens",
			cookies: []*http.Cookie{},
			wantToken: &oauth2.Token{
				AccessToken:  "",
				RefreshToken: "",
			},
		},
		{
			name:    "nil tokens",
			cookies: nil,
			wantToken: &oauth2.Token{
				AccessToken:  "",
				RefreshToken: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTokensFromCookies(tt.cookies)
			assert.Equal(t, tt.wantToken, got, "getTokensFromCookies() = %v, want %v", got, tt.wantToken)
		})
	}
}

func TestGetTokensFromCookieRequest(t *testing.T) {
	ctx := context.Background()

	accessCookie := http.Cookie{Name: "access_token", Value: "access_token"}
	refreshCookie := http.Cookie{Name: "refresh_token", Value: "refresh_token"}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080?session=sessionvalue", nil)
	require.NoError(t, err)

	req.AddCookie(&accessCookie)
	req.AddCookie(&refreshCookie)

	devRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080?session=sessionvalue", nil)
	require.NoError(t, err)

	devRequest.AddCookie(&accessCookie)
	devRequest.AddCookie(&refreshCookie)

	type args struct {
		r *http.Request
	}

	tests := []struct {
		name        string
		args        args
		wantToken   *oauth2.Token
		wantSession string
	}{
		{
			name: "default session",
			args: args{
				r: req,
			},
			wantToken: &oauth2.Token{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			},
			wantSession: "sessionvalue",
		},
		{
			name: "dev session",
			args: args{
				r: devRequest,
			},
			wantToken: &oauth2.Token{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			},
			wantSession: "sessionvalue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, gotSession := getTokensFromCookieRequest(tt.args.r)
			assert.Equal(t, tt.wantToken, gotToken, "getTokensFromCookieRequest() gotToken = %v, want %v", gotToken, tt.wantToken)
			assert.Equal(t, tt.wantSession, gotSession, "getTokensFromCookieRequest() gotSession = %v, want %v", gotSession, tt.wantSession)
		})
	}
}
