package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/stretchr/testify/assert"

	echo "github.com/datumforge/echox"
)

var sessionManager = scs.New()

func TestSession(t *testing.T) {
	sessionManager.Lifetime = 24 * time.Hour

	e := echo.New()

	// Call /put to set the message in the session manager
	req := httptest.NewRequest(http.MethodGet, "/put", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	defer c.Request().Response.Request.Body.Close()

	session := LoadAndSave(sessionManager)

	h := session(func(c echo.Context) error {
		sessionManager.Put(c.Request().Context(), "message", "Hello from a session!")
		return c.String(http.StatusOK, "")
	})

	h(c)

	assert.Equal(t, rec.Result().StatusCode, 200)
	assert.Equal(t, len(rec.Result().Cookies()), 1)

	sessionCookie := rec.Result().Cookies()[0]

	defer c.Request().Response.Request.Body.Close()

	assert.Equal(t, sessionCookie.Name, "session")

	// Make a request to /get to see if the message is still there
	req = httptest.NewRequest(http.MethodGet, "/get", nil)
	req.Header.Set("Cookie", sessionCookie.Name+"="+sessionCookie.Value)

	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	h = session(func(c echo.Context) error {
		msg := sessionManager.GetString(c.Request().Context(), "message")
		return c.String(http.StatusOK, msg)
	})

	h(c)

	assert.Equal(t, rec.Result().StatusCode, 200)
	assert.Equal(t, rec.Body.String(), "Hello from a session!")
}
