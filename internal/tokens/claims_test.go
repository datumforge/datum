package tokens_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/ulids"
)

func TestClaimsParseOrgID(t *testing.T) {
	claims := &tokens.Claims{}
	require.Equal(t, ulids.Null, claims.ParseOrgID())

	claims.OrgID = "notvalid"
	require.Equal(t, ulids.Null, claims.ParseOrgID())

	orgID := ulids.New()
	claims.OrgID = orgID.String()
	require.Equal(t, orgID, claims.ParseOrgID())
}

func TestClaimsParseUserID(t *testing.T) {
	claims := &tokens.Claims{}
	require.Equal(t, ulids.Null, claims.ParseUserID())

	claims.Subject = "notvalid"
	require.Equal(t, ulids.Null, claims.ParseUserID())

	userID := ulids.New()
	claims.Subject = userID.String()
	require.Equal(t, userID, claims.ParseUserID())
}
