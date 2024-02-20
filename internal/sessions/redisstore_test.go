package sessions_test

import (
	"context"
	"log"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/sessions"
	"github.com/datumforge/datum/internal/utils/ulids"
)

func TestExists(t *testing.T) {
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

			sessionID := sessions.GenerateSessionID()
			if tc.exists == int64(1) {
				err := ps.StoreSession(context.Background(), sessionID, tc.userID)
				require.NoError(t, err)
			}

			exists, err := ps.Exists(context.Background(), sessionID)
			require.NoError(t, err)
			assert.Equal(t, tc.exists, exists)
		})
	}
}

func TestGetSession(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		session string
		exists  bool
	}{
		{
			name:    "happy path",
			userID:  "MITB",
			session: ulids.New().String(),
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
				// store session in redis if the test expects it
				err := ps.StoreSession(context.Background(), tc.session, tc.userID)
				require.NoError(t, err)
			}

			// get stored value from redis
			value, err := ps.GetSession(context.Background(), tc.session)

			if tc.exists {
				require.NoError(t, err)
				assert.Equal(t, tc.userID, value)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestDeleteSession(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		session string
		exists  bool
	}{
		{
			name:    "happy path",
			userID:  "MITB",
			session: ulids.New().String(),
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
		Addr:             mr.Addr(),
		DisableIndentity: true,
	})

	return client
}
