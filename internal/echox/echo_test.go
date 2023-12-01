package echox

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_GetActorSubject(t *testing.T) {
	// context with no user set
	basicContext := NewTestEchoContext()

	missingSubCtx := NewTestEchoContext()
	jBasic := jwt.New(jwt.SigningMethodHS256)
	missingSubCtx.Set("user", jBasic)

	validCtx, err := NewTestContextWithValidUser("foobar")
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
			e:    *validCtx,
			err:  nil,
		},
		{
			name: "no user",
			e:    basicContext,
			err:  ErrJWTMissingInvalid,
		},
		{
			name: "no user",
			e:    missingSubCtx,
			err:  ErrSubjectNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			sub, err := GetActorSubject(tc.e)
			if tc.err != nil {
				assert.Error(t, err)
				assert.Empty(t, sub)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, "foobar", sub)
		})
	}
}
