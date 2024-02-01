package sessions_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/sessions"
	"github.com/datumforge/datum/internal/utils/ulids"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name         string
		sessionName  string
		expectedName string
	}{
		{
			name:         "happy path",
			sessionName:  "huddle",
			expectedName: "huddle",
		},
		{
			name:         "empty name, use default",
			sessionName:  "",
			expectedName: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cs := sessions.NewCookieStore[string](sessions.DefaultCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			session := cs.New(tc.sessionName)

			assert.Equal(t, tc.expectedName, session.Name())
		})
	}
}

func TestNewSessionCookie(t *testing.T) {
	tests := []struct {
		name    string
		session string
	}{
		{
			name:    "happy path",
			session: ulids.New().String(),
		},
		{
			name:    "empty string still results in session",
			session: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cs := sessions.NewCookieStore[string](sessions.DebugCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			cooky := cs.New(tc.name)

			cooky.Set("name", tc.name)
			cooky.Set("session", tc.session)

			name := cooky.Get("name")
			session := cooky.Get("session")

			assert.Equal(t, tc.name, name)
			assert.Equal(t, tc.session, session)
		})
	}
}

func TestSaveGet(t *testing.T) {
	tests := []struct {
		name    string
		session string
		userID  string
	}{
		{
			name:    "happy path",
			session: ulids.New().String(),
			userID:  "mitb",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cs := sessions.NewCookieStore[map[string]string](sessions.DebugCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			session := sessions.NewSession(cs, tc.name)
			sessionID := sessions.GenerateSessionID()

			setSessionMap := map[string]string{}
			setSessionMap["userID"] = tc.userID
			setSessionMap["name"] = tc.name
			setSessionMap["session"] = tc.session

			session.Set(sessionID, setSessionMap)

			err := cs.Save(recorder, session)
			require.NoError(t, err)

			// Copy the Cookie over to a new Request
			res := recorder.Result()
			defer res.Body.Close()

			cooky := res.Header["Set-Cookie"]
			request := &http.Request{Header: http.Header{"Cookie": cooky}}

			sess, err := cs.Get(request, tc.name)
			require.NoError(t, err)

			assert.Equal(t, tc.name, sess.Name)
		})
	}
}

func TestGetUserFromSession(t *testing.T) {
	tests := []struct {
		name    string
		session string
		userID  string
	}{
		{
			name:    "happy path",
			session: ulids.New().String(),
			userID:  "mitb",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cs := sessions.NewCookieStore[map[string]string](sessions.DebugCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			session := cs.New(tc.name)
			sessionID := sessions.GenerateSessionID()
			setSessionMap := map[string]string{}
			setSessionMap["userID"] = tc.userID
			setSessionMap["name"] = tc.name
			setSessionMap["session"] = tc.session

			session.Set(sessionID, setSessionMap)

			err := session.Save(recorder)
			require.NoError(t, err)

			// Copy the Cookie over to a new Request
			res := recorder.Result()
			defer res.Body.Close()

			cooky := res.Header["Set-Cookie"]
			request := &http.Request{Header: http.Header{"Cookie": cooky}}

			userID, err := sessions.UserIDFromContext(request.Context())
			require.NoError(t, err)

			require.Equal(t, tc.userID, userID)
		})
	}
}
