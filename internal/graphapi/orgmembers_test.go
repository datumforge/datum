package graphapi_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/hooks"
	mock_fga "github.com/datumforge/datum/internal/fga/mockery"
)

func TestQuery_OrgMembers(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{client: client}).MustNew(reqCtx, t)

	orgMember, err := org1.Members(reqCtx)
	require.NoError(t, err)
	require.Len(t, orgMember, 1)

	testCases := []struct {
		name     string
		queryID  string
		expected *ent.OrgMembership
	}{
		{
			name:     "happy path, get org member by org id",
			queryID:  org1.ID,
			expected: orgMember[0],
		},
		{
			name:     "invalid-id",
			queryID:  "tacos-for-dinner",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			orgID := tc.queryID
			whereInput := datumclient.OrgMembershipWhereInput{
				OrganizationID: &orgID,
			}

			resp, err := client.datum.GetOrgMembersByOrgID(reqCtx, &whereInput)
			require.NoError(t, err)

			if tc.expected == nil {
				assert.Empty(t, resp.OrgMemberships.Edges)

				return
			}

			require.NotNil(t, resp)
			require.NotNil(t, resp.OrgMemberships)
			assert.Equal(t, tc.expected.UserID, resp.OrgMemberships.Edges[0].Node.UserID)
			assert.Equal(t, tc.expected.Role, resp.OrgMemberships.Edges[0].Node.Role)
		})
	}

	// delete created org
	(&OrganizationCleanup{client: client, OrgID: org1.ID}).MustDelete(reqCtx, t)
}

func TestQuery_CreateOrgMembers(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{client: client}).MustNew(reqCtx, t)
	personalOrg := (&OrganizationBuilder{client: client, PersonalOrg: true}).MustNew(reqCtx, t)
	listObjects := []string{fmt.Sprintf("organization:%s", org1.ID), fmt.Sprintf("organization:%s", personalOrg.ID)}

	orgMember, err := org1.Members(reqCtx)
	require.NoError(t, err)
	require.Len(t, orgMember, 1)

	testUser1 := (&UserBuilder{client: client}).MustNew(reqCtx, t)
	testUser2 := (&UserBuilder{client: client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		orgID    string
		userID   string
		role     enums.Role
		checkOrg bool
		errMsg   string
	}{
		{
			name:     "happy path, add admin",
			orgID:    org1.ID,
			userID:   testUser1.ID,
			role:     enums.RoleAdmin,
			checkOrg: true,
		},
		{
			name:     "happy path, add member",
			orgID:    org1.ID,
			userID:   testUser2.ID,
			role:     enums.RoleMember,
			checkOrg: true,
		},
		{
			name:     "duplicate user, different role",
			orgID:    org1.ID,
			userID:   testUser1.ID,
			role:     enums.RoleMember,
			checkOrg: true,
			errMsg:   "constraint failed",
		},
		{
			name:     "add user to personal org not allowed",
			orgID:    personalOrg.ID,
			userID:   testUser1.ID,
			role:     enums.RoleMember,
			checkOrg: true,
			errMsg:   hooks.ErrPersonalOrgsNoMembers.Error(),
		},
		{
			name:     "invalid user",
			orgID:    org1.ID,
			userID:   "not-a-valid-user-id",
			role:     enums.RoleMember,
			checkOrg: true,
			errMsg:   "constraint failed", // TODO: better error messaging: https://github.com/datumforge/datum/issues/415
		},
		{
			name:     "invalid org",
			orgID:    "not-a-valid-org-id",
			userID:   testUser1.ID,
			role:     enums.RoleMember,
			checkOrg: true,
			errMsg:   "organization not found",
		},
		{
			name:     "invalid role",
			orgID:    org1.ID,
			userID:   testUser1.ID,
			role:     enums.Invalid,
			checkOrg: false,
			errMsg:   "not a valid OrgMembershipRole",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(client.fga)

			if tc.errMsg == "" {
				mock_fga.WriteAny(t, client.fga)
			}

			if tc.checkOrg {
				// checks for adding orgs to ensure not a personal org
				mock_fga.ListAny(t, client.fga, listObjects)
			}

			role := tc.role
			input := datumclient.CreateOrgMembershipInput{
				OrgID:  tc.orgID,
				UserID: tc.userID,
				Role:   &role,
			}

			resp, err := client.datum.AddUserToOrgWithRole(reqCtx, input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateOrgMembership)
			assert.Equal(t, tc.userID, resp.CreateOrgMembership.OrgMembership.UserID)
			assert.Equal(t, tc.orgID, resp.CreateOrgMembership.OrgMembership.OrganizationID)
			assert.Equal(t, tc.role, resp.CreateOrgMembership.OrgMembership.Role)
		})
	}

	// delete created org and users
	(&OrganizationCleanup{client: client, OrgID: org1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: client, UserID: testUser1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: client, UserID: testUser2.ID}).MustDelete(reqCtx, t)
}

func TestQuery_UpdateOrgMembers(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&OrgMemberBuilder{client: client}).MustNew(reqCtx, t)

	testCases := []struct {
		name       string
		role       enums.Role
		tupleWrite bool
		errMsg     string
	}{
		{
			name:       "happy path, update to admin from member",
			tupleWrite: true,
			role:       enums.RoleAdmin,
		},
		{
			name:       "happy path, update to member from admin",
			tupleWrite: true,
			role:       enums.RoleMember,
		},
		{
			name:       "update to same role",
			tupleWrite: false, // nothing should change
			role:       enums.RoleMember,
		},
		{
			name:       "invalid role",
			role:       enums.Invalid,
			tupleWrite: false,
			errMsg:     "not a valid OrgMembershipRole",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(client.fga)

			if tc.tupleWrite {
				mock_fga.WriteAny(t, client.fga)
			}

			role := tc.role
			input := datumclient.UpdateOrgMembershipInput{
				Role: &role,
			}

			resp, err := client.datum.UpdateUserRoleInOrg(reqCtx, om.ID, input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdateOrgMembership)
			assert.Equal(t, tc.role, resp.UpdateOrgMembership.OrgMembership.Role)
		})
	}

	// delete created org and users
	(&OrgMemberCleanup{client: client, ID: om.ID}).MustDelete(reqCtx, t)
}

func TestQuery_DeleteOrgMembers(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&OrgMemberBuilder{client: client}).MustNew(reqCtx, t)

	mock_fga.WriteAny(t, client.fga)

	resp, err := client.datum.RemoveUserFromOrg(reqCtx, om.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.DeleteOrgMembership)
	assert.Equal(t, om.ID, resp.DeleteOrgMembership.DeletedID)
}
