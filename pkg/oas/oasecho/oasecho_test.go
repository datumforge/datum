package oasecho_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	echo "github.com/datumforge/echox"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/require"

	openapi "github.com/datumforge/datum/pkg/oas"
	oasEcho "github.com/datumforge/datum/pkg/oas/oasecho"
)

const (
	OpenAPITitle   = "test openapi title"
	OpenAPIVersion = "test openapi version"
)

type echoOASRouter = openapi.Router[echo.HandlerFunc, echo.RouteInfo]

func TestEchoIntegration(t *testing.T) {
	t.Run("router works correctly - echo", func(t *testing.T) {
		echoRouter, oasRouter := setupEchoOpenAPI(t)

		err := oasRouter.GenerateAndExposeOpenAPI()
		require.NoError(t, err)

		t.Run("/hello", func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/hello", nil)

			echoRouter.ServeHTTP(w, r)

			require.Equal(t, http.StatusOK, w.Result().StatusCode) //nolint: bodyclose

			body := readBody(t, w.Result().Body) //nolint: bodyclose
			require.Equal(t, "OK", body)
		})

		t.Run("/hello/:value", func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/hello/something", nil)

			echoRouter.ServeHTTP(w, r)

			require.Equal(t, http.StatusOK, w.Result().StatusCode) //nolint: bodyclose

			body := readBody(t, w.Result().Body) //nolint: bodyclose
			require.Equal(t, "OK", body)
		})

		t.Run("and generate openapi", func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, openapi.DefaultJSONDocumentationPath, nil)

			echoRouter.ServeHTTP(w, r)

			require.Equal(t, http.StatusOK, w.Result().StatusCode) //nolint: bodyclose

			body := readBody(t, w.Result().Body) //nolint: bodyclose
			require.JSONEq(t, readFile(t, "./testdata/integration.json"), body)
		})
	})
}

func readBody(t *testing.T, requestBody io.ReadCloser) string {
	t.Helper()

	body, err := io.ReadAll(requestBody)
	require.NoError(t, err)

	return string(body)
}

func setupEchoOpenAPI(t *testing.T) (*echo.Echo, *echoOASRouter) {
	t.Helper()

	context := context.Background()
	e := echo.New()

	router, err := openapi.NewRouter(oasEcho.NewRouter(e), openapi.Options{
		Context: context,
		OpenAPI: &openapi3.T{
			Info: &openapi3.Info{
				Title:   OpenAPITitle,
				Version: OpenAPIVersion,
			},
		},
	})
	require.NoError(t, err)

	operation := openapi.Operation{}

	_, err = router.AddRawRoute(http.MethodGet, "/hello", okHandler, operation)
	require.NoError(t, err)

	_, err = router.AddRoute(http.MethodPost, "/hello/:value", okHandler, openapi.Definitions{})
	require.NoError(t, err)

	return e, router
}

func okHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func readFile(t *testing.T, path string) string {
	t.Helper()

	fileContent, err := os.ReadFile(path)
	require.NoError(t, err)

	return string(fileContent)
}
