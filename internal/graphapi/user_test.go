package graphapi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/graphapi"
	auth "github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
)

func TestQuery_User(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	user1 := (&UserBuilder{}).MustNew(ctx)
	user2 := (&UserBuilder{}).MustNew(ctx)

	// setup valid user context
	userCtx, err := auth.NewTestContextWithValidUser(user1.ID)
	if err != nil {
		t.Fatal()
	}

	reqCtx := context.WithValue(userCtx.Request().Context(), echocontext.EchoContextKey, userCtx)

	userCtx.SetRequest(ec.Request().WithContext(reqCtx))

	testCases := []struct {
		name     string
		queryID  string
		expected *ent.User
		errorMsg string
	}{
		{
			name:     "happy path user",
			queryID:  user1.ID,
			expected: user1,
		},
		{
			name:     "valid user, but no auth",
			queryID:  user2.ID,
			errorMsg: "user not found",
		},
		{
			name:     "invalid-id",
			queryID:  "tacos-for-dinner",
			errorMsg: "user not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			if tc.errorMsg == "" {
				// placeholder until authz is enforced on org members
				// at which point this needs to be the correct organization id
				listObjects := []string{fmt.Sprintf("organization:%s", "test")}
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			}

			resp, err := authClient.gc.GetUserByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.User)
		})
	}

	(&UserCleanup{UserID: user1.ID}).MustDelete(reqCtx)
	(&UserCleanup{UserID: user2.ID}).MustDelete(reqCtx)
}

func TestQuery_Users(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	user1 := (&UserBuilder{}).MustNew(reqCtx)
	user2 := (&UserBuilder{}).MustNew(reqCtx)

	t.Run("Get Users", func(t *testing.T) {
		resp, err := authClient.gc.GetAllUsers(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Users.Edges)

		// make sure only the current user is returned
		assert.Equal(t, len(resp.Users.Edges), 1)

		// set new context with another user id
		ec, err := auth.NewTestContextWithValidUser(user1.ID)
		if err != nil {
			t.Fatal()
		}

		reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

		ec.SetRequest(ec.Request().WithContext(reqCtx))

		resp, err = authClient.gc.GetAllUsers(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Users.Edges)

		// only user that is making the request should be returned
		assert.Equal(t, len(resp.Users.Edges), 1)

		user1Found := false
		user2Found := false
		for _, o := range resp.Users.Edges {
			if o.Node.ID == user1.ID {
				user1Found = true
			} else if o.Node.ID == user2.ID {
				user2Found = true
			}
		}

		// only user 1 should be found
		if !user1Found {
			t.Errorf("user 1 was expected to be found but was not")
		}

		// user 2 should not be found
		if user2Found {
			t.Errorf("user 2 was not expected to be found but was returned")
		}

		// TODO: Add admin test that should be able to do a `GET` on all users
		// user1Found := false
		// user2Found := false
		// for _, o := range resp.Users.Edges {
		// 	if o.Node.ID == user1.ID {
		// 		user1Found = true
		// 	} else if o.Node.ID == user2.ID {
		// 		user2Found = true
		// 	}
		// }

		// // if one of the users isn't found, fail the test
		// if !user1Found || !user2Found {
		// 	t.Fail()
		// }
	})
}

func TestMutation_CreateUserNoAuth(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	// Bypass auth checks to ensure input checks for now
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	displayName := gofakeit.LetterN(50)

	weakPassword := "notsecure"
	strongPassword := "my&supers3cr3tpassw0rd!"

	testCases := []struct {
		name      string
		userInput datumclient.CreateUserInput
		errorMsg  string
	}{
		{
			name: "happy path user",
			userInput: datumclient.CreateUserInput{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				DisplayName: &displayName,
				Email:       gofakeit.Email(),
				Password:    &strongPassword,
			},
			errorMsg: "",
		},
		{
			name: "no email",
			userInput: datumclient.CreateUserInput{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				DisplayName: &displayName,
				Email:       "",
			},
			errorMsg: "mail: no address",
		},
		{
			name: "no first name",
			userInput: datumclient.CreateUserInput{
				FirstName:   "",
				LastName:    gofakeit.LastName(),
				DisplayName: &displayName,
				Email:       gofakeit.Email(),
			},
			errorMsg: "value is less than the required length",
		},
		{
			name: "no last name",
			userInput: datumclient.CreateUserInput{
				FirstName:   gofakeit.FirstName(),
				LastName:    "",
				DisplayName: &displayName,
				Email:       gofakeit.Email(),
			},
			errorMsg: "value is less than the required length",
		},
		{
			name: "no display name, should default to email",
			userInput: datumclient.CreateUserInput{
				FirstName: gofakeit.FirstName(),
				LastName:  gofakeit.LastName(),
				Email:     gofakeit.Email(),
			},
			errorMsg: "",
		},
		{
			name: "weak password",
			userInput: datumclient.CreateUserInput{
				FirstName: gofakeit.FirstName(),
				LastName:  gofakeit.LastName(),
				Email:     gofakeit.Email(),
				Password:  &weakPassword,
			},
			errorMsg: auth.ErrPasswordTooWeak.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			if tc.errorMsg == "" {
				// mock writes to create personal org membership
				mockWriteAny(authClient.mockCtrl, authClient.mc, ctx, nil)
			}

			resp, err := authClient.gc.CreateUser(ctx, tc.userInput)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateUser.User)

			// Make sure provided values match
			assert.Equal(t, tc.userInput.FirstName, resp.CreateUser.User.FirstName)
			assert.Equal(t, tc.userInput.LastName, resp.CreateUser.User.LastName)
			assert.Equal(t, tc.userInput.Email, resp.CreateUser.User.Email)

			// display name defaults to email if not provided
			if tc.userInput.DisplayName == nil {
				assert.Equal(t, tc.userInput.Email, resp.CreateUser.User.DisplayName)
			} else {
				assert.Equal(t, *tc.userInput.DisplayName, resp.CreateUser.User.DisplayName)
			}

			// ensure a user setting was created
			assert.NotNil(t, resp.CreateUser.User.Setting)

			orgs := resp.CreateUser.User.OrgMemberships
			require.Len(t, orgs, 1)

			// setup valid user context
			ec, err := auth.NewTestContextWithValidUser(resp.CreateUser.User.ID)
			if err != nil {
				t.Fatal()
			}

			reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

			ec.SetRequest(ec.Request().WithContext(reqCtx))

			// mocks to check for org access
			listObjects := []string{fmt.Sprintf("organization:%s", orgs[0].OrgID)}
			mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
			mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)

			// Bypass auth checks to ensure input checks for now
			reqCtx = privacy.DecisionContext(reqCtx, privacy.Allow)

			personalOrg, err := authClient.gc.GetOrganizationByID(reqCtx, orgs[0].OrgID)
			require.NoError(t, err)

			assert.True(t, personalOrg.Organization.PersonalOrg)
			// make sure there is only one user
			require.Len(t, personalOrg.Organization.Members, 1)
			// make sure user was added as owner
			assert.Equal(t, personalOrg.Organization.Members[0].Role.String(), "OWNER")

			(&UserCleanup{UserID: resp.CreateUser.User.ID}).MustDelete(ctx)
		})
	}
}

func TestMutation_CreateUser(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	displayName := gofakeit.LetterN(50)

	// weakPassword := "notsecure"
	strongPassword := "my&supers3cr3tpassw0rd!"

	testCases := []struct {
		name      string
		userInput datumclient.CreateUserInput
		errorMsg  string
	}{
		{
			name: "no auth create user",
			userInput: datumclient.CreateUserInput{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				DisplayName: &displayName,
				Email:       gofakeit.Email(),
				Password:    &strongPassword,
			},
			errorMsg: graphapi.ErrPermissionDenied.Error(),
		},
		// TODO: These will all have no-auth failures
		// until a policy is added to add service user concepts
		// users should generally be created via register or invite, and not
		// the create user graph api
		// {
		// 	name: "happy path user",
		// 	userInput: datumclient.CreateUserInput{
		// 		FirstName:   gofakeit.FirstName(),
		// 		LastName:    gofakeit.LastName(),
		// 		DisplayName: &displayName,
		// 		Email:       gofakeit.Email(),
		// 		Password:    &strongPassword,
		// 	},
		// 	errorMsg: "",
		// },
		// {
		// 	name: "no email",
		// 	userInput: datumclient.CreateUserInput{
		// 		FirstName:   gofakeit.FirstName(),
		// 		LastName:    gofakeit.LastName(),
		// 		DisplayName: &displayName,
		// 		Email:       "",
		// 	},
		// 	errorMsg: "mail: no address",
		// },
		// {
		// 	name: "no first name",
		// 	userInput: datumclient.CreateUserInput{
		// 		FirstName:   "",
		// 		LastName:    gofakeit.LastName(),
		// 		DisplayName: &displayName,
		// 		Email:       gofakeit.Email(),
		// 	},
		// 	errorMsg: "value is less than the required length",
		// },
		// {
		// 	name: "no last name",
		// 	userInput: datumclient.CreateUserInput{
		// 		FirstName:   gofakeit.FirstName(),
		// 		LastName:    "",
		// 		DisplayName: &displayName,
		// 		Email:       gofakeit.Email(),
		// 	},
		// 	errorMsg: "value is less than the required length",
		// },
		// {
		// 	name: "no display name, should default to email",
		// 	userInput: datumclient.CreateUserInput{
		// 		FirstName: gofakeit.FirstName(),
		// 		LastName:  gofakeit.LastName(),
		// 		Email:     gofakeit.Email(),
		// 	},
		// 	errorMsg: "",
		// },
		// {
		// 	name: "weak password",
		// 	userInput: datumclient.CreateUserInput{
		// 		FirstName: gofakeit.FirstName(),
		// 		LastName:  gofakeit.LastName(),
		// 		Email:     gofakeit.Email(),
		// 		Password:  &weakPassword,
		// 	},
		// 	errorMsg: auth.ErrPasswordTooWeak.Error(),
		// },
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			if tc.errorMsg == "" {
				// mock writes to create personal org membership
				mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)
			}
			resp, err := authClient.gc.CreateUser(reqCtx, tc.userInput)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateUser.User)

			// Make sure provided values match
			assert.Equal(t, tc.userInput.FirstName, resp.CreateUser.User.FirstName)
			assert.Equal(t, tc.userInput.LastName, resp.CreateUser.User.LastName)
			assert.Equal(t, tc.userInput.Email, resp.CreateUser.User.Email)

			// display name defaults to email if not provided
			if tc.userInput.DisplayName == nil {
				assert.Equal(t, tc.userInput.Email, resp.CreateUser.User.DisplayName)
			} else {
				assert.Equal(t, *tc.userInput.DisplayName, resp.CreateUser.User.DisplayName)
			}

			// ensure a user setting was created
			assert.NotNil(t, resp.CreateUser.User.Setting)

			// ensure personal org is created
			personalOrg := true
			whereInput := &datumclient.OrganizationWhereInput{
				PersonalOrg: &personalOrg,
			}

			// TODO: update to pull by user once https://github.com/datumforge/datum/issues/293 is complete
			orgs, err := authClient.gc.OrganizationsWhere(reqCtx, whereInput)
			require.NoError(t, err)

			orgCreated := false
			for _, o := range orgs.Organizations.GetEdges() {
				if *o.Node.Description == fmt.Sprintf("Personal Organization - %s %s", resp.CreateUser.User.FirstName, resp.CreateUser.User.LastName) {
					orgCreated = true
				}
			}

			assert.True(t, orgCreated, "personal org expected to be created")
		})
	}
}

func TestMutation_UpdateUser(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	firstNameUpdate := gofakeit.FirstName()
	lastNameUpdate := gofakeit.LastName()
	emailUpdate := gofakeit.Email()
	displayNameUpdate := gofakeit.LetterN(40)
	nameUpdateLong := gofakeit.LetterN(200)

	user := (&UserBuilder{}).MustNew(ctx)

	// setup valid user context
	userCtx, err := auth.NewTestContextWithValidUser(user.ID)
	if err != nil {
		t.Fatal()
	}

	reqCtx := context.WithValue(userCtx.Request().Context(), echocontext.EchoContextKey, userCtx)

	userCtx.SetRequest(ec.Request().WithContext(reqCtx))

	weakPassword := "notsecure"

	testCases := []struct {
		name        string
		updateInput datumclient.UpdateUserInput
		expectedRes datumclient.UpdateUser_UpdateUser_User
		errorMsg    string
	}{
		{
			name: "update first name and password, happy path",
			updateInput: datumclient.UpdateUserInput{
				FirstName: &firstNameUpdate,
			},
			expectedRes: datumclient.UpdateUser_UpdateUser_User{
				ID:          user.ID,
				FirstName:   firstNameUpdate,
				LastName:    user.LastName,
				DisplayName: user.DisplayName,
				Email:       user.Email,
			},
		},
		{
			name: "update last name, happy path",
			updateInput: datumclient.UpdateUserInput{
				LastName: &lastNameUpdate,
			},
			expectedRes: datumclient.UpdateUser_UpdateUser_User{
				ID:          user.ID,
				FirstName:   firstNameUpdate, // this would have been updated on the prior test
				LastName:    lastNameUpdate,
				DisplayName: user.DisplayName,
				Email:       user.Email,
			},
		},
		{
			name: "update email, happy path",
			updateInput: datumclient.UpdateUserInput{
				Email: &emailUpdate,
			},
			expectedRes: datumclient.UpdateUser_UpdateUser_User{
				ID:          user.ID,
				FirstName:   firstNameUpdate,
				LastName:    lastNameUpdate, // this would have been updated on the prior test
				DisplayName: user.DisplayName,
				Email:       emailUpdate,
			},
		},
		{
			name: "update display name, happy path",
			updateInput: datumclient.UpdateUserInput{
				DisplayName: &displayNameUpdate,
			},
			expectedRes: datumclient.UpdateUser_UpdateUser_User{
				ID:          user.ID,
				FirstName:   firstNameUpdate,
				LastName:    lastNameUpdate,
				DisplayName: displayNameUpdate,
				Email:       emailUpdate, // this would have been updated on the prior test
			},
		},
		{
			name: "update name, too long",
			updateInput: datumclient.UpdateUserInput{
				FirstName: &nameUpdateLong,
			},
			errorMsg: "value is greater than the required length",
		},
		{
			name: "update with weak password",
			updateInput: datumclient.UpdateUserInput{
				Password: &weakPassword,
			},
			errorMsg: auth.ErrPasswordTooWeak.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			// update user
			resp, err := authClient.gc.UpdateUser(reqCtx, user.ID, tc.updateInput)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdateUser.User)

			// Make sure provided values match
			updatedUser := resp.GetUpdateUser().User
			assert.Equal(t, tc.expectedRes.FirstName, updatedUser.FirstName)
			assert.Equal(t, tc.expectedRes.LastName, updatedUser.LastName)
			assert.Equal(t, tc.expectedRes.DisplayName, updatedUser.DisplayName)
			assert.Equal(t, tc.expectedRes.Email, updatedUser.Email)
		})
	}
}

func TestMutation_DeleteUser(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// Setup echo context
	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	// bypass auth on object creation
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	user := (&UserBuilder{}).MustNew(ctx)

	userSetting, err := user.Setting(ctx)
	require.NoError(t, err)

	orgs, err := user.Organizations(ctx)
	require.NoError(t, err)
	require.Len(t, orgs, 1)

	personalOrgID := orgs[0].ID

	listObjects := []string{fmt.Sprintf("organization:%s", personalOrgID)}

	// setup valid user context
	userCtx, err := auth.NewTestContextWithValidUser(user.ID)
	if err != nil {
		t.Fatal()
	}

	reqCtx := context.WithValue(userCtx.Request().Context(), echocontext.EchoContextKey, userCtx)

	userCtx.SetRequest(ec.Request().WithContext(reqCtx))

	testCases := []struct {
		name     string
		userID   string
		errorMsg string
	}{
		{
			name:   "delete user, happy path",
			userID: user.ID,
		},
		{
			name:     "delete user, not found",
			userID:   "tacos-tuesday",
			errorMsg: "not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			// mock check calls
			if tc.errorMsg == "" {
				mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
				mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
				mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
				mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)

				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
				mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)

				mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)
			}

			// delete user
			resp, err := authClient.gc.DeleteUser(reqCtx, tc.userID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.DeleteUser.DeletedID)

			// make sure the personal org is deleted
			org, err := authClient.gc.GetOrganizationByID(reqCtx, personalOrgID)
			require.Nil(t, org)
			require.Error(t, err)
			assert.ErrorContains(t, err, "not found")

			// make sure the deletedID matches the ID we wanted to delete
			assert.Equal(t, tc.userID, resp.DeleteUser.DeletedID)

			// make sure the user setting is deleted
			out, err := authClient.gc.GetUserSettingByID(reqCtx, userSetting.ID)
			require.Nil(t, out)
			require.Error(t, err)
			assert.ErrorContains(t, err, "not found")
		})
	}
}

func TestMutation_UserCascadeDelete(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	usr := (&UserBuilder{}).MustNew(ctx)

	token1 := (&PersonalAccessTokenBuilder{OwnerID: usr.ID}).MustNew(ctx)

	// setup valid user context
	userCtx, err := auth.NewTestContextWithValidUser(usr.ID)
	if err != nil {
		t.Fatal()
	}

	reqCtx := context.WithValue(userCtx.Request().Context(), echocontext.EchoContextKey, userCtx)

	userCtx.SetRequest(ec.Request().WithContext(reqCtx))

	// mock checks
	mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
	mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)

	orgs, err := usr.OrgMemberships(ctx)
	require.NoError(t, err)
	require.Len(t, orgs, 1)

	listObjects := []string{fmt.Sprintf("organization:%s", orgs[0].OrgID)}

	mockCheckAny(authClient.mockCtrl, authClient.mc, reqCtx, true)
	mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)
	mockListAny(authClient.mockCtrl, authClient.mc, reqCtx, listObjects)

	// mock writes to clean up personal org
	mockWriteAny(authClient.mockCtrl, authClient.mc, reqCtx, nil)

	// delete user
	resp, err := authClient.gc.DeleteUser(reqCtx, usr.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.DeleteUser.DeletedID)

	// make sure the deletedID matches the ID we wanted to delete
	assert.Equal(t, usr.ID, resp.DeleteUser.DeletedID)

	o, err := authClient.gc.GetUserByID(reqCtx, usr.ID)

	require.Nil(t, o)
	require.Error(t, err)
	assert.ErrorContains(t, err, "not found")

	g, err := authClient.gc.GetPersonalAccessTokenByID(reqCtx, token1.ID)
	require.Error(t, err)

	require.Nil(t, g)
	assert.ErrorContains(t, err, "not found")

	ctx = mixin.SkipSoftDelete(reqCtx)

	// skip checks because tuples will be deleted at this point
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	o, err = authClient.gc.GetUserByID(ctx, usr.ID)
	require.NoError(t, err)

	require.Equal(t, o.User.ID, usr.ID)

	// Bypass auth check to get owner of access token
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	g, err = authClient.gc.GetPersonalAccessTokenByID(ctx, token1.ID)
	require.NoError(t, err)

	require.Equal(t, g.PersonalAccessToken.ID, token1.ID)
}

func TestMutation_SoftDeleteUniqueIndex(t *testing.T) {
	// setup entdb with authz
	authClient := setupAuthEntDB(t)
	defer authClient.entDB.Close()

	// Setup echo context
	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	input := datumclient.CreateUserInput{
		FirstName: "Abraxos",
		LastName:  "Funk",
		Email:     "abraxos@datum.net",
	}

	// skip auth checks because user creation is not generally allowed via a direct mutation
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	// mocks
	// the object ID doesn't matter here because we end up needing to bypass auth checks
	listObjects := []string{fmt.Sprintf("organization:%s", "test")}

	// write tuples for personal org on create, delete, and create again
	mockWriteAny(authClient.mockCtrl, authClient.mc, ctx, nil)
	mockWriteAny(authClient.mockCtrl, authClient.mc, ctx, nil)
	mockWriteAny(authClient.mockCtrl, authClient.mc, ctx, nil)

	// check access for requests
	mockCheckAny(authClient.mockCtrl, authClient.mc, ctx, true)
	mockListAny(authClient.mockCtrl, authClient.mc, ctx, listObjects)
	mockListAny(authClient.mockCtrl, authClient.mc, ctx, listObjects)

	resp, err := authClient.gc.CreateUser(ctx, input)
	require.NoError(t, err)

	// should fail on unique
	_, err = authClient.gc.CreateUser(ctx, input)
	require.Error(t, err)
	assert.ErrorContains(t, err, "UNIQUE")

	// setup valid user context
	userCtx, err := auth.NewTestContextWithValidUser(resp.CreateUser.User.ID)
	if err != nil {
		t.Fatal()
	}

	reqCtx := context.WithValue(userCtx.Request().Context(), echocontext.EchoContextKey, userCtx)

	userCtx.SetRequest(ec.Request().WithContext(reqCtx))

	// delete user
	_, err = authClient.gc.DeleteUser(userCtx.Request().Context(), resp.CreateUser.User.ID)
	require.NoError(t, err)

	// skip checks because tuples will be deleted at this point
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	o, err := authClient.gc.GetUserByID(ctx, resp.CreateUser.User.ID)

	require.Nil(t, o)
	require.Error(t, err)
	assert.ErrorContains(t, err, "not found")

	// Ensure user is soft deleted
	ctx = mixin.SkipSoftDelete(userCtx.Request().Context())

	o, err = authClient.gc.GetUserByID(ctx, resp.CreateUser.User.ID)
	require.NoError(t, err)

	require.Equal(t, o.User.ID, resp.CreateUser.User.ID)

	// create the user again, this should work because we should ignore soft deleted
	// records on unique email
	// skip auth checks because user creation is not generally allowed via a direct mutation
	ctx = privacy.DecisionContext(ctx, privacy.Allow)
	resp, err = authClient.gc.CreateUser(ctx, input)
	require.NoError(t, err)
	assert.Equal(t, input.Email, resp.CreateUser.User.Email)
}
