package auth

import (
	echojwt "github.com/datumforge/echo-jwt/v5"
	echo "github.com/datumforge/echox"
)

// CreateJwtMiddleware creates a middleware function for JWTs
// TODO expand the config settings
func CreateJwtMiddleware(secret []byte) echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: secret,
	}

	return echojwt.WithConfig(config)
}
