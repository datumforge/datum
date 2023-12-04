// Package webmiddleware provides middleware for http Handlers.
package middleware

import "net/http"

// MiddlewareFunc is a function that acts as middleware for http Handlers.
type MiddlewareFunc func(next http.Handler) http.Handler

// Chain returns a http.Handler that chains the middleware onion-style around the handler.
func Chain(middlewares []MiddlewareFunc, handler http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// Conditional is a middleware that only executes middleware if the condition
// returns true for the request. If the condition returns false, the middleware
// is skipped, and request handling moves on to the next handler in the chain.
func Conditional(middleware MiddlewareFunc, condition func(r *http.Request) bool) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler := next
			if condition(r) {
				handler = middleware(next)
			}
			handler.ServeHTTP(w, r)
		})
	}
}
