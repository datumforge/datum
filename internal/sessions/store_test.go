package sessions_test

import (
	"context"
	"log"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/sessions"
)

func Test_Exists(t *testing.T) {
	tests := []struct {
		name   string
		userID string
		exists int64
	}{
		{
			name:   "happy path",
			userID: "MITB",
			exists: 1,
		},
		{
			name:   "session does not exist",
			userID: "SITB",
			exists: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rc := newRedisClient()
			ps := sessions.NewStore(rc)

			if tc.exists == int64(1) {
				err := ps.StoreSession(context.Background(), gofakeit.UUID(), tc.userID)
				require.NoError(t, err)
			}

			exists, err := ps.Exists(context.Background(), tc.userID)
			require.NoError(t, err)
			assert.Equal(t, tc.exists, exists)
		})
	}
}

func Test_GetSession(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		session string
		exists  bool
	}{
		{
			name:    "happy path",
			userID:  "MITB",
			session: gofakeit.UUID(),
			exists:  true,
		},
		{
			name:   "session does not exist",
			userID: "SITB",
			exists: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rc := newRedisClient()
			ps := sessions.NewStore(rc)

			if tc.exists {
				err := ps.StoreSession(context.Background(), tc.session, tc.userID)
				require.NoError(t, err)
			}

			sessionID, err := ps.GetSession(context.Background(), tc.userID)

			if tc.exists {
				require.NoError(t, err)
				assert.Equal(t, tc.session, sessionID)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_DeleteSession(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		session string
		exists  bool
	}{
		{
			name:    "happy path",
			userID:  "MITB",
			session: gofakeit.UUID(),
			exists:  true,
		},
		{
			name:   "session does not exist, should not error",
			userID: "SITB",
			exists: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rc := newRedisClient()
			ps := sessions.NewStore(rc)

			if tc.exists {
				err := ps.StoreSession(context.Background(), tc.session, tc.userID)
				require.NoError(t, err)
			}

			err := ps.DeleteSession(context.Background(), tc.userID)
			require.NoError(t, err)
		})
	}
}

func newRedisClient() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client
}
