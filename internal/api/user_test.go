package api_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/echox"
	ent "github.com/datumforge/datum/internal/ent/generated"
)

func TestQuery_User(t *testing.T) {
	// Setup Test Graph Client
	client := graphTestClient(EntClient)

	ec, err := echox.NewTestContextWithValidUser(subClaim)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	user1 := (&UserBuilder{}).MustNew(reqCtx)

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
			name:     "invalid-id",
			queryID:  "tacos-for-dinner",
			errorMsg: "user not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := client.GetUserByID(reqCtx, tc.queryID)

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
}

func TestQuery_Users(t *testing.T) {
	// Setup Test Graph Client
	client := graphTestClient(EntClient)

	ec, err := echox.NewTestContextWithValidUser(subClaim)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	user1 := (&UserBuilder{}).MustNew(reqCtx)
	user2 := (&UserBuilder{}).MustNew(reqCtx)

	t.Run("Get Users", func(t *testing.T) {
		resp, err := client.GetAllUsers(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Users.Edges)

		// make sure at least two users are returned
		assert.GreaterOrEqual(t, len(resp.Users.Edges), 2)

		user1Found := false
		user2Found := false
		for _, o := range resp.Users.Edges {
			if o.Node.ID == user1.ID {
				user1Found = true
			} else if o.Node.ID == user2.ID {
				user2Found = true
			}
		}

		// if one of the users isn't found, fail the test
		if !user1Found || !user2Found {
			t.Fail()
		}
	})
}

func TestMutation_CreateUser(t *testing.T) {
	// Setup Test Graph Client
	client := graphTestClient(EntClient)

	// Setup echo context
	ec, err := echox.NewTestContextWithValidUser(subClaim)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	displayName := gofakeit.LetterN(50)

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
			name: "no display name", // TODO: confirm we want to require display name for users
			userInput: datumclient.CreateUserInput{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				DisplayName: new(string),
				Email:       gofakeit.Email(),
			},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			resp, err := client.CreateUser(reqCtx, tc.userInput)

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
			assert.Equal(t, *tc.userInput.DisplayName, resp.CreateUser.User.DisplayName)

			// TODO: make sure an org was created, requires auth checks first
		})
	}
}

func TestMutation_UpdateUser(t *testing.T) {
	// Setup Test Graph Client
	client := graphTestClient(EntClient)

	// Setup echo context
	ec, err := echox.NewTestContextWithValidUser(subClaim)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	firstNameUpdate := gofakeit.FirstName()
	lastNameUpdate := gofakeit.LastName()
	emailUpdate := gofakeit.Email()
	displayNameUpdate := gofakeit.LetterN(40)
	nameUpdateLong := gofakeit.LetterN(200)

	user := (&UserBuilder{}).MustNew(reqCtx)

	testCases := []struct {
		name        string
		updateInput datumclient.UpdateUserInput
		expectedRes datumclient.UpdateUser_UpdateUser_User
		errorMsg    string
	}{
		{
			name: "update first name, happy path",
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
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			// update user
			resp, err := client.UpdateUser(reqCtx, user.ID, tc.updateInput)

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
	// Setup Test Graph Client
	client := graphTestClient(EntClient)

	// Setup echo context
	ec, err := echox.NewTestContextWithValidUser(subClaim)
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	reqCtx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(reqCtx))

	user := (&UserBuilder{}).MustNew(reqCtx)

	testCases := []struct {
		name     string
		userID   string
		errorMsg string
	}{
		{
			name:   "delete org, happy path",
			userID: user.ID,
		},
		{
			name:     "delete org, not found",
			userID:   "tacos-tuesday",
			errorMsg: "not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			// delete org
			resp, err := client.DeleteUser(reqCtx, tc.userID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.DeleteUser.DeletedID)

			// TODO: ensure personal org is also deleted when user is deleted

			// make sure the deletedID matches the ID we wanted to delete
			assert.Equal(t, tc.userID, resp.DeleteUser.DeletedID)
		})
	}
}
