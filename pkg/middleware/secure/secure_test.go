package secure_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/datumforge/echox"
	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/pkg/middleware/secure"
)

func TestSecureMiddleware(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up the middleware
	mw := secure.Secure()

	// Invoke the middleware
	err := mw(func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})(c)

	// Assert that the middleware did not return an error
	assert.NoError(t, err)

	// Assert that the response has the expected headers
	assert.Equal(t, "1; mode=block", rec.Header().Get("X-XSS-Protection"))
	assert.Equal(t, "nosniff", rec.Header().Get("X-Content-Type-Options"))
	assert.Equal(t, "SAMEORIGIN", rec.Header().Get("X-Frame-Options"))
	assert.Equal(t, "default-src 'self'", rec.Header().Get("Content-Security-Policy"))

	// Assert that the response body is "OK"
	assert.Equal(t, "OK", rec.Body.String())
}
