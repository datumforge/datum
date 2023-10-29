# Datum Core

This repo will hold the core server / handler for Datum services - for complete detailed references please check out the go-template used to generate this repository; it can be found [here](https://github.com/datumforge/go-template)

## Development

Datum's core server operates with the following utilities:

1. [ent](https://entgo.io/) - insane entity mapping tool, definitely not an ORM but kind of an ORM
1. [atlas](https://atlasgo.io/) - Schema generation and migration
1. [gqlgen](https://gqlgen.com/) - Code generation from schema definitions
1. [openfga](https://openfga.dev/) - Authorization

### Dependencies

Setup Taskfile by following the instructions and using one of the various convenient package managers or installation scripts. You can then simply run `task install` to load the associated dependencies. Nearly everything in this repository assumes you already have a local golang environment setup so this is not included. Please see the associated documentation.
