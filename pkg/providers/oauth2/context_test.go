package oauth2

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestContextState(t *testing.T) {
	expectedState := "state"
	ctx := WithState(context.Background(), expectedState)
	state, err := StateFromContext(ctx)
	assert.Equal(t, expectedState, state)
	assert.Nil(t, err)
}

func TestMissingState(t *testing.T) {
	state, err := StateFromContext(context.Background())
	assert.Equal(t, "", state)

	if assert.NotNil(t, err) {
		assert.Equal(t, ErrContextMissingStateValue, ErrContextMissingStateValue)
	}
}

func TestContextToken(t *testing.T) {
	expectedToken := &oauth2.Token{AccessToken: "access_token"}
	ctx := WithToken(context.Background(), expectedToken)
	token, err := TokenFromContext(ctx)
	assert.Equal(t, expectedToken, token)
	assert.Nil(t, err)
}

func TestFailTokenContext(t *testing.T) {
	token, err := TokenFromContext(context.Background())
	assert.Nil(t, token)

	if assert.NotNil(t, err) {
		assert.Equal(t, ErrContextMissingToken, ErrContextMissingToken)
	}
}
