# OAS Schemagen

The initial purpose of this setup is so that eventually any of our repos would be able to import the package and use the methods to compose a wholly unique OpenAPI specification, append to an existing "base" specification, or merge between specifications.

This setup is not yet complete. There are weird quirks between all of the libraries and today `ogen` is controlling the largest swath of our generated specs, this is the first attempt at composing + outputting specifications based on the REST handlers and interfaces.

Eventually we may be able to use `ast`, `packages`, and other walking methods to reference / fetch the structs but for now the handler methods are added through basic functions which map them in the specification correctly.

## TO DO 

As you may be able to see from the `mitb.gen.yaml` file, the current setup is only composing `Components`, and the objects underneath it. The main remaining body of work is:

- Compose full `Operations` as a rollup of Requests, Responses, etc.
- Add methods to `range` over the registered routes and map them to `PathItems` which relate to `Operations`
- Add meethods to properly merge both json and yaml specifications
- Append the security method(s) to each operation
- Add examples to all `Components` and associated
- Add basic validations (e.g. every operation must have at least 1 4XX response)

redocly join openapi.json doc.json  -o apis-combined.yaml