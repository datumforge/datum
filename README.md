<div align="center">
  <picture>
    <img alt="Datum logo" src="https://github.com/datumforge/datum/raw/main/assets/datumlogo.png" width="50%">
  </picture>
</div>

<br>

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/datumforge/datum)](https://goreportcard.com/report/github.com/datumforge/datum)
[![Build status](https://badge.buildkite.com/a3a38b934ca2bb7fc771e19bc5a986a1452fa2962e4e1c63bf.svg?branch=main)](https://buildkite.com/datum/datum)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=datumforge_datum&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=datumforge_datum)
[![Go Reference](https://pkg.go.dev/badge/github.com/datumforge/datum.svg)](https://pkg.go.dev/github.com/datumforge/datum)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache2.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)

</div>


This repository houses the core server and orchestration elements which are at the heart of the [Datum](https://datum.net) cloud service. We have no plans to ever gate / silo elements of the code that may fall under our "enterprise licensing" (or any other commercial license we offer) and intend to keep the code Apache 2.0 licensed and free for use, forever. Given that, if you find value in anything we're doing here, our cloud services, or use this software yourself (for any purpose) - don't be afraid to become a contributor! If you have any questions please reach out to `contribute@datum.net`.

## What is Datum?

> _datum: a fixed starting point of a scale or operation_

We’re on a mission to build a better, more sustainable world by providing digital leaders with open source solutions which help them innovate faster and remove massive toil. We believe that in order to change the world for good, we’ll benefit from having thousands of companies working at scale, not just a handful (the “hyperscalers”). Check out our [documentation](https://docs.datum.net) or reach out to get involved!

## Features

At it's core, Datum is a collection of services built on top of an entity framework which allows us to:
- Model database schemas as graph structures
- Define schemas as programmatic go code
- Execute complex database queries and graph traversals easily
- Extend and customize using templates and code generation utilities
- Type-safe resolvers and GraphQL schema stitching
- Code generated audit / history tables for defined schemas

On top of this powerful core we also have an incredible amount of pluggable, extensible services:
- Authentication: we today support password, OAuth2 / Social login providers (Github, Google), Passkeys as well as standard OIDC Discovery flows
- Multi-factor: built-in 2FA mechanisms, TOTP
- Authorization: extensible and flexible permissions constructs via openFGA based on Google Zanzibar
- Session Management: built-in session manaagement with JWKS key validation, encrypted cookies and sessions
- Robust Middleware: cache control, CORS, Rate Limiting, transaction rollbacks, and more
- Queuing and Scheduling: Task management and scheduling with Marionette
- External Storage Providers: store data in AWS S3, Google GCS, or locally
- External Database Providers: Leverage Turso, or other PostgreSQL / SQLite compatible vendors and libraries
- Data Isolation and Management: Heirarchal organizations and granular permissions controls


### Hermit

While working on Datum, you can take advantage of the included [Hermit](https://cashapp.github.io/hermit/) dev
environment to get Go & other tooling without having to install them separately on your local machine.

Just use the following command to activate the environment, and you're good to go:

```zsh
. ./bin/activate-hermit
```

### Mac GUI Editors

For other editors and IDEs, the best solution in lieu of native plugins is to
open up a terminal, activate the Hermit environment, then launch the editor
from the terminal. This is not ideal, but does work until a plugin is
available.

1. Close your editor.
2. From a terminal activate your Hermit environment: `. ./bin/activate-hermit`
3. Launch your editor from the terminal:

	| Editor     | Launch command |
	|------------|----------------|
	| [Sublime](https://www.sublimetext.com/docs/3/osx_command_line.html)  | `subl -and .`   |
	| [Visual Studio Code](https://code.visualstudio.com/docs/setup/mac)    | `code .`   |

At this point your editor should be running with environment variables from
the Hermit environment.

## Development

Developing against this repo involves a few mandatory tools; please read up on these and familiarize yourself if you're interested in making additions or changes!

1. [ent](https://entgo.io/) - insane entity mapping tool, definitely not an ORM but kind of an ORM (handles our relational data storage, mappings, codegen processes)
1. [atlas](https://atlasgo.io/) - Schema generation and migrations (can be disabled in lieu of migrations on disk)
1. [goose](https://github.com/pressly/goose) - Secondary database migration utility we also use for seeding data
1. [gqlgen](https://gqlgen.com/) - Code generation + GraphQL server building from from `ent` schema definitions
1. [gqlgenc](https://github.com/Yamashou/gqlgenc) - Client building utilities with GraphQL
1. [openfga](https://openfga.dev/) - Flexible authorization/permission engine inspired by Google Zanzibar
1. [echo](https://echo.labstack.com/) - High performance, extensible, minimalist Go web framework
1. [koanf](https://github.com/knadh/koanf) - Configuration management library which parses command line arguments, Go structs + creates our main configuration files

We also leverage many secondary technologies in use, including (but not limited to!):

1. [taskfile](https://taskfile.dev/usage/) - So much better than Make zomg
1. [redis](https://redis.io/) - in-memory datastore used for sessions, caching
1. [sqlite](https://www.sqlite.org/) - currently planned database system but also offer additional support for PostgreSQL
1. [golangci-lint](https://github.com/golangci/golangci-lint) - an annoyingly opinionated linter
1. [buildkite](https://buildkite.com/datum) - our CI system of choice (with github actions providing some intermediary support)
1. [sonar](https://sonarcloud.io/summary/overall?id=datumforge_datum) - used for code scanning, vulnerability scanning

Lastly we're already ourselves using (and plan to support our customers usage in our cloud service) these third party integrations:

1. [turso/libsql](https://github.com/tursodatabase/libsql) - Turso is an edge-hosted, distributed database that's based on libSQL , an open-source and open-contribution fork of SQLite
1. [posthog](https://posthog.com/) - Product analytics
1. [sendgrid](https://sendgrid.com/en-us) - Transactional email send provider

All of these components are bundled into our respective Docker images; for additional information / instructions, see the [contributing guide](.github/CONTRIBUTING.md) in this repository. We're constantly adding and changing things, but have tried to list all the great open source tools and projects we rely on; if you see your project (or one you use) in here and wish to list it, feel free to open a PR!

## Dependencies

The vast majority of behaviors of the system can be turned on or off by updating the configuration parameters found in `config`; in some instances, we've made features or integrations with third party systems which are "always on", but we're happy to receive PR's wrapping those dependencies if you are interested in running the software without them!

### Installing Dependencies

Setup [Taskfile](https://taskfile.dev/installation/) by following the instructions and using one of the various convenient package managers or installation scripts. After installation, you can then simply run `task install` to load the associated dependencies. Nearly everything in this repository assumes you already have a local golang environment setup so this is not included. Please see the associated documentation.

### Updating Configuration Settings

See the [README](/config/README.md) in the `config` directory.

## Deploying

The only "supported" method of deploying today is locally, but we have a WIP Helm chart which can be found [here](https://github.com/datumforge/helm-charts)

## Contributing

Please read the [contributing](.github/CONTRIBUTING.md) guide as well as the [Developer Certificate of Origin](https://developercertificate.org/). You will be required to sign all commits to the Datum project, so if you're unfamiliar with how to set that up, see [github's documentation](https://docs.github.com/en/authentication/managing-commit-signature-verification/about-commit-signature-verification).

## Licensing

This repository contains `datum` which is open source software under [Apache 2.0](LICENSE). Datum is a product produced from this open source software exclusively by Datum Technology, Inc. This product is produced under our published commercial terms (which are subject to change), and any logos or trademarks in this repository or the broader [datumforge](https://github.com/datumforge) organization are not covered under the Apache License.

Others are allowed to make their own distribution of this software or include this software in other commercial offerings, but cannot use any of the Datum logos, trademarks, cloud services, etc.

## Security

We take the security of our software products and services seriously, including all of the open source code repositories managed through our Github Organizations, such as [datumforge](https://github.com/datumforge). If you believe you have found a security vulnerability in any of our repositories, please report it to us through coordinated disclosure.

**Please do NOT report security vulnerabilities through public github issues, discussions, or pull requests!**

Instead, please send an email to `security@datum.net` with as much information as possible to best help us understand and resolve the issues. See the security policy attached to this repository for more details.

## Questions?

You can email us at `info@datum.net`, open a github issue in this repository, or reach out to [matoszz](https://github.com/matoszz) directly.