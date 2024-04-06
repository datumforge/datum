package graphapi_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
)

func (suite *GraphTestSuite) TestQueryInvite() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	invite := (&InviteBuilder{client: suite.client}).MustNew(reqCtx, t)

	user := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
	inviteExistingUser := (&InviteBuilder{client: suite.client, Recipient: user.Email}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		queryID     string
		shouldCheck bool
		expected    *ent.Invite
		wantErr     bool
	}{
		{
			name:        "happy path",
			queryID:     invite.ID,
			shouldCheck: true,
			expected:    invite,
			wantErr:     false,
		},
		{
			name:        "invite accepted, should be deleted and not found",
			queryID:     inviteExistingUser.ID,
			shouldCheck: true,
			expected:    nil,
			wantErr:     true,
		},
		{
			name:        "invalid id",
			queryID:     "allthefooandbar",
			shouldCheck: false,
			expected:    nil,
			wantErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.shouldCheck {
				mock_fga.CheckAny(t, suite.client.fga, true)
				mock_fga.ListAny(t, suite.client.fga, []string{fmt.Sprintf("organization:%s", invite.OwnerID)})
			}

			resp, err := suite.client.datum.GetInvite(reqCtx, tc.queryID)

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

func (suite *GraphTestSuite) TestMutationCreateInvite() {
	t := suite.T()

	// setup user context
	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	orgAdmin := (&UserBuilder{client: suite.client}).MustNew(ctx, t)

	// setup valid user context
	userCtx, err := auth.NewTestContextWithValidUser(orgAdmin.ID)
	if err != nil {
		t.Fatal()
	}

	reqCtx := context.WithValue(userCtx.Request().Context(), echocontext.EchoContextKey, userCtx)

	userCtx.SetRequest(ec.Request().WithContext(reqCtx))

	// Org to invite users to
	org := (&OrganizationBuilder{client: suite.client}).MustNew(userCtx.Request().Context(), t)

	// Existing user to invite to org
	existingUser := (&UserBuilder{client: suite.client}).MustNew(userCtx.Request().Context(), t)

	// Existing user already a member of org
	existingUser2 := (&UserBuilder{client: suite.client}).MustNew(userCtx.Request().Context(), t)
	_ = (&OrgMemberBuilder{client: suite.client, OrgID: org.ID, UserID: existingUser2.ID}).MustNew(reqCtx, t)

	testCases := []struct {
		name             string
		recipient        string
		orgID            string
		role             enums.Role
		accessAllowed    bool
		expectedStatus   enums.InviteStatus
		expectedAttempts int64
		wantErr          bool
	}{
		{
			name:             "happy path, new user as member",
			recipient:        "meow@datum.net",
			orgID:            org.ID,
			role:             enums.RoleMember,
			accessAllowed:    true,
			expectedStatus:   enums.InvitationSent,
			expectedAttempts: 0,
			wantErr:          false,
		},
		{
			name:             "re-invite new user as member",
			recipient:        "meow@datum.net",
			orgID:            org.ID,
			role:             enums.RoleMember,
			accessAllowed:    true,
			expectedStatus:   enums.InvitationSent,
			expectedAttempts: 1,
			wantErr:          false,
		},
		{
			name:             "happy path, new user as admin",
			recipient:        "woof@datum.net",
			orgID:            org.ID,
			role:             enums.RoleAdmin,
			accessAllowed:    true,
			expectedStatus:   enums.InvitationSent,
			expectedAttempts: 0,
			wantErr:          false,
		},
		// TODO: uncomment with https://github.com/datumforge/datum/issues/405
		// {
		// 	name:         "new user as owner should fail",
		// 	existingUser: false,
		// 	recipient:    "woof@datum.net",
		// 	orgID:        org.ID,
		// 	role:         enums.RoleOwner,
		//  accessAllowed: true,
		// 	wantErr:      true,
		// },
		{
			name:          "user not allowed to add to org",
			recipient:     "oink@datum.net",
			orgID:         org.ID,
			role:          enums.RoleAdmin,
			accessAllowed: false,
			wantErr:       true,
		},
		{
			name:             "happy path, existing user as member",
			recipient:        existingUser.Email,
			orgID:            org.ID,
			role:             enums.RoleMember,
			accessAllowed:    true,
			expectedStatus:   enums.InvitationSent,
			expectedAttempts: 0,
			wantErr:          false,
		},
		{
			name:             "user already a member, will still send an invite",
			recipient:        existingUser2.Email,
			orgID:            org.ID,
			role:             enums.RoleMember,
			accessAllowed:    true,
			expectedStatus:   enums.InvitationSent,
			expectedAttempts: 0,
			wantErr:          false,
		},
		{
			name:          "invalid org",
			recipient:     existingUser.Email,
			orgID:         "boommeowboom",
			role:          enums.RoleMember,
			accessAllowed: true,
			wantErr:       true,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, tc.accessAllowed)

			if tc.accessAllowed {
				mock_fga.ListAny(t, suite.client.fga, []string{fmt.Sprintf("organization:%s", tc.orgID)})
			}

			role := tc.role
			input := datumclient.CreateInviteInput{
				Recipient: tc.recipient,
				OwnerID:   tc.orgID,
				Role:      &role,
			}

			resp, err := suite.client.datum.CreateInvite(reqCtx, input)

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
			assert.Equal(t, tc.expectedAttempts, resp.CreateInvite.Invite.SendAttempts)
			assert.WithinDuration(t, time.Now().UTC().AddDate(0, 0, 14), resp.CreateInvite.Invite.Expires, time.Minute)
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteInvite() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	invite := (&InviteBuilder{client: suite.client}).MustNew(reqCtx, t)

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
			name:     "invalid id",
			queryID:  "allthefooandbar",
			allowed:  true,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if !tc.wantErr {
				mock_fga.CheckAny(t, suite.client.fga, true)
			}

			resp, err := suite.client.datum.DeleteInvite(reqCtx, tc.queryID)

			if tc.wantErr {
				require.Error(t, err)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			// assert equal
			assert.Equal(t, tc.queryID, resp.DeleteInvite.DeletedID)
		})
	}
}
