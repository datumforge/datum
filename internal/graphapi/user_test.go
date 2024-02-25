package graphapi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/datumforge/entx"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/graphapi"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	auth "github.com/datumforge/datum/pkg/auth"
)

func TestQueryUser(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	user1 := (&UserBuilder{client: client}).MustNew(ctx, t)
	user2 := (&UserBuilder{client: client}).MustNew(ctx, t)

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
			defer mock_fga.ClearMocks(client.fga)

			if tc.errorMsg == "" {
				// placeholder until authz is enforced on org members
				// at which point this needs to be the correct organization id
				listObjects := []string{fmt.Sprintf("organization:%s", "test")}
				mock_fga.ListAny(t, client.fga, listObjects)
			}

			resp, err := client.datum.GetUserByID(reqCtx, tc.queryID)

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

	(&UserCleanup{client: client, UserID: user1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: client, UserID: user2.ID}).MustDelete(reqCtx, t)
}

func TestQueryUsers(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	user1 := (&UserBuilder{client: client}).MustNew(reqCtx, t)
	user2 := (&UserBuilder{client: client}).MustNew(reqCtx, t)

	t.Run("Get Users", func(t *testing.T) {
		resp, err := client.datum.GetAllUsers(reqCtx)

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

		resp, err = client.datum.GetAllUsers(reqCtx)

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
	})
}

func TestMutationCreateUserNoAuth(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	// Bypass auth checks to ensure input checks for now
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	weakPassword := "notsecure"
	strongPassword := "my&supers3cr3tpassw0rd!"

	email := gofakeit.Email()

	testCases := []struct {
		name      string
		userInput datumclient.CreateUserInput
		errorMsg  string
	}{
		{
			name: "happy path user",
			userInput: datumclient.CreateUserInput{
				FirstName:    gofakeit.FirstName(),
				LastName:     gofakeit.LastName(),
				DisplayName:  gofakeit.LetterN(50),
				Email:        email,
				AuthProvider: &enums.Credentials,
				Password:     &strongPassword,
			},
			errorMsg: "",
		},
		{
			name: "same email, same auth provider",
			userInput: datumclient.CreateUserInput{
				FirstName:    gofakeit.FirstName(),
				LastName:     gofakeit.LastName(),
				DisplayName:  gofakeit.LetterN(50),
				Email:        email,
				AuthProvider: &enums.Credentials,
				Password:     &strongPassword,
			},
			errorMsg: "constraint failed",
		},
		{
			name: "same email, different auth provider",
			userInput: datumclient.CreateUserInput{
				FirstName:    gofakeit.FirstName(),
				LastName:     gofakeit.LastName(),
				DisplayName:  gofakeit.LetterN(50),
				Email:        email,
				AuthProvider: &enums.GitHub,
			},
			errorMsg: "",
		},
		{
			name: "no email",
			userInput: datumclient.CreateUserInput{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				DisplayName: gofakeit.LetterN(50),
				Email:       "",
			},
			errorMsg: "mail: no address",
		},
		{
			name: "no first name",
			userInput: datumclient.CreateUserInput{
				FirstName:   "",
				LastName:    gofakeit.LastName(),
				DisplayName: gofakeit.LetterN(50),
				Email:       gofakeit.Email(),
			},
			errorMsg: "value is less than the required length",
		},
		{
			name: "no last name",
			userInput: datumclient.CreateUserInput{
				FirstName:   gofakeit.FirstName(),
				LastName:    "",
				DisplayName: gofakeit.LetterN(50),
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
			defer mock_fga.ClearMocks(client.fga)

			if tc.errorMsg == "" {
				// mock writes to create personal org membership
				mock_fga.WriteAny(t, client.fga)
			}

			resp, err := client.datum.CreateUser(ctx, tc.userInput)

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

			if tc.userInput.AuthProvider != nil {
				assert.Equal(t, tc.userInput.AuthProvider, &resp.CreateUser.User.AuthProvider)
			} else {
				// default is credentials if not provided
				assert.Equal(t, enums.Credentials, resp.CreateUser.User.AuthProvider)
			}
			// display name defaults to email if not provided
			if tc.userInput.DisplayName == "" {
				assert.Equal(t, tc.userInput.Email, resp.CreateUser.User.DisplayName)
			} else {
				assert.Equal(t, tc.userInput.DisplayName, resp.CreateUser.User.DisplayName)
			}

			// subject should always be set
			assert.Equal(t, resp.CreateUser.User.ID, *resp.CreateUser.User.Sub)

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
			listObjects := []string{fmt.Sprintf("organization:%s", orgs[0].OrganizationID)}
			mock_fga.ListAny(t, client.fga, listObjects)

			// Bypass auth checks to ensure input checks for now
			reqCtx = privacy.DecisionContext(reqCtx, privacy.Allow)

			personalOrg, err := client.datum.GetOrganizationByID(reqCtx, orgs[0].OrganizationID)
			require.NoError(t, err)

			assert.True(t, personalOrg.Organization.PersonalOrg)
			// make sure there is only one user
			require.Len(t, personalOrg.Organization.Members, 1)
			// make sure user was added as owner
			assert.Equal(t, personalOrg.Organization.Members[0].Role.String(), "OWNER")
		})
	}
}

func TestMutationCreateUser(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

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
				DisplayName: gofakeit.LetterN(50),
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
		// 		DisplayName: gofakeit.LetterN(50),
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
		// 		DisplayName: gofakeit.LetterN(50),
		// 		Email:       "",
		// 	},
		// 	errorMsg: "mail: no address",
		// },
		// {
		// 	name: "no first name",
		// 	userInput: datumclient.CreateUserInput{
		// 		FirstName:   "",
		// 		LastName:    gofakeit.LastName(),
		// 		DisplayName: gofakeit.LetterN(50),
		// 		Email:       gofakeit.Email(),
		// 	},
		// 	errorMsg: "value is less than the required length",
		// },
		// {
		// 	name: "no last name",
		// 	userInput: datumclient.CreateUserInput{
		// 		FirstName:   gofakeit.FirstName(),
		// 		LastName:    "",
		// 		DisplayName: gofakeit.LetterN(50),
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
			defer mock_fga.ClearMocks(client.fga)

			if tc.errorMsg == "" {
				// mock writes to create personal org membership
				mock_fga.WriteAny(t, client.fga)
			}

			resp, err := client.datum.CreateUser(reqCtx, tc.userInput)

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
			if tc.userInput.DisplayName == "" {
				assert.Equal(t, tc.userInput.Email, resp.CreateUser.User.DisplayName)
			} else {
				assert.Equal(t, tc.userInput.DisplayName, resp.CreateUser.User.DisplayName)
			}

			// ensure a user setting was created
			assert.NotNil(t, resp.CreateUser.User.Setting)

			// ensure personal org is created
			// default org will always be the personal org when the user is first created
			personalOrgID := *resp.CreateUser.User.Setting.DefaultOrg

			org, err := client.datum.GetOrganizationByID(reqCtx, personalOrgID, nil)
			require.NoError(t, err)
			assert.Equal(t, personalOrgID, org.Organization.ID)
			assert.True(t, org.Organization.PersonalOrg)
		})
	}
}

func TestMutationUpdateUser(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	firstNameUpdate := gofakeit.FirstName()
	lastNameUpdate := gofakeit.LastName()
	emailUpdate := gofakeit.Email()
	displayNameUpdate := gofakeit.LetterN(40)
	nameUpdateLong := gofakeit.LetterN(200)

	user := (&UserBuilder{client: client}).MustNew(ctx, t)

	// setup valid user context
	reqCtx, err := userContextWithID(user.ID)
	require.NoError(t, err)

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
			defer mock_fga.ClearMocks(client.fga)

			// checks for member tables
			if tc.errorMsg == "" {
				mock_fga.CheckAny(t, client.fga, true)
			}

			// update user
			resp, err := client.datum.UpdateUser(reqCtx, user.ID, tc.updateInput)

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

func TestMutationDeleteUser(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// Setup echo context
	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	// bypass auth on object creation
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	user := (&UserBuilder{client: client}).MustNew(ctx, t)

	userSetting, err := user.Setting(ctx)
	require.NoError(t, err)

	// setup valid user context
	reqCtx, err := userContextWithID(user.ID)
	require.NoError(t, err)

	// bypass auth
	ctx = privacy.DecisionContext(reqCtx, privacy.Allow)
	userSettings, err := user.Setting(ctx)
	require.NoError(t, err)

	// personal org will be the default org when the user is created
	personalOrgID := userSettings.DefaultOrg

	listObjects := []string{fmt.Sprintf("organization:%s", personalOrgID)}

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
			defer mock_fga.ClearMocks(client.fga)

			// mock check calls
			if tc.errorMsg == "" {
				mock_fga.CheckAny(t, client.fga, true)

				mock_fga.ListAny(t, client.fga, listObjects)

				mock_fga.WriteAny(t, client.fga)
			}

			// delete user
			resp, err := client.datum.DeleteUser(reqCtx, tc.userID)

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
			org, err := client.datum.GetOrganizationByID(reqCtx, personalOrgID)
			require.Nil(t, org)
			require.Error(t, err)
			assert.ErrorContains(t, err, "not found")

			// make sure the deletedID matches the ID we wanted to delete
			assert.Equal(t, tc.userID, resp.DeleteUser.DeletedID)

			// make sure the user setting is deleted
			out, err := client.datum.GetUserSettingByID(reqCtx, userSetting.ID)
			require.Nil(t, out)
			require.Error(t, err)
			assert.ErrorContains(t, err, "not found")
		})
	}
}

func TestMutationUserCascadeDelete(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	ec := echocontext.NewTestEchoContext()

	ctx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(ctx))

	user := (&UserBuilder{client: client}).MustNew(ctx, t)

	reqCtx, err := userContextWithID(user.ID)
	require.NoError(t, err)

	token := (&PersonalAccessTokenBuilder{client: client, OwnerID: user.ID}).MustNew(reqCtx, t)

	userSettings, err := user.Setting(ctx)
	require.NoError(t, err)

	listObjects := []string{fmt.Sprintf("organization:%s", userSettings.DefaultOrg)}

	// mock checks
	mock_fga.CheckAny(t, client.fga, true)
	mock_fga.ListAny(t, client.fga, listObjects)
	// mock writes to clean up personal org
	mock_fga.WriteAny(t, client.fga)

	// delete user
	resp, err := client.datum.DeleteUser(reqCtx, user.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.DeleteUser.DeletedID)

	// make sure the deletedID matches the ID we wanted to delete
	assert.Equal(t, user.ID, resp.DeleteUser.DeletedID)

	o, err := client.datum.GetUserByID(reqCtx, user.ID)

	require.Nil(t, o)
	require.Error(t, err)
	assert.ErrorContains(t, err, "not found")

	g, err := client.datum.GetPersonalAccessTokenByID(reqCtx, token.ID)
	require.Error(t, err)

	require.Nil(t, g)
	assert.ErrorContains(t, err, "not found")

	ctx = entx.SkipSoftDelete(reqCtx)

	// skip checks because tuples will be deleted at this point
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	o, err = client.datum.GetUserByID(ctx, user.ID)
	require.NoError(t, err)

	require.Equal(t, o.User.ID, user.ID)

	// Bypass auth check to get owner of access token
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	g, err = client.datum.GetPersonalAccessTokenByID(ctx, token.ID)
	require.NoError(t, err)

	require.Equal(t, g.PersonalAccessToken.ID, token.ID)
}

func TestMutationSoftDeleteUniqueIndex(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

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
	mock_fga.WriteAny(t, client.fga)

	// check access for requests
	mock_fga.CheckAny(t, client.fga, true)
	mock_fga.ListAny(t, client.fga, listObjects)

	resp, err := client.datum.CreateUser(ctx, input)
	require.NoError(t, err)

	// should fail on unique
	_, err = client.datum.CreateUser(ctx, input)
	require.Error(t, err)
	assert.ErrorContains(t, err, "constraint failed")

	// setup valid user context
	userCtx, err := auth.NewTestContextWithValidUser(resp.CreateUser.User.ID)
	if err != nil {
		t.Fatal()
	}

	reqCtx := context.WithValue(userCtx.Request().Context(), echocontext.EchoContextKey, userCtx)

	userCtx.SetRequest(ec.Request().WithContext(reqCtx))

	// delete user
	_, err = client.datum.DeleteUser(userCtx.Request().Context(), resp.CreateUser.User.ID)
	require.NoError(t, err)

	// skip checks because tuples will be deleted at this point
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	o, err := client.datum.GetUserByID(ctx, resp.CreateUser.User.ID)

	require.Nil(t, o)
	require.Error(t, err)
	assert.ErrorContains(t, err, "not found")

	// Ensure user is soft deleted
	ctx = entx.SkipSoftDelete(userCtx.Request().Context())

	o, err = client.datum.GetUserByID(ctx, resp.CreateUser.User.ID)
	require.NoError(t, err)

	require.Equal(t, o.User.ID, resp.CreateUser.User.ID)

	// create the user again, this should work because we should ignore soft deleted
	// records on unique email
	// skip auth checks because user creation is not generally allowed via a direct mutation
	ctx = privacy.DecisionContext(ctx, privacy.Allow)
	resp, err = client.datum.CreateUser(ctx, input)
	require.NoError(t, err)
	assert.Equal(t, input.Email, resp.CreateUser.User.Email)
}
