package apirouter

import (
	"strings"
)

// TransformPathParamsWithColon replaces the colon in the path with ze curly braces {}
func TransformPathParamsWithColon(path string) string {
	pathParams := strings.Split(path, "/")
	for i, param := range pathParams {
		if strings.HasPrefix(param, ":") {
			pathParams[i] = strings.Replace(param, ":", "{", 1) + "}"
		}
	}

	return strings.Join(pathParams, "/")
}
