[![Build status](https://badge.buildkite.com/a3a38b934ca2bb7fc771e19bc5a986a1452fa2962e4e1c63bf.svg?branch=main)](https://buildkite.com/datum/datum) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=datumforge_datum&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=datumforge_datum)

# Datum Core

> This repository is experimental meaning that it's based on untested ideas or techniques and not yet established or finalized or involves a radically new and innovative style!
> This means that support is best effort (at best!) and we strongly encourage you to NOT use this in production - reach out to [@matoszz](https://github.com/matoszz) with any questions

## Development

Datum's core server operates with the following utilities:

1. [ent](https://entgo.io/) - insane entity mapping tool, definitely not an ORM but kind of an ORM
1. [atlas](https://atlasgo.io/) - Schema generation and migration
1. [gqlgen](https://gqlgen.com/) - Code generation from schema definitions
1. [gqlgenc](https://github.com/Yamashou/gqlgenc) - client building utilities with GraphQL
1. [openfga](https://openfga.dev/) - Authorization
1. [echo](https://echo.labstack.com/) - High performance, extensible, minimalist Go web framework
1. [koanf](github.com/knadh/koanf) - configuration management
1. [viper](https://github.com/spf13/viper) - command line flags / management

We also leverage many secondary technologies in use, including (but not limited to!):

1. [redis](https://redis.io/) - in-memory datastore used for sessions, caching
1. [sqlite](https://www.sqlite.org/) - currently planned database system but also offer additional support for PostgreSQL
1. [golangci-lint](https://github.com/golangci/golangci-lint) - an annoyingly opinionated linter
1. [buildkite](https://buildkite.com/datum) - our CI system of choice (with github actions providing some intermediary support)
1. [sonar](https://sonarcloud.io/summary/overall?id=datumforge_datum) - used for code scanning, vulnerability scanning
1. [posthog](https://posthog.com/) - product analytics
1. [sentry](https://sentry.io) - error montioring, tracing
1. [sendgrid](https://sendgrid.com/en-us) - transactional email send provider

All of these components are bundled into our respective Docker images; for additional information / instructions, see the [contributing guide](.github/CONTRIBUTING.md) in this repository. 

### Dependencies

Setup [Taskfile](https://taskfile.dev/installation/) by following the instructions and using one of the various convenient package managers or installation scripts. After installation, you can then simply run `task install` to load the associated dependencies. Nearly everything in this repository assumes you already have a local golang environment setup so this is not included. Please see the associated documentation.

### Updating Configuration Settings

Within the `config` directory in the root of this repository there are several `example.yaml` files prefixed with `config` or similar; these hold examples of environment configurations which you should review and potentially override depending on your needs. Anything which is launched out of the `Taskfile` will source it's configuration from these files.

You will need to perform a 1-time action of creating a `.config.yaml` file based on the `.example` files. 
The Taskfiles will also source a `.dotenv` files which match the naming conventions called for `{{.ENV}}` to ease the overriding of environment variables. These files are intentionally added to the `.gitignore` within this repository to prevent you from accidentally committing secrets or other sensitive information which may live inside the server's environment variables.

All settings in the `yaml` configuration can also be overwritten with environment variables prefixed with `DATUM_`. For example, to override the Google `client_secret` set in the yaml configuration with an environment variable you can use: 

```
export DATUM_AUTH_PROVIDERS_GOOGLE_CLIENT_SECRET
```

Configuration precedence is as follows, the latter overriding the former:

1. `default` values set in the config struct within the code
1. `.config.yaml` values
1. Environment variables
