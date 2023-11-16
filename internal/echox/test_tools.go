package echox

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// newEchoContext used for testing purposes ONLY
func newEchoContext() echo.Context {
	// create echo context
	e := echo.New()
	req := &http.Request{
		Header: http.Header{},
	}
	res := &echo.Response{}

	return e.NewContext(req, res)
}

// newValidSignedJWT creates a jwt with a fake subject for testing purposes ONLY
func newValidSignedJWT() (*jwt.Token, error) {
	iat := time.Now().Unix()
	nbf := time.Now().Unix()
	exp := time.Now().Add(time.Hour).Unix()

	claims := jwt.MapClaims{
		"sub":    "foobar",
		"issuer": "test suite",
		"iat":    iat,
		"nbf":    nbf,
		"exp":    exp,
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwt, nil
}

// NewContextWithValidUser creates an echo context with a fake subject for testing purposes ONLY
func NewContextWithValidUser() (*echo.Context, error) {
	ec := newEchoContext()

	j, err := newValidSignedJWT()
	if err != nil {
		return nil, err
	}

	ec.Set("user", j)

	return &ec, nil
}
