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
)

func TestQuery_OrgMembers(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)

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
				OrgID: &orgID,
			}

			resp, err := authClient.gc.GetOrgMembersByOrgID(reqCtx, &whereInput)
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
	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
}

func TestQuery_CreateOrgMembers(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org1 := (&OrganizationBuilder{}).MustNew(reqCtx)
	personalOrg := (&OrganizationBuilder{PersonalOrg: true}).MustNew(reqCtx)
	listObjects := []string{fmt.Sprintf("organization:%s", org1.ID), fmt.Sprintf("organization:%s", personalOrg.ID)}

	orgMember, err := org1.Members(reqCtx)
	require.NoError(t, err)
	require.Len(t, orgMember, 1)

	testUser1 := (&UserBuilder{}).MustNew(reqCtx)
	testUser2 := (&UserBuilder{}).MustNew(reqCtx)

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
			if tc.errMsg == "" {
				mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)
			}

			if tc.checkOrg {
				// checks for adding orgs to ensure not a personal org
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			}

			role := tc.role
			input := datumclient.CreateOrgMembershipInput{
				OrgID:  tc.orgID,
				UserID: tc.userID,
				Role:   &role,
			}

			resp, err := authClient.gc.AddUserToOrgWithRole(reqCtx, input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateOrgMembership)
			assert.Equal(t, tc.userID, resp.CreateOrgMembership.OrgMembership.UserID)
			assert.Equal(t, tc.orgID, resp.CreateOrgMembership.OrgMembership.OrgID)
			assert.Equal(t, tc.role, resp.CreateOrgMembership.OrgMembership.Role)
		})
	}

	// delete created org and users
	(&OrganizationCleanup{OrgID: org1.ID}).MustDelete(reqCtx)
	(&UserCleanup{UserID: testUser1.ID}).MustDelete(reqCtx)
	(&UserCleanup{UserID: testUser2.ID}).MustDelete(reqCtx)
}

func TestQuery_UpdateOrgMembers(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&OrgMemberBuilder{}).MustNew(reqCtx)

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
			if tc.tupleWrite {
				mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)
			}

			role := tc.role
			input := datumclient.UpdateOrgMembershipInput{
				Role: &role,
			}

			resp, err := authClient.gc.UpdateUserRoleInOrg(reqCtx, om.ID, input)

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
	(&OrgMemberCleanup{ID: om.ID}).MustDelete(reqCtx)
}

func TestQuery_DeleteOrgMembers(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&OrgMemberBuilder{}).MustNew(reqCtx)

	mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)

	resp, err := authClient.gc.RemoveUserFromOrg(reqCtx, om.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.DeleteOrgMembership)
	assert.Equal(t, om.ID, resp.DeleteOrgMembership.DeletedID)
}
