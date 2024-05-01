package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	echo "github.com/datumforge/echox"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/rShetty/asyncwait"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/emails/mock"
)

func (suite *HandlerTestSuite) TestRegisterHandler() {
	t := suite.T()

	// add handler
	suite.e.POST("register", suite.h.RegisterHandler)

	var bonkers = "b!a!n!a!n!a!s!"

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
			password:       bonkers,
			emailExpected:  true,
			expectedStatus: http.StatusCreated,
		},
		{
			name:               "invalid email",
			email:              "bananas.net",
			firstName:          "Princess",
			lastName:           "Fiona",
			password:           bonkers,
			emailExpected:      false,
			expectedErrMessage: "email was invalid",
			expectedStatus:     http.StatusBadRequest,
		},
		{
			name:               "missing email",
			firstName:          "Princess",
			lastName:           "Fiona",
			password:           bonkers,
			emailExpected:      false,
			expectedErrMessage: "missing required field: email",
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
			defer mock_fga.ClearMocks(suite.fga)

			sent := time.Now()

			mock.ResetEmailMock()

			// setup mock authz writes
			if tc.expectedErrMessage == "" {
				mock_fga.WriteAny(t, suite.fga)
			}

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
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			suite.e.ServeHTTP(recorder, req)

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
				ec, err := auth.NewTestEchoContextWithValidUser(out.ID)
				require.NoError(t, err)

				ctx := ec.Request().Context()

				// get the user and make sure things were created as expected
				user, err := suite.db.User.Get(ctx, out.ID)
				require.NoError(t, err)

				// setup echo context
				ctx = context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

				// Bypass auth check because user is not authenticated before verified
				ctx = privacy.DecisionContext(ctx, privacy.Allow)

				// mocks to check for org access
				listObjects := []string{"organization:test"}
				mock_fga.ListOnce(t, suite.fga, listObjects, nil)

				// make sure user is an owner of their personal org
				orgMemberships, err := user.OrgMemberships(ctx)
				require.NoError(t, err)
				require.Len(t, orgMemberships, 1)
				assert.Equal(t, orgMemberships[0].Role, enums.RoleOwner)

				// delete user
				suite.db.User.DeleteOneID(out.ID).ExecX(ctx)
			} else {
				assert.Contains(t, out.Error, tc.expectedErrMessage)
			}

			// Test that one verify email was sent to each user
			messages := []*mock.EmailMetadata{
				{
					To:        tc.email,
					From:      "mitb@datum.net",
					Subject:   emails.VerifyEmailRE,
					Timestamp: sent,
				},
			}

			// wait for messages
			predicate := func() bool {
				return suite.h.TaskMan.GetQueueLength() == 0
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
