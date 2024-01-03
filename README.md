[![Build status](https://badge.buildkite.com/a3a38b934ca2bb7fc771e19bc5a986a1452fa2962e4e1c63bf.svg?branch=main)](https://buildkite.com/datum/datum)

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

Setup [Taskfile](https://taskfile.dev/installation/) by following the instructions and using one of the various convenient package managers or installation scripts. Two of the more common installation methods are below for your convenience:

```
brew install go-task
```

```
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d
```
(by default, this installs on the ``./bin`` directory relative to the working directory)

After installation, you can then simply run `task install` to load the associated dependencies. Nearly everything in this repository assumes you already have a local golang environment setup so this is not included. Please see the associated documentation.

To include Taskfile's created in other directories / to call the respective tasks, you would add an `includes` per the Taskfile documentation and then reference it by name, e.g. `task cli:createorg`

### Updating Environment Variables

Within the `config` directory in the root of this repository there are several `.example` files prefixed with `.env-dev` or similar; these hold examples of environment configurations which you should review and potentially override depending on your needs. Anything which is launched out of the `Taskfile` will source it's environment from these files and their configurations. Different tasks can be made to source from different files as can be seen by several of the tasks within the Taskfile.

You will need to perform a 1-time action of either removing the `.example` suffix from these files or creating your own files which match the naming convensions called for `{{.ENV}}` within the Taskfile. These files are intentionally added to the `.gitignore` within this repository to prevent you from accidentally committing secrets or other sensitive information which may live inside the server's environment variables.

### Pre-requisites to a PR

This repository contains a number of code generating functions / utilities which take schema modifications and scaffold out resolvers, graphql API schemas, openAPI specifications, among other things. To ensure you've generated all the necessary dependencies run `task pr`; this will run the entirety of the commands required to safely generate a PR. If for some reason one of the commands fails / encounters an error, you will need to debug the individual steps. It should be decently easy to follow the `Taskfile` in the root of this repository.

## Querying

The best method of forming / testing queries against the server is to run `task rover` which will launch an interactive query UI.

## OpenFGA Playground

You can load up a local openFGA environment with the compose setup in this repository; `task fga:up` - this will launch an interactive playground where you can model permissions model(s) or changes to the models

## Migrations

`task atlas` or `task atlas:create` will generate the necessary migrations

## Creating a new Schema

To ease the effort required to add additional schemas into the system a template + task function has been created. This isn't doing anything terribly complex, but it's attempting to ensure you have the _minimum_ set of required things needed to create a schema - most notably: you need to ensure the IDMixin is present (otherwise you will get ID type conflicts) and a standard set of schema annotations.

**NOTE: you still have to make intelligent decisions around things like the presence / integration of hooks, interceptors, policies, etc. This is saving you about 10 seconds of copy-paste, so don't over estimate the automation, here.

To generate a new schema, you can run `task newschema -- [yourschemaname]` where you replace the name within `[]`. Please be sure to note that this isn't a command line flag so there's a space between `--` and the name.