package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	echo "github.com/datumforge/echox"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/httpserve/handlers"
)

func TestRegisterHandler(t *testing.T) {
	h := handlerSetup(t)

	testCases := []struct {
		name               string
		email              string
		firstName          string
		lastName           string
		password           string
		expectedErrMessage string
		expectedStatus     int
	}{
		{
			name:           "happy path",
			email:          "bananas@datum.net",
			firstName:      "Princess",
			lastName:       "Fiona",
			password:       "b!a!n!a!n!a!s!",
			expectedStatus: http.StatusCreated,
		},
		{
			name:               "duplicate email",
			email:              "bananas@datum.net",
			firstName:          "Princess",
			lastName:           "Fiona",
			password:           "b!a!n!a!n!a!s!",
			expectedErrMessage: "user already exists",
			expectedStatus:     http.StatusBadRequest,
		},
		{
			name:               "invalid email",
			email:              "bananas.net",
			firstName:          "Princess",
			lastName:           "Fiona",
			password:           "b!a!n!a!n!a!s!",
			expectedErrMessage: "email was invalid",
			expectedStatus:     http.StatusBadRequest,
		},
		{
			name:               "missing email",
			firstName:          "Princess",
			lastName:           "Fiona",
			password:           "b!a!n!a!n!a!s!",
			expectedErrMessage: "missing required field: email",
			expectedStatus:     http.StatusBadRequest,
		},
		{
			name:               "missing first name",
			email:              "tacos@datum.net",
			lastName:           "Fiona",
			password:           "b!a!n!a!n!a!s!",
			expectedErrMessage: "missing required field: first name",
			expectedStatus:     http.StatusBadRequest,
		},
		{
			name:               "missing last name",
			email:              "waffles@datum.net",
			firstName:          "Princess",
			password:           "b!a!n!a!n!a!s!",
			expectedErrMessage: "missing required field: last name",
			expectedStatus:     http.StatusBadRequest,
		},
		{
			name:               "bad password",
			email:              "pancakes@datum.net",
			firstName:          "Princess",
			lastName:           "Fiona",
			password:           "asfghjkl",
			expectedErrMessage: "password is too weak",
			expectedStatus:     http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create echo context
			e := echo.New()

			registerJSON := handlers.RegisterRequest{
				FirstName: tc.firstName,
				LastName:  tc.lastName,
				Email:     tc.email,
				Password:  tc.password,
			}

			body, err := json.Marshal(registerJSON)
			if err != nil {
				t.Error("error creating register json")
			}

			req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(string(body)))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			ctx := e.NewContext(req, recorder)

			err = h.RegisterHandler(ctx)
			require.NoError(t, err)

			res := recorder.Result()
			defer res.Body.Close()

			var out *handlers.RegisterReply

			// parse request body
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				t.Error("error parsing response", err)
			}

			assert.Equal(t, tc.expectedStatus, ctx.Response().Status)

			if tc.expectedStatus == http.StatusOK {
				assert.Equal(t, out.Email, tc.email)
				assert.NotEmpty(t, out.Message)
				assert.NotEmpty(t, out.ID)

				// cleanup after
				EntClient.User.DeleteOneID(out.ID).ExecX(context.Background())
			} else {
				assert.Contains(t, out.Message, tc.expectedErrMessage)
			}
		})
	}
}
