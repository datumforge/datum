package graphapi_test

import (
	"testing"
	"time"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/enums"
)

func (suite *GraphTestSuite) TestQueryInvite() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// Org to invite users to
	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup valid user context with org
	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, org.ID)
	require.NoError(t, err)

	invite := (&InviteBuilder{client: suite.client}).MustNew(reqCtx, t)

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
	ctx, err := userContext()
	require.NoError(t, err)

	orgAdmin := (&UserBuilder{client: suite.client}).MustNew(ctx, t)

	userCtx, err := auth.NewTestContextWithValidUser(orgAdmin.ID)
	require.NoError(t, err)

	// Org to invite users to
	org := (&OrganizationBuilder{client: suite.client}).MustNew(userCtx, t)

	// setup valid user context with org
	reqCtx, err := auth.NewTestContextWithOrgID(orgAdmin.ID, org.ID)
	require.NoError(t, err)

	// Existing user to invite to org
	existingUser := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	// Existing user already a member of org
	existingUser2 := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
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
			accessAllowed: false,
			wantErr:       true,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, tc.accessAllowed)

			role := tc.role
			input := datumclient.CreateInviteInput{
				Recipient: tc.recipient,
				OwnerID:   &tc.orgID,
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

	// Org to invite users to
	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup valid user context with org
	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, org.ID)
	require.NoError(t, err)

	invite := (&InviteBuilder{client: suite.client}).MustNew(reqCtx, t)

	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, invite.OwnerID)
	require.NoError(t, err)

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

			mock_fga.CheckAny(t, suite.client.fga, true)

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
