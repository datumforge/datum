package apirouter

// Router is the interface for the router
type Router[HandlerFunc any, Route any] interface {
	// AddRoute adds a route to the router
	AddRoute(method string, path string, handler HandlerFunc) Route
	// OASHandler returns a handler that serves the OAS content
	OASHandler(contentType string, blob []byte) HandlerFunc
	// TransformPathToOASPath transforms the path to an OAS path
	TransformPathToOASPath(path string) string
}
