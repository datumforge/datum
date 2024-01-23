package graphapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	mock_fga "github.com/datumforge/datum/internal/fga/mockery"
)

func TestQuery_GroupMembers(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	group := (&GroupBuilder{client: client}).MustNew(reqCtx, t)

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

			resp, err := client.datum.GetGroupMembersByGroupID(reqCtx, &whereInput)
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
	(&GroupCleanup{client: client, GroupID: group.ID}).MustDelete(reqCtx, t)
}

func TestQuery_CreateGroupMembers(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	group1 := (&GroupBuilder{client: client}).MustNew(reqCtx, t)

	groupMember, err := group1.Members(reqCtx)
	require.NoError(t, err)
	require.Len(t, groupMember, 1)

	testUser1 := (&UserBuilder{client: client}).MustNew(reqCtx, t)
	testUser2 := (&UserBuilder{client: client}).MustNew(reqCtx, t)

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
			defer mock_fga.ClearMocks(client.fga)

			if tc.errMsg == "" {
				mock_fga.WriteAny(t, client.fga)
			}

			role := tc.role
			input := datumclient.CreateGroupMembershipInput{
				GroupID: tc.groupID,
				UserID:  tc.userID,
				Role:    &role,
			}

			resp, err := client.datum.AddUserToGroupWithRole(reqCtx, input)

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
	(&GroupCleanup{client: client, GroupID: group1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: client, UserID: testUser1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: client, UserID: testUser2.ID}).MustDelete(reqCtx, t)
}

func TestQuery_UpdateGroupMembers(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&GroupMemberBuilder{client: client}).MustNew(reqCtx, t)

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
			defer mock_fga.ClearMocks(client.fga)

			if tc.errMsg == "" {
				mock_fga.WriteAny(t, client.fga)
			}

			role := tc.role
			input := datumclient.UpdateGroupMembershipInput{
				Role: &role,
			}

			resp, err := client.datum.UpdateUserRoleInGroup(reqCtx, om.ID, input)

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

	// delete created group
	(&GroupMemberCleanup{client: client, ID: om.ID}).MustDelete(reqCtx, t)
}

func TestQuery_DeleteGroupMembers(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&GroupMemberBuilder{client: client}).MustNew(reqCtx, t)

	mock_fga.WriteAny(t, client.fga)

	resp, err := client.datum.RemoveUserFromGroup(reqCtx, om.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.DeleteGroupMembership)
	assert.Equal(t, om.ID, resp.DeleteGroupMembership.DeletedID)
}
