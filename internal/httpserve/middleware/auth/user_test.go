package auth_test

import (
	"testing"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/utils/ulids"
)

func Test_GetActorUserID(t *testing.T) {
	// context with no user set
	basicContext := echocontext.NewTestEchoContext()

	missingSubCtx := echocontext.NewTestEchoContext()
	jBasic := jwt.New(jwt.SigningMethodHS256)
	missingSubCtx.Set("user", jBasic)

	sub := ulids.New().String()

	validCtx, err := auth.NewTestContextWithValidUser(sub)
	if err != nil {
		t.Fatal()
	}

	invalidUserCtx, err := auth.NewTestContextWithValidUser(ulids.Null.String())
	if err != nil {
		t.Fatal()
	}

	testCases := []struct {
		name string
		e    echo.Context
		err  error
	}{
		{
			name: "happy path",
			e:    validCtx,
			err:  nil,
		},
		{
			name: "no user",
			e:    basicContext,
			err:  auth.ErrNoClaims,
		},
		{
			name: "no user",
			e:    missingSubCtx,
			err:  auth.ErrNoClaims,
		},
		{
			name: "null user",
			e:    invalidUserCtx,
			err:  auth.ErrNoUserInfo,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			got, err := auth.GetActorUserID(tc.e)
			if tc.err != nil {
				assert.Error(t, err)
				assert.Empty(t, got)
				assert.ErrorContains(t, err, tc.err.Error())

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, sub, got)
		})
	}
}
