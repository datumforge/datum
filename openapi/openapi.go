package openapi

// @title Datum API
// @version 1.0
// @description API Specifications for Datum Services
// @termsOfService https://datum.net/terms/

// @contact.name API Support
// @contact.url http://datum.net/support
// @contact.email support@datum.net

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.api_key
// @in header
// @name Authorization
// @type apiKey
// @description API Key Authentication

import (
	"os"
	"strings"

	"github.com/datumforge/datum/docs"
)

func ConfigSwagger() {
	var schemes []string
	host := os.Getenv("API_SERVER_URL")

	if host == "" {
		host = "http://localhost:17608"
	}

	if strings.HasPrefix(host, "http://") {
		schemes = append(schemes, "http")
		host = strings.Replace(host, "http://", "", 1)
	} else if strings.HasPrefix(host, "https://") {
		schemes = append(schemes, "https")
		host = strings.Replace(host, "https://", "", 1)
	}

	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Schemes = schemes
}
