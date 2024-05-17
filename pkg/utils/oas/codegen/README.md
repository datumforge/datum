# OAS Schemagen

The initial purpose of this setup is so that eventually any of our repos would be able to import the package and use the methods to compose a wholly unique OpenAPI specification, append to an existing "base" specification, or merge between specifications.

This setup is not yet complete. Eventually we may be able to use `ast`, `packages`, and other walking methods to reference / fetch the structs but for now the handler methods are added through basic functions which map them in the specification correctly.

## TO DO

- Add examples to all `Components` and associated
