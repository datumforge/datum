package sessions_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/cookies"
	"github.com/datumforge/datum/internal/sessions"
	"github.com/datumforge/datum/internal/utils/ulids"
)

func Test_New(t *testing.T) {
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
			expectedName: sessions.DefaultSessionName,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cs := sessions.NewCookieStore(&cookies.DebugOnlyCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			session := cs.New(tc.sessionName)

			assert.Equal(t, tc.expectedName, session.Name())
		})
	}
}

func Test_NewSessionCookie(t *testing.T) {
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
			cooky := sessions.NewSessionCookie(tc.session)

			assert.Equal(t, sessions.DefaultSessionName, cooky.Name)
			assert.Equal(t, tc.session, cooky.Value)
			assert.Equal(t, true, cooky.Secure)
		})
	}
}

func Test_NewDebugSessionCookie(t *testing.T) {
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
			cooky := sessions.NewDebugSessionCookie(tc.session)

			assert.Equal(t, sessions.DefaultSessionName, cooky.Name)
			assert.Equal(t, tc.session, cooky.Value)
			assert.Equal(t, false, cooky.Secure) // debug cookies should have secure off
		})
	}
}

func Test_SaveGet(t *testing.T) {
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
			cs := sessions.NewCookieStore(&cookies.DebugOnlyCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			session := cs.New(sessions.DefaultSessionName)
			session.Set(tc.userID, tc.session)

			err := cs.Save(recorder, session)
			require.NoError(t, err)

			// Copy the Cookie over to a new Request
			res := recorder.Result()
			defer res.Body.Close()

			cooky := res.Header["Set-Cookie"]
			request := &http.Request{Header: http.Header{"Cookie": cooky}}

			sess, err := cs.Get(request, sessions.DefaultSessionName)
			require.NoError(t, err)

			assert.Equal(t, tc.session, sess.Get(tc.userID))
		})
	}
}

func Test_GetUserFromSession(t *testing.T) {
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
			cs := sessions.NewCookieStore(&cookies.DebugOnlyCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			session := cs.New(sessions.DefaultSessionName)
			session.Set(tc.userID, tc.session)

			err := cs.Save(recorder, session)
			require.NoError(t, err)

			// Copy the Cookie over to a new Request
			res := recorder.Result()
			defer res.Body.Close()

			cooky := res.Header["Set-Cookie"]
			request := &http.Request{Header: http.Header{"Cookie": cooky}}

			userID, err := cs.GetUserFromSession(request, sessions.DefaultSessionName)
			require.NoError(t, err)

			require.Equal(t, tc.userID, userID)
		})
	}
}
