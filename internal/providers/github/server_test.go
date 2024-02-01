package github

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/datumforge/datum/internal/testutils"
)

// newGithubTestServer mocks the GitHub user endpoint and a client
func newGithubTestServer(routePrefix, jsonData string) (*http.Client, *httptest.Server) {
	client, mux, server := testutils.TestServer()
	mux.HandleFunc(routePrefix+"/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, jsonData)
	})

	return client, server
}
