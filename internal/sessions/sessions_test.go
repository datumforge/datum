package sessions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/internal/cookies"
	"github.com/datumforge/datum/internal/sessions"
	"github.com/datumforge/datum/internal/utils/ulids"
)

func Test_Set(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		session string
	}{
		{
			name:    "happy path",
			userID:  "01HMDBSNBGH4DTEP0SR8118Y96",
			session: ulids.New().String(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cs := sessions.NewCookieStore(&cookies.DebugOnlyCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			session := sessions.New(cs)

			// Set sessions
			session.Set(tc.userID, tc.session)

			assert.Equal(t, tc.session, session.Get(tc.userID))
		})
	}
}

func Test_GetOk(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		session string
		exists  bool
	}{
		{
			name:    "happy path",
			userID:  "01HMDBSNBGH4DTEP0SR8118Y96",
			session: ulids.New().String(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cs := sessions.NewCookieStore(&cookies.DebugOnlyCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			s := sessions.New(cs)

			// Set sessions
			if tc.exists {
				s.Set(tc.userID, tc.session)
			}

			session, ok := s.GetOk(tc.userID)
			assert.Equal(t, tc.exists, ok)

			if tc.exists {
				assert.Equal(t, tc.session, session)
			}
		})
	}
}
