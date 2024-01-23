package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/rShetty/asyncwait"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	mock_client "github.com/datumforge/datum/internal/fga/mocks"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/emails/mock"
)

func TestRegisterHandler(t *testing.T) {
	// setup mock controller
	mockCtrl := gomock.NewController(t)

	mc := mock_client.NewMockSdkClient(mockCtrl)

	// setup entdb with authz
	entClient := setupAuthEntDB(t, mockCtrl, mc)
	defer entClient.Close()

	h := handlerSetup(t, entClient)

	testCases := []struct {
		name               string
		email              string
		firstName          string
		lastName           string
		password           string
		emailExpected      bool
		expectedErrMessage string
		expectedStatus     int
		from               string
	}{
		{
			name:           "happy path",
			email:          "bananas@datum.net",
			firstName:      "Princess",
			lastName:       "Fiona",
			password:       "b!a!n!a!n!a!s!",
			emailExpected:  true,
			expectedStatus: http.StatusCreated,
		},
		{
			name:               "invalid email",
			email:              "bananas.net",
			firstName:          "Princess",
			lastName:           "Fiona",
			password:           "b!a!n!a!n!a!s!",
			emailExpected:      false,
			expectedErrMessage: "email was invalid",
			expectedStatus:     http.StatusBadRequest,
		},
		{
			name:               "missing email",
			firstName:          "Princess",
			lastName:           "Fiona",
			password:           "b!a!n!a!n!a!s!",
			emailExpected:      false,
			expectedErrMessage: "missing required field: email",
			expectedStatus:     http.StatusBadRequest,
		},
		{
			name:               "missing first name",
			email:              "tacos@datum.net",
			lastName:           "Fiona",
			password:           "b!a!n!a!n!a!s!",
			emailExpected:      false,
			expectedErrMessage: "missing required field: first name",
			expectedStatus:     http.StatusBadRequest,
		},
		{
			name:               "missing last name",
			email:              "waffles@datum.net",
			firstName:          "Princess",
			password:           "b!a!n!a!n!a!s!",
			emailExpected:      false,
			expectedErrMessage: "missing required field: last name",
			expectedStatus:     http.StatusBadRequest,
		},
		{
			name:               "bad password",
			email:              "pancakes@datum.net",
			firstName:          "Princess",
			lastName:           "Fiona",
			password:           "asfghjkl",
			emailExpected:      false,
			expectedErrMessage: "password is too weak",
			expectedStatus:     http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sent := time.Now()
			mock.ResetEmailMock()

			// create echo context with middleware
			e := setupEchoAuth(entClient)

			// setup mock authz writes
			if tc.expectedErrMessage == "" {
				mockWriteAny(mockCtrl, mc, context.Background(), nil)
			}

			e.POST("register", h.RegisterHandler)

			registerJSON := handlers.RegisterRequest{
				FirstName: tc.firstName,
				LastName:  tc.lastName,
				Email:     tc.email,
				Password:  tc.password,
			}

			body, err := json.Marshal(registerJSON)
			if err != nil {
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(string(body)))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.RegisterReply

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus == http.StatusCreated {
				assert.Equal(t, out.Email, tc.email)
				assert.NotEmpty(t, out.Message)
				assert.NotEmpty(t, out.ID)

				// setup context to get the data back
				ec := echocontext.NewTestEchoContext().Request().Context()
				ctx := privacy.DecisionContext(ec, privacy.Allow)

				// get the user and make sure things were created as expected
				user, err := EntClient.User.Get(ctx, out.ID)
				require.NoError(t, err)

				// make sure personal org was created
				orgs, err := user.Organizations(ctx)
				require.NoError(t, err)
				require.Len(t, orgs, 1)
				assert.True(t, orgs[0].PersonalOrg)

				// make sure user is an owner of their personal org
				orgMemberships, err := user.OrgMemberships(ctx)
				require.NoError(t, err)
				require.Len(t, orgs, 1)
				assert.Equal(t, orgMemberships[0].Role, enums.RoleOwner)

				// make sure the personal org was set as the user's default org
				userSettings, err := user.Setting(ctx)
				require.NoError(t, err)
				assert.Equal(t, userSettings.DefaultOrg, orgs[0].ID)

				// delete user
				EntClient.User.DeleteOneID(out.ID).ExecX(ctx)
			} else {
				assert.Contains(t, out.Message, tc.expectedErrMessage)
			}

			// Test that one verify email was sent to each user
			messages := []*mock.EmailMetadata{
				{
					To:        tc.email,
					From:      "no-reply@datum.net",
					Subject:   emails.VerifyEmailRE,
					Timestamp: sent,
				},
			}

			// wait for messages
			predicate := func() bool {
				return h.TaskMan.GetQueueLength() == 0
			}
			successful := asyncwait.NewAsyncWait(maxWaitInMillis, pollIntervalInMillis).Check(predicate)

			if successful != true {
				t.Errorf("max wait of email send")
			}

			if tc.emailExpected {
				mock.CheckEmails(t, messages)
			} else {
				mock.CheckEmails(t, nil)
			}
		})
	}
}
