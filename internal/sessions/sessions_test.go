package sessions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/internal/sessions"
	"github.com/datumforge/datum/internal/utils/ulids"
)

func TestSet(t *testing.T) {
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
			cs := sessions.NewCookieStore[string](sessions.DebugCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			session := cs.New(tc.name)

			// Set sessions
			session.Set(tc.userID, tc.session)

			assert.Equal(t, tc.session, session.Get(tc.userID))
		})
	}
}

func TestGetOk(t *testing.T) {
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
		{
			name:    "MeOWzErZ!",
			userID:  ulids.New().String(),
			session: "01HMDBSNBGH4DTEP0SR8118Y96",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cs := sessions.NewCookieStore[string](sessions.DebugCookieConfig,
				[]byte("my-signing-secret"), []byte("encryptionsecret"))

			s := cs.New(tc.name)

			s.Set("userID", tc.userID)
			s.Set("session", tc.session)

			uID, _ := s.GetOk("userID")
			sess, _ := s.GetOk("session")

			assert.Equal(t, tc.userID, uID)
			assert.Equal(t, tc.session, sess)
		})
	}
}
