package graphapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
)

func TestQuery_GroupMembers(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	group := (&GroupBuilder{}).MustNew(reqCtx)

	groupMember, err := group.Members(reqCtx)
	require.NoError(t, err)
	require.Len(t, groupMember, 1)

	testCases := []struct {
		name     string
		queryID  string
		expected *ent.GroupMembership
	}{
		{
			name:     "happy path, get group member by group id",
			queryID:  group.ID,
			expected: groupMember[0],
		},
		{
			name:     "invalid-id",
			queryID:  "tacos-for-dinner",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			groupID := tc.queryID
			whereInput := datumclient.GroupMembershipWhereInput{
				GroupID: &groupID,
			}

			resp, err := authClient.gc.GetGroupMembersByGroupID(reqCtx, &whereInput)
			require.NoError(t, err)

			if tc.expected == nil {
				assert.Empty(t, resp.GroupMemberships.Edges)

				return
			}

			require.NotNil(t, resp)
			require.NotNil(t, resp.GroupMemberships)
			assert.Equal(t, tc.expected.UserID, resp.GroupMemberships.Edges[0].Node.UserID)
			assert.Equal(t, tc.expected.Role, resp.GroupMemberships.Edges[0].Node.Role)
		})
	}

	// delete created group
	(&GroupCleanup{GroupID: group.ID}).MustDelete(reqCtx)
}

func TestQuery_CreateGroupMembers(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	group1 := (&GroupBuilder{}).MustNew(reqCtx)

	groupMember, err := group1.Members(reqCtx)
	require.NoError(t, err)
	require.Len(t, groupMember, 1)

	testUser1 := (&UserBuilder{}).MustNew(reqCtx)
	testUser2 := (&UserBuilder{}).MustNew(reqCtx)

	testCases := []struct {
		name    string
		groupID string
		userID  string
		role    enums.Role
		errMsg  string
	}{
		{
			name:    "happy path, add admin",
			groupID: group1.ID,
			userID:  testUser1.ID,
			role:    enums.RoleAdmin,
		},
		{
			name:    "happy path, add member",
			groupID: group1.ID,
			userID:  testUser2.ID,
			role:    enums.RoleMember,
		},
		{
			name:    "owner relation not valid for groups",
			groupID: group1.ID,
			userID:  testUser2.ID,
			role:    enums.RoleOwner,
			errMsg:  "OWNER is not a valid GroupMembershipRole",
		},
		{
			name:    "duplicate user, different role",
			groupID: group1.ID,
			userID:  testUser1.ID,
			role:    enums.RoleMember,
			errMsg:  "constraint failed",
		},
		{
			name:    "invalid user",
			groupID: group1.ID,
			userID:  "not-a-valid-user-id",
			role:    enums.RoleMember,
			errMsg:  "constraint failed",
		},
		{
			name:    "invalid group",
			groupID: "not-a-valid-group-id",
			userID:  testUser1.ID,
			role:    enums.RoleMember,
			errMsg:  "constraint failed",
		},
		{
			name:    "invalid role",
			groupID: group1.ID,
			userID:  testUser1.ID,
			role:    enums.Invalid,
			errMsg:  "not a valid GroupMembershipRole",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			if tc.errMsg == "" {
				mockWriteTuplesAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)
			}

			role := tc.role
			input := datumclient.CreateGroupMembershipInput{
				GroupID: tc.groupID,
				UserID:  tc.userID,
				Role:    &role,
			}

			resp, err := authClient.gc.AddUserToGroupWithRole(reqCtx, input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateGroupMembership)
			assert.Equal(t, tc.userID, resp.CreateGroupMembership.GroupMembership.UserID)
			assert.Equal(t, tc.groupID, resp.CreateGroupMembership.GroupMembership.GroupID)
			assert.Equal(t, tc.role, resp.CreateGroupMembership.GroupMembership.Role)
		})
	}

	// delete created group and users
	(&GroupCleanup{GroupID: group1.ID}).MustDelete(reqCtx)
	(&UserCleanup{UserID: testUser1.ID}).MustDelete(reqCtx)
	(&UserCleanup{UserID: testUser2.ID}).MustDelete(reqCtx)
}

func TestQuery_UpdateGroupMembers(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&GroupMemberBuilder{}).MustNew(reqCtx)

	testCases := []struct {
		name   string
		role   enums.Role
		errMsg string
	}{
		{
			name: "happy path, update to admin from member",
			role: enums.RoleAdmin,
		},
		{
			name: "happy path, update to member from admin",
			role: enums.RoleMember,
		},
		{
			name:   "invalid role",
			role:   enums.Invalid,
			errMsg: "not a valid GroupMembershipRole",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			if tc.errMsg == "" {
				mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)
			}

			role := tc.role
			input := datumclient.UpdateGroupMembershipInput{
				Role: &role,
			}

			resp, err := authClient.gc.UpdateUserRoleInGroup(reqCtx, om.ID, input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdateGroupMembership)
			assert.Equal(t, tc.role, resp.UpdateGroupMembership.GroupMembership.Role)
		})
	}

	// delete created group and users
	(&GroupMemberCleanup{ID: om.ID}).MustDelete(reqCtx)
}

func TestQuery_DeleteGroupMembers(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&GroupMemberBuilder{}).MustNew(reqCtx)

	mockDeleteTuplesAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)

	resp, err := authClient.gc.RemoveUserFromGroup(reqCtx, om.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.DeleteGroupMembership)
	assert.Equal(t, om.ID, resp.DeleteGroupMembership.DeletedID)
}
