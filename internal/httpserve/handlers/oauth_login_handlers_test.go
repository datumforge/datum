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
	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/middleware/transaction"
)

func (suite *HandlerTestSuite) TestHandlerCheckAndCreateUser() {
	t := suite.T()

	// add login handler
	suite.e.POST("login", suite.h.LoginHandler)

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
				provider: enums.AuthProviderGitHub,
			},
			want: &ent.User{
				FirstName:    "Wanda",
				LastName:     "Maximoff",
				Email:        "wmaximoff@marvel.com",
				AuthProvider: enums.AuthProviderGitHub,
			},
			writes: true,
		},
		{
			name: "happy path, same email, different provider",
			args: args{
				name:     "Wanda Maximoff",
				email:    "wmaximoff@marvel.com",
				provider: enums.AuthProviderGoogle,
			},
			want: &ent.User{
				FirstName:    "Wanda",
				LastName:     "Maximoff",
				Email:        "wmaximoff@marvel.com",
				AuthProvider: enums.AuthProviderGoogle,
			},
			writes: true,
		},
		{
			name: "user already exists, should not fail, just update last seen",
			args: args{
				name:     "Wanda Maximoff",
				email:    "wmaximoff@marvel.com",
				provider: enums.AuthProviderGoogle,
			},
			want: &ent.User{
				FirstName:    "Wanda",
				LastName:     "Maximoff",
				Email:        "wmaximoff@marvel.com",
				AuthProvider: enums.AuthProviderGoogle,
			},
			writes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.writes {
				// add mocks for writes when a new user is created
				mock_fga.WriteOnce(t, suite.fga)
			}

			now := time.Now()

			// start transaction because the query expects a transaction in the context
			tx, err := suite.h.DBClient.Tx(ctx)
			require.NoError(t, err)

			// commit transaction after test finishes
			defer tx.Commit() //nolint:errcheck

			// set transaction in the context
			ctx = transaction.NewContext(ctx, tx)

			got, err := suite.h.CheckAndCreateUser(ctx, tt.args.name, tt.args.email, tt.args.provider)
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
