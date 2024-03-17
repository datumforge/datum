package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
)

func (suite *HandlerTestSuite) TestUnsubscribeHandler() {
	t := suite.T()

	// add handler
	suite.e.GET("unsubscribe", suite.h.UnsubscribeHandler)

	ec := echocontext.NewTestEchoContext().Request().Context()

	email := gofakeit.Email()

	user := handlers.User{
		Email: email,
	}

	// create token
	if err := user.CreateVerificationToken(); err != nil {
		require.NoError(t, err)
	}

	ttl, err := time.Parse(time.RFC3339Nano, user.EmailVerificationExpires.String)
	if err != nil {
		require.NoError(t, err)
	}

	// set privacy allow in order to allow the creation of the users without
	// authentication in the tests
	ctx := privacy.DecisionContext(ec, privacy.Allow)

	// store token in db
	et := suite.db.Subscriber.Create().
		SetToken(user.EmailVerificationToken.String).
		SetEmail(user.Email).
		SetSecret(user.EmailVerificationSecret).
		SetTTL(ttl).
		SaveX(ctx)

	testCases := []struct {
		name               string
		token              string
		email              string
		expectedErrMessage string
		expectedStatus     int
	}{
		{
			name:               "happy path, unsubscriber exists",
			token:              et.Token,
			email:              email,
			expectedStatus:     http.StatusOK,
			expectedErrMessage: "",
		},
		{
			name:               "invalid token",
			token:              "invalid",
			email:              email,
			expectedStatus:     http.StatusBadRequest,
			expectedErrMessage: "subscriber not found",
		},
		{
			name:               "missing email",
			token:              et.Token,
			email:              "",
			expectedStatus:     http.StatusBadRequest,
			expectedErrMessage: "email is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			target := fmt.Sprintf("/unsubscribe?email=%s", tc.email)

			req := httptest.NewRequest(http.MethodGet, target, nil)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			suite.e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.SubscribeReply

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus == http.StatusOK {
				assert.NotEmpty(t, out.Message)
			} else {
				assert.Contains(t, out.Error, tc.expectedErrMessage)
			}
		})
	}
}
