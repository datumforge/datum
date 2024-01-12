package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticHandler(t *testing.T) {
	h := handlerSetup(t, EntClient)

	testCases := []struct {
		name         string
		path         string
		expectStatus int
	}{
		{
			"without .txt",
			"/security",
			http.StatusNotFound,
		},
		{
			"full call",
			"/security.txt",
			http.StatusOK,
		},
		{
			"with v1",
			"/v1/security.txt",
			http.StatusNotFound,
		},
		{
			"without .txt",
			"/robots",
			http.StatusNotFound,
		},
		{
			"full call",
			"/robots.txt",
			http.StatusOK,
		},
		{
			"with v1",
			"/v1/robots.txt",
			http.StatusNotFound,
		},
		{
			"full",
			"/.well-known",
			http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create echo context with middleware
			e := setupEchoAuth(h.SM, EntClient)

			e.GET("security.txt", h.SecurityHandler)
			e.GET("robots.txt", h.RobotsHandler)
			e.GET(".well-known", h.JWKSWellKnownHandler)

			req := httptest.NewRequest(http.MethodGet, tc.path, nil)

			// Set writer for tests that write on the response
			recorder := httptest.NewRecorder()

			// Using the ServerHTTP on echo will trigger the router and middleware
			e.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.expectStatus, recorder.Code)

			defer res.Body.Close()
		})
	}
}
