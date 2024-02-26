# OpenAPI specifications

Manual management of OpenAPI specifications is cumbersome and error-prone, leading to inconsistencies between the actual implementation and
the documentation. Currently (as of my knowledge) there is no out of the box solution for generating OpenAPI specifications in Go based on code.
This package is intended to remedy this, by providing the possibility to generate `json` based OpenAPI specifications using code. In order to
generate OpenAPI specifications, this package uses the amazing `kin-openapi` project and custom type definitions, which can be found in `types.go`.

## Why did you make this?

Well, first, I'll just go fuck myself.




go-openapi does not suuport 3.x
https://github.com/go-openapi/spec?tab=readme-ov-file#faq


go-swagger uses go-openapi and also thus, does not support current specifications
https://github.com/go-swagger/go-swagger


https://github.com/go-openapi/swag
also unusable


https://github.com/ogen-go/ogen
is closer to what we want


https://github.com/oapi-codegen/echo-middleware


https://spec.openapis.org/oas/latest.html

3.1.0


npx @redocly/cli@latest push [openapi.yaml] --destination="Datum API@v2" --organization="datum"
