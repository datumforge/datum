package graphapi_test

import (
	"fmt"
	"testing"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/enums"
)

func (suite *GraphTestSuite) TestQueryGroupMembers() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	group := (&GroupBuilder{client: suite.client}).MustNew(reqCtx, t)

	// allow access to group
	checkCtx := privacy.DecisionContext(reqCtx, privacy.Allow)

	groupMember, err := group.Members(checkCtx)
	require.NoError(t, err)
	require.Len(t, groupMember, 1)

	testCases := []struct {
		name        string
		queryID     string
		allowed     bool
		expected    *ent.GroupMembership
		errExpected bool
	}{
		{
			name:     "happy path, get group member by group id",
			queryID:  group.ID,
			allowed:  true,
			expected: groupMember[0],
		},
		{
			name:        "get group member by group id, no access",
			queryID:     group.ID,
			allowed:     false,
			expected:    nil,
			errExpected: true,
		},
		{
			name:     "invalid-id",
			queryID:  "tacos-for-dinner",
			allowed:  true,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			groupID := tc.queryID
			whereInput := datumclient.GroupMembershipWhereInput{
				GroupID: &groupID,
			}

			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.GetGroupMembersByGroupID(reqCtx, &whereInput)

			if tc.errExpected {
				require.Error(t, err)
				assert.ErrorContains(t, err, "deny rule")

				return
			}

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
	(&GroupCleanup{client: suite.client, GroupID: group.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestMutationCreateGroupMembers() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	group1 := (&GroupBuilder{client: suite.client}).MustNew(reqCtx, t)

	// allow access to group
	checkCtx := privacy.DecisionContext(reqCtx, privacy.Allow)

	groupMember, err := group1.Members(checkCtx)
	require.NoError(t, err)
	require.Len(t, groupMember, 1)

	testUser1 := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
	testUser2 := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	reqCtx, err = auth.NewTestContextWithOrgID(testUser.ID, group1.OwnerID)
	require.NoError(t, err)

	testCases := []struct {
		name    string
		groupID string
		userID  string
		role    enums.Role
		allowed bool
		check   bool
		list    bool
		errMsg  string
	}{
		{
			name:    "happy path, add admin",
			groupID: group1.ID,
			userID:  testUser1.ID,
			role:    enums.RoleAdmin,
			allowed: true,
			check:   true,
			list:    true,
		},
		{
			name:    "happy path, add member",
			groupID: group1.ID,
			userID:  testUser2.ID,
			role:    enums.RoleMember,
			allowed: true,
			check:   true,
			list:    true,
		},
		{
			name:    "add member, no access",
			groupID: group1.ID,
			userID:  testUser2.ID,
			role:    enums.RoleMember,
			allowed: false,
			check:   true,
			list:    false,
			errMsg:  "you are not authorized to perform this action",
		},
		{
			name:    "owner relation not valid for groups",
			groupID: group1.ID,
			userID:  testUser2.ID,
			role:    enums.RoleOwner,
			allowed: true,
			check:   false,
			list:    false,
			errMsg:  "OWNER is not a valid GroupMembershipRole",
		},
		{
			name:    "duplicate user, different role",
			groupID: group1.ID,
			userID:  testUser1.ID,
			role:    enums.RoleMember,
			allowed: true,
			check:   true,
			list:    true,
			errMsg:  "already exists",
		},
		{
			name:    "invalid user",
			groupID: group1.ID,
			userID:  "not-a-valid-user-id",
			role:    enums.RoleMember,
			allowed: true,
			check:   true,
			list:    true,
			errMsg:  "constraint failed",
		},
		{
			name:    "invalid group",
			groupID: "not-a-valid-group-id",
			userID:  testUser1.ID,
			role:    enums.RoleMember,
			allowed: true,
			check:   true,
			list:    true,
			errMsg:  "constraint failed",
		},
		{
			name:    "invalid role",
			groupID: group1.ID,
			userID:  testUser1.ID,
			role:    enums.RoleInvalid,
			allowed: true,
			check:   false,
			list:    false,
			errMsg:  "not a valid GroupMembershipRole",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.errMsg == "" {
				mock_fga.WriteAny(t, suite.client.fga)
			}

			if tc.check {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			if tc.list {
				mock_fga.ListAny(t, suite.client.fga, []string{fmt.Sprintf("organization:%s", group1.OwnerID)})
			}

			role := tc.role
			input := datumclient.CreateGroupMembershipInput{
				GroupID: tc.groupID,
				UserID:  tc.userID,
				Role:    &role,
			}

			resp, err := suite.client.datum.AddUserToGroupWithRole(reqCtx, input)

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
	(&GroupCleanup{client: suite.client, GroupID: group1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: suite.client, UserID: testUser1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: suite.client, UserID: testUser2.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestMutationUpdateGroupMembers() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	gm := (&GroupMemberBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name    string
		role    enums.Role
		allowed bool
		check   bool
		errMsg  string
	}{
		{
			name:    "happy path, update to admin from member",
			role:    enums.RoleAdmin,
			allowed: true,
			check:   true,
		},
		{
			name:    "happy path, update to member from admin",
			role:    enums.RoleMember,
			allowed: true,
			check:   true,
		},
		{
			name:    "invalid role",
			role:    enums.RoleInvalid,
			errMsg:  "not a valid GroupMembershipRole",
			allowed: true,
			check:   false,
		},
		{
			name:    "no access",
			role:    enums.RoleMember,
			errMsg:  "you are not authorized to perform this action",
			allowed: false,
			check:   true,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.errMsg == "" {
				mock_fga.WriteAny(t, suite.client.fga)
			}

			if tc.check {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			role := tc.role
			input := datumclient.UpdateGroupMembershipInput{
				Role: &role,
			}

			resp, err := suite.client.datum.UpdateUserRoleInGroup(reqCtx, gm.ID, input)

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
	(&GroupMemberCleanup{client: suite.client, ID: gm.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestMutationDeleteGroupMembers() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	om := (&GroupMemberBuilder{client: suite.client}).MustNew(reqCtx, t)

	mock_fga.WriteAny(t, suite.client.fga)
	mock_fga.CheckAny(t, suite.client.fga, true)

	resp, err := suite.client.datum.RemoveUserFromGroup(reqCtx, om.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.DeleteGroupMembership)
	assert.Equal(t, om.ID, resp.DeleteGroupMembership.DeletedID)
}
