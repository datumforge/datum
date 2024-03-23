package oasecho

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/pkg/oas/apirouter"
)

// Route is a type alias for echo.RouteInfo
type Route = *echo.Route

// echoRouter is a struct that implements the apirouter.Router interface
type echoRouter struct {
	router *echo.Echo
}

// AddRoute adds a route to the router
func (r echoRouter) AddRoute(method string, path string, handler echo.HandlerFunc) echo.RouteInfo {
	return r.router.Add(method, path, handler)
}

// TransformPathToOASPath function in the `echoRouter` struct is responsible for transforming a
// given path to an OpenAPI Specification (OAS) compliant path
func (r echoRouter) TransformPathToOASPath(path string) string {
	return apirouter.TransformPathParamsWithColon(path)
}

// OASHandler function is defining a method on the `echoRouter` struct that
// returns a handler function for serving OAS content. This method takes two parameters: `contentType`
// which specifies the content type of the response, and `blob` which is the byte array containing the
// OAS content (the OAS specification in JSON format)
func (r echoRouter) OASHandler(contentType string, blob []byte) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Content-Type", contentType)
		return c.JSONBlob(http.StatusOK, blob)
	}
}

// NewRouter function is a constructor function that creates and returns a new instance of
// the `echoRouter` struct, which implements the `apirouter.Router` interface. It takes a pointer to an
// `echo.Echo` instance as a parameter and initializes the `echoRouter` with that instance. This
// function is used to create a new router for handling routes in an Echo framework based application.
func NewRouter(router *echo.Echo) apirouter.Router[echo.HandlerFunc, echo.RouteInfo] {
	return echoRouter{
		router: router,
	}
}
