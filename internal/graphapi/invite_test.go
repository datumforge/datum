package graphapi_test

import (
	"context"
	"testing"
	"time"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/enums"
)

func (suite *GraphTestSuite) TestQueryInvite() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	invite := (&InviteBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		queryID     string
		client      *datumclient.DatumClient
		ctx         context.Context
		shouldCheck bool
		wantErr     bool
	}{
		{
			name:        "happy path",
			queryID:     invite.ID,
			client:      suite.client.datum,
			ctx:         reqCtx,
			shouldCheck: true,
			wantErr:     false,
		},
		{
			name:        "happy path with api token",
			queryID:     invite.ID,
			client:      suite.client.datumWithAPIToken,
			ctx:         context.Background(),
			shouldCheck: true,
			wantErr:     false,
		},
		{
			name:        "invalid id",
			queryID:     "allthefooandbar",
			client:      suite.client.datum,
			ctx:         reqCtx,
			shouldCheck: false,
			wantErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.shouldCheck {
				mock_fga.CheckAny(t, suite.client.fga, true)
			}

			resp, err := tc.client.GetInviteByID(tc.ctx, tc.queryID)

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
	reqCtx, err := userContext()
	require.NoError(t, err)

	// Existing user to invite to org
	existingUser := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	// Existing user already a member of org
	existingUser2 := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
	_ = (&OrgMemberBuilder{client: suite.client, OrgID: testOrgID, UserID: existingUser2.ID}).MustNew(reqCtx, t)

	testCases := []struct {
		name             string
		recipient        string
		orgID            string
		role             enums.Role
		client           *datumclient.DatumClient
		ctx              context.Context
		accessAllowed    bool
		expectedStatus   enums.InviteStatus
		expectedAttempts int64
		wantErr          bool
	}{
		{
			name:             "happy path, new user as member",
			recipient:        "meow@datum.net",
			orgID:            testOrgID,
			role:             enums.RoleMember,
			client:           suite.client.datum,
			ctx:              reqCtx,
			accessAllowed:    true,
			expectedStatus:   enums.InvitationSent,
			expectedAttempts: 0,
			wantErr:          false,
		},
		{
			name:             "re-invite new user as member using api token",
			recipient:        "meow@datum.net",
			orgID:            testOrgID,
			role:             enums.RoleMember,
			client:           suite.client.datumWithAPIToken,
			ctx:              context.Background(),
			accessAllowed:    true,
			expectedStatus:   enums.InvitationSent,
			expectedAttempts: 1,
			wantErr:          false,
		},
		{
			name:             "happy path, new user as admin using pat",
			recipient:        "woof@datum.net",
			orgID:            testOrgID,
			role:             enums.RoleAdmin,
			client:           suite.client.datumWithPAT,
			ctx:              context.Background(),
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
		// 	orgID:        testOrgID,
		// 	role:         enums.RoleOwner,
		//  client:      suite.client.datum,
		//  ctx:         reqCtx,
		//  accessAllowed: true,
		// 	wantErr:      true,
		// },
		{
			name:          "user not allowed to add to org",
			recipient:     "oink@datum.net",
			orgID:         testOrgID,
			role:          enums.RoleAdmin,
			client:        suite.client.datum,
			ctx:           reqCtx,
			accessAllowed: false,
			wantErr:       true,
		},
		{
			name:             "happy path, existing user as member",
			recipient:        existingUser.Email,
			orgID:            testOrgID,
			role:             enums.RoleMember,
			client:           suite.client.datum,
			ctx:              reqCtx,
			accessAllowed:    true,
			expectedStatus:   enums.InvitationSent,
			expectedAttempts: 0,
			wantErr:          false,
		},
		{
			name:             "user already a member, will still send an invite",
			recipient:        existingUser2.Email,
			orgID:            testOrgID,
			role:             enums.RoleMember,
			client:           suite.client.datum,
			ctx:              reqCtx,
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
			client:        suite.client.datum,
			ctx:           reqCtx,
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

			resp, err := tc.client.CreateInvite(tc.ctx, input)

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
			assert.Equal(t, testUser.ID, *resp.CreateInvite.Invite.RequestorID)
			assert.Equal(t, tc.expectedStatus, resp.CreateInvite.Invite.Status)
			assert.Equal(t, tc.expectedAttempts, resp.CreateInvite.Invite.SendAttempts)
			assert.WithinDuration(t, time.Now().UTC().AddDate(0, 0, 14), *resp.CreateInvite.Invite.Expires, time.Minute)
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteInvite() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	invite1 := (&InviteBuilder{client: suite.client}).MustNew(reqCtx, t)
	invite2 := (&InviteBuilder{client: suite.client}).MustNew(reqCtx, t)
	invite3 := (&InviteBuilder{client: suite.client}).MustNew(reqCtx, t)

	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, invite1.OwnerID)
	require.NoError(t, err)

	testCases := []struct {
		name    string
		queryID string
		client  *datumclient.DatumClient
		ctx     context.Context
		allowed bool
		wantErr bool
	}{
		{
			name:    "happy path",
			queryID: invite1.ID,
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
			wantErr: false,
		},
		{
			name:    "happy path, using api token",
			queryID: invite2.ID,
			client:  suite.client.datumWithAPIToken,
			ctx:     context.Background(),
			allowed: true,
			wantErr: false,
		},
		{
			name:    "happy path, using personal access token",
			queryID: invite3.ID,
			client:  suite.client.datumWithPAT,
			ctx:     context.Background(),
			allowed: true,
			wantErr: false,
		},
		{
			name:    "invalid id",
			queryID: "allthefooandbar",
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, true)

			resp, err := tc.client.DeleteInvite(tc.ctx, tc.queryID)

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
