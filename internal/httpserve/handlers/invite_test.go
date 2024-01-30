package handlers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	mock_fga "github.com/datumforge/fgax/mockery"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/rShetty/asyncwait"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/emails/mock"
)

func TestOrgInviteAcceptHandler(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// add handler
	client.e.POST("invite", client.h.OrganizationInviteAccept)

	// bypass auth
	ctx := context.Background()
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	mock_fga.WriteAny(t, client.fga)

	// setup test data
	user := client.db.User.Create().
		SetEmail("rocket@datum.net").
		SetFirstName("Rocket").
		SetLastName("Racoon").
		SaveX(ctx)

	ec, err := auth.NewTestContextWithValidUser(user.ID)
	require.NoError(t, err)

	newCtx := ec.Request().Context()
	newCtx = privacy.DecisionContext(newCtx, privacy.Allow)

	reqCtx := context.WithValue(newCtx, echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	org := client.db.Organization.Create().
		SetName("avengers").
		SaveX(reqCtx)

	testCases := []struct {
		name          string
		email         string
		firstName     string
		lastName      string
		password      string
		tokenSet      bool
		emailExpected bool
		wantErr       bool
		errMsg        string
	}{
		{
			name:          "happy path",
			email:         "groot@datum.net",
			firstName:     "Groot",
			lastName:      "JustGroot",
			password:      "IAmGr00t!",
			emailExpected: true,
			tokenSet:      true,
		},
		{
			name:      "missing token",
			email:     "drax@datum.net",
			firstName: "Drax",
			lastName:  "TheDestroyer",
			password:  "IllD0YoU1B3tt3r",
			tokenSet:  false,
			wantErr:   true,
			errMsg:    "token is required",
		},
		{
			name:      "missing password",
			email:     "gamora@datum.net",
			firstName: "Gamora",
			lastName:  "Zen Whoberi Ben Titan",
			tokenSet:  true,
			wantErr:   true,
			errMsg:    "missing required field: password",
		},
		{
			name:      "missing last name",
			email:     "yondu@datum.net",
			firstName: "Yondu",
			password:  "RememberB0y!",
			tokenSet:  true,
			wantErr:   true,
			errMsg:    "missing required field: last name",
		},
		{
			name:     "missing first name",
			email:    "thanos@datum.net",
			lastName: "Thanos",
			password: "RealityISOft3n",
			tokenSet: true,
			wantErr:  true,
			errMsg:   "missing required field: first name",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(client.fga)
			sent := time.Now()
			mock.ResetEmailMock()

			// mock auth
			mock_fga.ListAny(t, client.fga, []string{fmt.Sprintf("organization:%s", org.ID)})

			acceptInviteJSON := handlers.InviteRequest{
				FirstName: tc.firstName,
				LastName:  tc.lastName,
				Password:  tc.password,
			}

			invite := client.db.Invite.Create().
				SetOwnerID(org.ID).
				SetRecipient(tc.email).SaveX(reqCtx)

			body, err := json.Marshal(acceptInviteJSON)
			if err != nil {
				require.NoError(t, err)
			}

			target := "/invite"
			if tc.tokenSet {
				target = fmt.Sprintf("/invite?token=%s", invite.Token)
			}

			req := httptest.NewRequest(http.MethodPost, target, strings.NewReader(string(body)))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			client.e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.InviteReply

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			if tc.wantErr {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				assert.Equal(t, tc.errMsg, out.Message)

				return
			}

			assert.Equal(t, http.StatusCreated, recorder.Code)
			assert.Equal(t, org.ID, out.JoinedOrgID)
			assert.Equal(t, tc.email, out.Email)

			// Test that one verify email was sent to each user
			// one for invite, one for accepted
			messages := []*mock.EmailMetadata{
				{
					To:        tc.email,
					From:      "mitb@datum.net",
					Subject:   fmt.Sprintf(emails.InviteRE, user.FirstName),
					Timestamp: sent,
				},
				{
					To:        tc.email,
					From:      "mitb@datum.net",
					Subject:   emails.InviteBeenAccepted,
					Timestamp: sent,
				},
			}

			// wait for messages
			predicate := func() bool {
				return client.h.TaskMan.GetQueueLength() == 0
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
