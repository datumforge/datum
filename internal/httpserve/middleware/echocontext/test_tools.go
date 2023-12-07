package echocontext

import (
	"net/http"
	"time"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"
)

// NewTestEchoContext used for testing purposes ONLY
func NewTestEchoContext() echo.Context {
	// create echo context
	e := echo.New()
	req := &http.Request{
		Header: http.Header{},
	}
	res := &echo.Response{}

	return e.NewContext(req, res)
}

// newValidSignedJWT creates a jwt with a fake subject for testing purposes ONLY
func newValidSignedJWT(subject string) (*jwt.Token, error) {
	iat := time.Now().Unix()
	nbf := time.Now().Unix()
	exp := time.Now().Add(time.Hour).Unix()

	claims := jwt.MapClaims{
		"sub":    subject,
		"issuer": "test suite",
		"iat":    iat,
		"nbf":    nbf,
		"exp":    exp,
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwt, nil
}

// NewTestContextWithValidUser creates an echo context with a fake subject for testing purposes ONLY
func NewTestContextWithValidUser(subject string) (*echo.Context, error) {
	ec := NewTestEchoContext()

	j, err := newValidSignedJWT(subject)
	if err != nil {
		return nil, err
	}

	ec.Set("user", j)

	return &ec, nil
}
