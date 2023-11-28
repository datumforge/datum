# Datum Core

> This repository is experimental meaning that it's based on untested ideas or techniques and not yet established or finalized or involves a radically new and innovative style!
> This means that support is best effort (at best!) and we strongly encourage you to NOT use this in production - reach out to [@matoszz](https://github.com/matoszz) with any questions

This repo will hold the core server / handler for Datum services - for complete detailed references please check out the go-template used to generate this repository; it can be found [here](https://github.com/datumforge/go-template)

## Development

Datum's core server operates with the following utilities:

1. [ent](https://entgo.io/) - insane entity mapping tool, definitely not an ORM but kind of an ORM
1. [atlas](https://atlasgo.io/) - Schema generation and migration
1. [gqlgen](https://gqlgen.com/) - Code generation from schema definitions
1. [openfga](https://openfga.dev/) - Authorization

### Dependencies

Setup [Taskfile](https://taskfile.dev/installation/) by following the instructions and using one of the various convenient package managers or installation scripts. You can then simply run `task install` to load the associated dependencies. Nearly everything in this repository assumes you already have a local golang environment setup so this is not included. Please see the associated documentation.

To include Taskfile's created in other directories / to call the respective tasks, you would add an `includes` per the Taskfile documentation and then reference it by name, e.g. `task cli:createorg`

### Pre-requisites to a PR

This repository contains a number of code generating functions / utilities which take schema modifications and scaffold out resolvers, graphql API schemas, openAPI specifications, among other things. To ensure you've generated all the necessary dependencies run `task pr`; this will run the entirety of the commands required to safely generate a PR. If for some reason one of the commands fails / encounters an error, you will need to debug the individual steps. It should be decently easy to follow the `Taskfile` in the root of this repository.

## Querying

The best method of forming / testing queries against the server is to run `task rover` which will launch an interactive query UI.

## OpenFGA Playground

You can load up a local openFGA environment with the compose setup in this repository; `task fga:up` - this will launch an interactive playground where you can model permissions model(s) or changes to the models

## Migrations

`task atlas` or `task atlas:create` will generate the necessary migrations
