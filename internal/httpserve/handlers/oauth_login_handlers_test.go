package handlers_test

import (
	"testing"
	"time"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/httpserve/middleware/transaction"
)

func TestHandlerCheckAndCreateUser(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// add login handler
	client.e.POST("login", client.h.LoginHandler)

	ec := echocontext.NewTestEchoContext().Request().Context()

	// set privacy allow in order to allow the creation of the users without
	// authentication in the tests
	ctx := privacy.DecisionContext(ec, privacy.Allow)

	type args struct {
		name     string
		email    string
		provider enums.AuthProvider
	}

	tests := []struct {
		name    string
		args    args
		want    *ent.User
		writes  bool
		wantErr bool
	}{
		{
			name: "happy path, github",
			args: args{
				name:     "Wanda Maximoff",
				email:    "wmaximoff@marvel.com",
				provider: enums.GitHub,
			},
			want: &ent.User{
				FirstName:    "Wanda",
				LastName:     "Maximoff",
				Email:        "wmaximoff@marvel.com",
				AuthProvider: enums.GitHub,
			},
			writes: true,
		},
		{
			name: "happy path, same email, different provider",
			args: args{
				name:     "Wanda Maximoff",
				email:    "wmaximoff@marvel.com",
				provider: enums.Google,
			},
			want: &ent.User{
				FirstName:    "Wanda",
				LastName:     "Maximoff",
				Email:        "wmaximoff@marvel.com",
				AuthProvider: enums.Google,
			},
			writes: true,
		},
		{
			name: "user already exists, should not fail, just update last seen",
			args: args{
				name:     "Wanda Maximoff",
				email:    "wmaximoff@marvel.com",
				provider: enums.Google,
			},
			want: &ent.User{
				FirstName:    "Wanda",
				LastName:     "Maximoff",
				Email:        "wmaximoff@marvel.com",
				AuthProvider: enums.Google,
			},
			writes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.writes {
				// add mocks for writes when a new user is created
				mock_fga.WriteOnce(t, client.fga)
			}

			now := time.Now()

			// start transaction because the query expects a transaction in the context
			tx, err := client.h.DBClient.Tx(ctx)
			require.NoError(t, err)

			// commit transaction after test finishes
			defer tx.Commit() //nolint:errcheck

			// set transaction in the context
			ctx = transaction.NewContext(ctx, tx)

			got, err := client.h.CheckAndCreateUser(ctx, tt.args.name, tt.args.email, tt.args.provider)
			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, got)

				return
			}

			// check if user was created
			require.NoError(t, err)
			require.NotNil(t, got)

			// verify fields
			assert.Equal(t, tt.want.FirstName, got.FirstName)
			assert.Equal(t, tt.want.LastName, got.LastName)
			assert.Equal(t, tt.want.Email, got.Email)
			assert.Equal(t, tt.want.AuthProvider, got.AuthProvider)
			assert.WithinDuration(t, now, *got.LastSeen, time.Second*5)
		})
	}
}
