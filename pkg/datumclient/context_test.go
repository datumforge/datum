package datumclient_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	api "github.com/datumforge/datum/pkg/datumclient"
)

func TestAPIContex(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	creds, ok := api.CredsFromContext(ctx)
	require.False(t, ok, "should not be able to collect creds from default context")
	require.Nil(t, creds, "creds should be returned as nil")

	authCtx := api.ContextWithToken(ctx, "access_token")

	creds, ok = api.CredsFromContext(authCtx)
	require.True(t, ok, "should be able to collect creds from context")
	require.NotNil(t, creds, "credentials should have been returned")

	accessToken, _ := creds.AccessToken()
	require.Equal(t, "access_token", accessToken)

	// Should be able to update the authentication on an access token
	authCtx = api.ContextWithToken(authCtx, "different_access_token")
	creds, ok = api.CredsFromContext(authCtx)
	require.True(t, ok, "could not fetch replaced access token")

	token, _ := creds.AccessToken()
	require.Equal(t, "different_access_token", token)
}
