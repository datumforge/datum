package secure

import (
	"strings"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
)

// Secure is middleware that provides protection against cross-site scripting (XSS) attack,
// content type sniffing, clickjacking, insecure connection and other code injection attacks
func Secure() echo.MiddlewareFunc {
	secureConfig := middleware.SecureConfig{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
		//		HSTSMaxAge:            31536000, // nolint: gomnd
		//		HSTSExcludeSubdomains: true,
		//		HSTSPreloadEnabled:    true,
		ContentSecurityPolicy: "default-src 'self'",
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/api-docs")
		},
	}

	// TODO: make this configurable with inputs
	return middleware.SecureWithConfig(secureConfig)
}
