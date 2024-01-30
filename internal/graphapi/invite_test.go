package graphapi_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	mock_fga "github.com/datumforge/datum/internal/fga/mockery"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
)

func TestQuery_Invite(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	invite := (&InviteBuilder{client: client}).MustNew(reqCtx, t)

	user := (&UserBuilder{client: client}).MustNew(reqCtx, t)
	inviteExistingUser := (&InviteBuilder{client: client, Recipient: user.Email}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		allowed  bool
		expected *ent.Invite
		wantErr  bool
	}{
		{
			name:     "happy path",
			queryID:  invite.ID,
			allowed:  true,
			expected: invite,
			wantErr:  false,
		},
		{
			name:     "invite accepted, should be deleted and not found",
			queryID:  inviteExistingUser.ID,
			allowed:  true,
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "invalid id",
			queryID:  "allthefooandbar",
			allowed:  true,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(client.fga)

			if !tc.wantErr {
				mock_fga.ListAny(t, client.fga, []string{fmt.Sprintf("organization:%s", invite.OwnerID)})
			}

			resp, err := client.datum.GetInvite(reqCtx, tc.queryID)

			if tc.wantErr {
				require.Error(t, err)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Invite)
		})
	}
}

func TestQuery_CreateInvite(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	orgAdmin := (&UserBuilder{client: client}).MustNew(ctx, t)

	// setup valid user context
	userCtx, err := auth.NewTestContextWithValidUser(orgAdmin.ID)
	if err != nil {
		t.Fatal()
	}

	reqCtx := context.WithValue(userCtx.Request().Context(), echocontext.EchoContextKey, userCtx)

	userCtx.SetRequest(ec.Request().WithContext(reqCtx))

	// Org to invite users to
	org := (&OrganizationBuilder{client: client}).MustNew(userCtx.Request().Context(), t)

	// // Existing user to invite to org
	existingUser := (&UserBuilder{client: client}).MustNew(userCtx.Request().Context(), t)

	// // Existing user already a member of org
	existingUser2 := (&UserBuilder{client: client}).MustNew(userCtx.Request().Context(), t)
	_ = (&OrgMemberBuilder{client: client, OrgID: org.ID, UserID: existingUser2.ID}).MustNew(reqCtx, t)

	testCases := []struct {
		name           string
		existingUser   bool
		recipient      string
		orgID          string
		role           enums.Role
		expectedStatus enums.InviteStatus
		wantErr        bool
	}{
		{
			name:           "happy path, new user as member",
			existingUser:   false,
			recipient:      "meow@datum.net",
			orgID:          org.ID,
			role:           enums.RoleMember,
			expectedStatus: enums.InvitationSent,
			wantErr:        false,
		},
		{
			name:           "happy path, existing user as member",
			existingUser:   true,
			recipient:      existingUser.Email,
			orgID:          org.ID,
			role:           enums.RoleMember,
			expectedStatus: enums.InvitationAccepted,
			wantErr:        false,
		},
		{
			name:      "user already a member",
			recipient: existingUser2.Email,
			orgID:     org.ID,
			role:      enums.RoleMember,
			wantErr:   true,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(client.fga)

			mock_fga.CheckAny(t, client.fga, true)
			mock_fga.ListAny(t, client.fga, []string{fmt.Sprintf("organization:%s", tc.orgID)})

			if tc.existingUser {
				mock_fga.WriteAny(t, client.fga)
			}

			role := tc.role
			input := datumclient.CreateInviteInput{
				Recipient: tc.recipient,
				OwnerID:   tc.orgID,
				Role:      &role,
			}

			resp, err := client.datum.CreateInvite(reqCtx, input)

			if tc.wantErr {
				require.Error(t, err)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			// Assert matching fields
			assert.Equal(t, tc.orgID, resp.CreateInvite.Invite.Owner.ID)
			assert.Equal(t, tc.role, resp.CreateInvite.Invite.Role)
			assert.Equal(t, orgAdmin.ID, resp.CreateInvite.Invite.RequestorID)
			assert.Equal(t, tc.expectedStatus, resp.CreateInvite.Invite.Status)
			assert.WithinDuration(t, time.Now().AddDate(0, 0, 14), resp.CreateInvite.Invite.Expires, time.Minute)
		})
	}
}
