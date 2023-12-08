package authtest_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	authtest "github.com/datumforge/datum/internal/httpserve/middleware/authtest"
	"github.com/datumforge/datum/internal/tokens"
)

// This test generates an example token with fake RSA keys for use in examples,
// documentation and other tests that don't need a valid token (since it will expire).
func TestGenerateToken(t *testing.T) {
	t.Skip("comment the skip out if you want to generate a token")

	srv, err := authtest.NewServer()
	require.NoError(t, err, "could not start authtest server")

	defer srv.Close()

	claims := &tokens.Claims{
		UserID:      "Rusty Shackleford",
		Email:       "rustys@datum.net",
		OrgID:       "01H6PGFG71N0AFEVTK3NJB71T9",
		ParentOrgID: "01H6PGFTK2X53RGG2KMSGR2M61",
		Tier:        "Pro",
	}

	accessToken, refreshToken, err := srv.CreateTokenPair(claims)
	require.NoError(t, err, "could not generate access token")

	// Log the tokens then fail the test so the tokens are printed out.
	t.Logf("access token: %s", accessToken)
	t.Logf("refresh token: %s", refreshToken)
	t.FailNow()
}
