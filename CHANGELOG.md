# CHANGELOG

## v0.2.6 (2024-02-03)

### Others

- Update module github.com/mattn/go-sqlite3 to v1.14.22 (#480) (2024-02-02)

- upgrade fgax to fix delete bug (#479) (2024-02-01)

- add and validate email url config (#477) (2024-02-01)

- switch to buildkite atlas plugin (#478) (2024-02-01)

- Update module github.com/Yamashou/gqlgenc to v0.17.0 (#475) (2024-01-31)

- Bump github.com/opencontainers/runc from 1.1.7 to 1.1.12 (#474) (2024-01-31)

- update display_name to no default (#473) (2024-01-31)

- Update anchore/sbom-action action to v0.15.8 (#470) (2024-01-31)

- Update sigstore/cosign-installer action to v3.4.0 (#471) (2024-01-31)

- Update module github.com/mattn/go-sqlite3 to v1.14.21 (#469) (2024-01-31)

- Switch to fgax package (#468) (2024-01-30)

- Organization Invitations (#462) (2024-01-30)

- Update anchore/sbom-action action to v0.15.7 (#467) (2024-01-30)

- Update module github.com/wundergraph/graphql-go-tools to v1.67.1 (#466) (2024-01-29)

- Update anchore/sbom-action action to v0.15.6 (#465) (2024-01-29)

- Update github.com/openfga/language/pkg/go digest to c40760d (#464) (2024-01-29)

- Fix redis name in docker compose (#463) (2024-01-28)

- split up task files, follow style guide (#461) (2024-01-28)

- user test check personal org by id instead of loop (#457) (2024-01-28)

- consistently use tx clients in all requests (#460) (2024-01-26)

- Update openfga/openfga Docker tag to v1.4.3 (#459) (2024-01-26)

- Update otel/opentelemetry-collector Docker tag to v0.93.0 (#458) (2024-01-25)

- Update module github.com/mattn/go-sqlite3 to v1.14.20 (#456) (2024-01-25)

- missing schema change to entitlement tiers for goType (#455) (2024-01-24)

- align enums to all use GoType (#454) (2024-01-23)

- no direct uuid, we use ulids here (#453) (2024-01-23)

- Remove no-auth options from datum server, replace mockgen with mockery (#449) (2024-01-23)

- Update github.com/openfga/language/pkg/go digest to d4c7970 (#451) (2024-01-23)

- Update github.com/openfga/language/pkg/go digest to 314566b (#450) (2024-01-22)

- Update module github.com/openfga/go-sdk to v0.3.4 (#448) (2024-01-22)

- Update github.com/openfga/language/pkg/go digest to d1175b3 (#446) (2024-01-22)

- Update anchore/sbom-action action to v0.15.5 (#447) (2024-01-22)

- Update github.com/openfga/language/pkg/go digest to aaa86ab (#445) (2024-01-22)

- Add Redis Client, Initial cut at session implementation (#441) (2024-01-20)

- Update module github.com/brianvoe/gofakeit/v6 to v6.28.0 (#442) (2024-01-20)

- Update golang.org/x/exp digest to 1b97071 (#440) (2024-01-19)

- Update module github.com/vektah/gqlparser/v2 to v2.5.11 (#436) (2024-01-18)

- Update module github.com/ogen-go/ogen to v0.81.2 (#439) (2024-01-18)

- Update anchore/sbom-action action to v0.15.4 (#437) (2024-01-18)

- Update opentelemetry-go monorepo to v1.22.0 (#435) (2024-01-17)

- First cut at simplifying tuple creation (#429) (2024-01-17)

- drop sessions schema, sessions manager (#434) (2024-01-17)

- Unleash the cookie monster (#432) (2024-01-17)

- fix session schema, hook for issued at, org not required (#431) (2024-01-16)

- Update sonar docker version for main (#433) (2024-01-16)

- use sonar scanner v5 (#427) (2024-01-16)

- Update module github.com/brianvoe/gofakeit/v6 to v6.27.0 (#430) (2024-01-15)

- Update github.com/openfga/language/pkg/go digest to 08a9a21 (#428) (2024-01-15)

- group member cli, resolvers, authz hooks (#421) (2024-01-14)

- Add default org to user settings (#419) (2024-01-14)

- temp sonar fix (#425) (2024-01-14)

- add members as part of org update, org settings for create and update (#418) (2024-01-14)

- set authmw vars to token config (#424) (2024-01-14)

- remove aio configs that are defaults or we might want to change (#422) (2024-01-12)

- Update golang.org/x/exp digest to db7319d (#417) (2024-01-12)

- Adds org member resolvers, cli commands, and hooks  (#388) (2024-01-11)

- add static file handlers (#416) (2024-01-11)

- cli doesn't need the sqlite driver (#410) (2024-01-11)

- fix unique index with soft deletes (#408) (2024-01-11)

- Update module github.com/Yamashou/gqlgenc to v0.16.2 (#406) (2024-01-10)

- set context type for response (#402) (2024-01-10)

- Update openfga/openfga Docker tag to v1.4.2 (#404) (2024-01-10)

- Update golang.org/x/exp digest to 0dcbfd6 (#403) (2024-01-10)

- Update otel/opentelemetry-collector Docker tag to v0.92.0 (#400) (2024-01-10)

- cascade delete through edges, soft delete org + group members (#395) (2024-01-10)

- try not setting these in the aio (#398) (2024-01-10)

- add debug log to ensure testing var (#397) (2024-01-10)

- Update dependency go to v1.21.6 (#396) (2024-01-10)

- remove hooks (#394) (2024-01-10)

- add roles, group and org memberships, org owned mixin (#382) (2024-01-09)

- Update module github.com/brianvoe/gofakeit/v6 to v6.26.4 (#393) (2024-01-09)

- user owned mixin for sesions (#389) (2024-01-09)

- Add display name and avatar url to JWT (#391) (2024-01-09)

- Update github.com/openfga/language/pkg/go digest to a66ff55 (#392) (2024-01-09)

- Update module github.com/lestrrat-go/jwx/v2 to v2.0.19 (#390) (2024-01-08)

- Update module golang.org/x/oauth2 to v0.16.0 (#387) (2024-01-08)

- Update jaegertracing/all-in-one Docker tag to v1.53 (#386) (2024-01-08)

- Update module golang.org/x/crypto to v0.18.0 (#385) (2024-01-08)

- Add password reset policy, user-owned mixiin (#378) (2024-01-08)

- do not add the authorization header more than once (#380) (2024-01-08)

- Update anchore/sbom-action action to v0.15.3 (#384) (2024-01-08)

- Update github.com/openfga/language/pkg/go digest to a678056 (#383) (2024-01-08)

- Add initial support for otel tracing (#381) (2024-01-07)

- add a echo logger (#379) (2024-01-06)

- Ent Privacy rules for Users, Tokens (#352) (2024-01-06)

- Add schemagen hook (#376) (2024-01-06)

- password should not be returned on queries (#375) (2024-01-06)

- fix template mixin (#374) (2024-01-06)

- remove task graph (#372) (2024-01-06)

- Update module github.com/docker/go-connections to v0.5.0 (#371) (2024-01-05)

- Update openfga/openfga Docker tag to v1.4.1 (#365) (2024-01-05)

- auth enabled flag set with new envconfig (#370) (2024-01-05)

- fix missing var in publish all in one (#369) (2024-01-04)

- publish aio image to gcr (#368) (2024-01-04)

- Add initial support for postgres supported databases (#354) (2024-01-04)

- Update module golang.org/x/term to v0.16.0 (#364) (2024-01-04)

- Add Password Reset Handler, Routes, token validation (#359) (2024-01-04)

- Update golang.org/x/exp digest to be819d1 (#361) (2024-01-03)

- Update github.com/openfga/language/pkg/go digest to 720992f (#360) (2024-01-03)

- Update module github.com/ogen-go/ogen to v0.81.1 (#355) (2024-01-02)

- fix naming conventions in docker (#358) (2024-01-02)

- add ToTile case for emails; changes to env files (#357) (2024-01-02)

- Update anchore/sbom-action action to v0.15.2 (#356) (2024-01-02)

- Update graph generation to entc; Drop gen_graph.sh (#332) (2024-01-02)

- switch to envconfig from viper flags for server settings (#349) (2024-01-01)

- Adds session middleware (#346) (2024-01-01)

- Update github.com/openfga/language/pkg/go digest to 53d68ae (#353) (2024-01-01)

- Update ghcr.io/grpc-ecosystem/grpc-health-probe Docker tag to v0.4.24 (#351) (2023-12-31)

- Initial all-in-one image (#350) (2023-12-31)

- fix cascade deletes, add tests for user setting cascade (#347) (2023-12-31)

- Update module github.com/Yamashou/gqlgenc to v0.16.1 (#348) (2023-12-31)

- add transaction middelware for rest endpoints (#344) (2023-12-30)

- limit register, forgot-password, and resend endpoints (#343) (2023-12-30)

- Forgot pass handler (#342) (2023-12-30)

- add resend handler (#340) (2023-12-30)

- add entoas mutation (#339) (2023-12-29)

- add template and basic readme (#334) (2023-12-28)

- Cleanup login tests, update responses in register handler (#338) (2023-12-28)

- Run sonar PR on all PRs (#337) (2023-12-28)

- move marrionatte email send to the right place, add register tests (#335) (2023-12-28)

- Marionettefollowup (#333) (2023-12-28)

- add verify handler, fixup email urls (#330) (2023-12-28)

- Split auth override into sep. compose file (#329) (2023-12-28)

- rollback trasnaction on email failure too (#328) (2023-12-28)

- user last_seen updated on login (#324) (2023-12-28)

- Cleanup compose configuration; Split no-auth/auth into sep. configs; Add compose:datum-auth task (#325) (2023-12-28)

- Update module gocloud.dev to v0.36.0 (#327) (2023-12-28)

- Update golang.org/x/exp digest to 02704c9 (#315) (2023-12-27)

- Add register handler (#321) (2023-12-27)

- Update github.com/openfga/language/pkg/go digest to 41ecd3d (#322) (2023-12-27)

- Update module github.com/prometheus/client_golang to v1.18.0 (#320) (2023-12-27)

- create email verification tokens schemas and generate (#317) (2023-12-27)

- add marionette to server startup (#319) (2023-12-27)

- Add emails package, sendgrid interfaces (#316) (2023-12-27)

- Update github.com/openfga/language/pkg/go digest to f13fb33 (#314) (2023-12-25)

- Add initial sonar scanning (codecoverage and gosec) (#309) (2023-12-25)

- setup v1 routes, middleware for rest (#313) (2023-12-24)

- be more consistent with the logger, use sugared (#312) (2023-12-24)

- pass context for graceful shutdown (#311) (2023-12-23)

- create dev yaml config for auth'ed server (#310) (2023-12-23)

- do not allow personal orgs to have child orgs (#308) (2023-12-22)

- remove codecov config cuz codecov sucks (#307) (2023-12-22)

- Update buildkite plugin docker-login to v3 (#306) (2023-12-22)

- add docker build and publish to bk pipeline  (#304) (2023-12-22)

- set version of base image (#303) (2023-12-22)

- ability to identify personal orgs from a regular org (#294) (2023-12-22)

- update cyclic dep (#301) (2023-12-22)

- fix trasnaction tests (#299) (2023-12-22)

- missed schema change from PR 296 (#298) (2023-12-22)

- respect debug level flag for logging (#297) (2023-12-22)

- Add cascade deletes to user edges (#296) (2023-12-22)

- Update module google.golang.org/protobuf to v1.32.0 (#295) (2023-12-22)

- Update module go.uber.org/mock to v0.4.0 (#286) (2023-12-21)

- Update module github.com/openfga/go-sdk to v0.3.3 (#292) (2023-12-21)

- Test to confirm Transactions roll back on Authz Tuple write errors (#291) (2023-12-21)

- add secure, cachecontrol, middleware (#288) (2023-12-21)

- remove base mixin (#289) (2023-12-21)

- Codecov, CHANGELOG.md, release automation (#290) (2023-12-21)

- remove TODO comment on rollback (#284) (2023-12-20)

- switch mutations to transacationer client from context (#283) (2023-12-20)

- Update github.com/openfga/language/pkg/go digest to 7cb4a2c (#281) (2023-12-20)

- Update module github.com/openfga/go-sdk to v0.3.2 (#280) (2023-12-20)

- remove entviz (#279) (2023-12-19)

- Update module github.com/openfga/go-sdk to v0.3.1 (#274) (2023-12-19)

- add authz group tests (#273) (2023-12-19)

- Generate sanity check (#270) (2023-12-19)

- Run tidy after updates (#272) (2023-12-19)

- Update golang.org/x/exp digest to dc181d7 (#271) (2023-12-19)

- Update module github.com/spf13/viper to v1.18.2 (#263) (2023-12-19)

- Update module golang.org/x/crypto to v0.17.0 [SECURITY] (#269) (2023-12-19)

- Groupauthz (#266) (2023-12-18)

- fix length access token works (#260) (2023-12-18)

- add org setting get (#265) (2023-12-18)

- ensure user has write access to parent when creating child org  (#259) (2023-12-18)

- fix inifite loop when child orgs are requested (#255) (2023-12-18)

- Update module github.com/ogen-go/ogen to v0.81.0 (#257) (2023-12-18)

- missing schema changes from cascade delete changes (#256) (2023-12-17)

- Add initial edgecleanup helper functions (#251) (2023-12-17)

- email confirmed true; return user id on creation (#253) (2023-12-16)

- run forks gosec upload on pull (#252) (2023-12-16)

- Adds the ability to get groups (#250) (2023-12-15)

- Refresh access tokens when using the cli (#249) (2023-12-15)

- Adds the refresh endpoint (#248) (2023-12-15)

- Update module github.com/mattn/go-sqlite3 to v1.14.19 (#247) (2023-12-14)

- move helpers to hooks (#246) (2023-12-14)

- Hooks for display names to be set on orgs, users, groups (#245) (2023-12-14)

- auto create org settings on org creation  (#239) (2023-12-14)

- add fga dep to run-dev-auth (#238) (2023-12-14)

- Update golang.org/x/exp digest to aacd6d4 (#237) (2023-12-14)

- add gosec to buildkite, add go build-cli, move to groups (#169) (2023-12-14)

- Add CLI login with username/password auth (#235) (2023-12-14)

- Update module github.com/brianvoe/gofakeit/v6 to v6.26.3 (#233) (2023-12-13)

- Update github.com/openfga/language/pkg/go digest to cca4c43 (#234) (2023-12-13)

- refactor cli layout, add self flag on user command (#231) (2023-12-13)

- set durations for tokens (#229) (2023-12-13)

- upgrade openfga v0.3.0 (#226) (2023-12-13)

- Update github/codeql-action action to v3 (#228) (2023-12-13)

- Update module github.com/brianvoe/gofakeit/v6 to v6.26.2 (#227) (2023-12-12)

- Adds login handler (#225) (2023-12-12)

- Add tests for derived keys (#222) (2023-12-12)

- Update module github.com/google/uuid to v1.5.0 (#224) (2023-12-12)

- fix the db client nil pointer (#220) (2023-12-12)

- pass ent db to the handlers (#219) (2023-12-11)

- validate password strength before creating/updating user (#218) (2023-12-11)

- Marionette (#211) (2023-12-11)

- add more fields for the user query (#217) (2023-12-11)

- url tokens (#210) (2023-12-11)

- Update github.com/openfga/language/pkg/go digest to 8dfc3b8 (#216) (2023-12-11)

- Switch ids to ULIDS instead of nano ids (#214) (2023-12-11)

- fix error response to return errors properly (#215) (2023-12-11)

- Update github.com/openfga/language/pkg/go digest to 779e682 (#212) (2023-12-11)

- Update sigstore/cosign-installer action to v3.3.0 (#213) (2023-12-11)

- Add auth middleware (#204) (2023-12-11)

- Update module gocloud.dev to v0.35.0 (#207) (2023-12-09)

- run migrate even if linter failed on main (#209) (2023-12-08)

- add and register routes with not implemented (#208) (2023-12-08)

- schema diff on main (#205) (2023-12-08)

- Tokenmanager (#198) (2023-12-08)

- Update module gocloud.dev to v0.34.0 (#203) (2023-12-08)

- user hook, avatar, pass (#202) (2023-12-08)

- Update module github.com/ogen-go/ogen to v0.80.1 (#200) (2023-12-08)

- Update module github.com/spf13/viper to v1.18.1 (#201) (2023-12-08)

- add graph resolver using serveropts (#199) (2023-12-07)

- adds db healthcheck, moves to server opts (#197) (2023-12-07)

- error check cert files and panic when not found (#196) (2023-12-06)

- rename oidc flag to auth (#195) (2023-12-06)

- Update module github.com/spf13/viper to v1.18.0 (#190) (2023-12-06)

- Update github.com/datumforge/echo-jwt/v5 digest to 63228bd (#192) (2023-12-06)

- Update github.com/datumforge/echox digest to eb30d6b (#193) (2023-12-06)

- Update dependency go to v1.21.5 (#187) (2023-12-06)

- Update actions/setup-go action to v5 (#191) (2023-12-06)

- Revamp server setup with new httpserve packages, echo v5 (#175) (2023-12-06)

- Update github.com/openfga/language/pkg/go digest to 92fa8fb (#188) (2023-12-05)

- Update module github.com/brianvoe/gofakeit/v6 to v6.26.0 (#189) (2023-12-05)

- rename api package to graphapi (#186) (2023-12-05)

- take a cut at access tokens (#184) (2023-12-04)

- Groups (#179) (2023-12-04)

- fix labeler configuration for v5 (#183) (2023-12-04)

- Update actions/labeler action to v5 (#182) (2023-12-04)

- Update anchore/sbom-action action to v0.15.1 (#181) (2023-12-04)

- Update github.com/openfga/language/pkg/go digest to 50a2774 (#180) (2023-12-04)

- cli build and clean (#178) (2023-12-03)

- Update module github.com/99designs/gqlgen to v0.17.41 (#176) (2023-12-03)

- Update module github.com/golang-jwt/jwt/v5 to v5.2.0 (#173) (2023-12-02)

- Update module github.com/ogen-go/ogen to v0.79.1 (#174) (2023-12-02)

- allow org creation when oidc=false (#172) (2023-12-01)

- fix get all orgs with no auth (#170) (2023-12-01)

- Authz checks for org hierarchy - parent (#154) (2023-12-01)

- delete relationship tuples on soft delete (#166) (2023-12-01)

- permission denied per type and action (#167) (2023-11-30)

- allow org names to be reused if soft-deleted (#164) (2023-11-30)

- do not push the atlas migration on task:pr, this should happen in CI on merge to main (#165) (2023-11-30)

- Initial soft delete (#157) (2023-11-30)

- add basic caching, using entcache, to the db layer (#156) (2023-11-30)

- Aligns audit mixin values with others when no auth is used; Fixes bug with retrieving user when auth is not enabled (#158) (2023-11-30)

- Sets up basic user creation, personal orgs, and cli commands (#146) (2023-11-28)

- viperconfig and basic cleanup (#147) (2023-11-28)

- Update fga playground task command in README (#148) (2023-11-28)

- Update module golang.org/x/crypto to v0.16.0 (#144) (2023-11-28)

- add cookie and session store (#145) (2023-11-27)

- Update github.com/openfga/language/pkg/go digest to 50a8baa (#143) (2023-11-27)

- Update module github.com/brianvoe/gofakeit/v6 to v6.25.0 (#141) (2023-11-27)

- Update github.com/openfga/language/pkg/go digest to 9d2548a (#142) (2023-11-27)

- fix spelling typo on org settings schema (#139) (2023-11-26)

- update template command and add http client (#140) (2023-11-26)

- Cleanup getting user information in audit mixin (#138) (2023-11-26)

- move hooks to its own package (#135) (2023-11-26)

- add mockgen (#134) (2023-11-26)

- Update auditmixin to set createdby and updatedby; Set createdby to immutable so it can't be updated after the fact (#133) (2023-11-26)

- add passwd package (#130) (2023-11-25)

- add keygen package (#131) (2023-11-25)

- add utils package (#132) (2023-11-25)

- Adds ent interceptor to log query duration (#129) (2023-11-25)

- Update module github.com/prometheus/client_golang to v1.17.0 (#126) (2023-11-25)

- Add scaffolding for initial Prometheus metrics (#125) (2023-11-25)

- use passed context, not background (#124) (2023-11-25)

- Adding authz with openfga (#93) (2023-11-24)

- revert (#100) (2023-11-22)

- stub out login / register (2023-11-22)

- add readyz and livez (#98) (2023-11-21)

- add user sub (#97) (2023-11-21)

- fix case and pluralism mismatches (#96) (2023-11-20)

- fix naming of some queries and mutations (#95) (2023-11-20)

- Update anchore/sbom-action action to v0.15.0 (#94) (2023-11-20)

- Add tests for echox.GetActorSubject (#92) (2023-11-17)

- Adds a basic FGA model with organization, groups, subscriptions, and features (#91) (2023-11-17)

- version and goreleaser (#90) (2023-11-17)

- Update module github.com/ogen-go/ogen to v0.78.0 (#88) (2023-11-16)

- add descriptions to taskfiles (#86) (2023-11-14)

- update task file; minor docs (#85) (2023-11-14)

- Fix Docker-Compose for FGA (#84) (2023-11-14)

- add additional edges and create migrations (#83) (2023-11-14)

- ent v0.12.4 -> v0.12.5, run generate (#82) (2023-11-13)

- make display name test number of letters to ensure no spaces (#81) (2023-11-13)

- add cli with org CRUD operations (#80) (2023-11-13)

- update org tests to account for new fields, unique test (#78) (2023-11-11)

- User settings (#77) (2023-11-10)

- Add test utils and organization crud resolver tests (#76) (2023-11-10)

- add oauth provider (#75) (2023-11-10)

- add pat (#74) (2023-11-09)

- add entitlements (#71) (2023-11-09)

- organization setting (#70) (2023-11-09)

- Update module github.com/Yamashou/gqlgenc to v0.16.0 (#69) (2023-11-09)

- remove ogent, update scopes to strings array (#68) (2023-11-08)

- Update module golang.org/x/crypto to v0.15.0 (#67) (2023-11-08)

- add refresh token (#64) (2023-11-08)

- add  organization queries and mutations for generated client (#66) (2023-11-08)

- Update module github.com/golang-jwt/jwt/v5 to v5.1.0 (#65) (2023-11-08)

- Upgrade images, set GOTOOLCHAIN=auto (#63) (2023-11-07)

- update golang versions (#61) (2023-11-07)

- Update dependency go to v1.21.4 (#60) (2023-11-07)

- Bump google.golang.org/grpc from 1.58.2 to 1.58.3 (#59) (2023-11-07)

- Update module github.com/labstack/echo/v4 to v4.11.3 (#57) (2023-11-07)

- Update dependency go to v1.21.3 (#55) (2023-11-07)

- add secrets keeper (#58) (2023-11-07)

- Adding TLS Config  (#46) (2023-11-06)

- Update module github.com/go-faster/errors to v0.7.0 (#54) (2023-11-06)

- Use nanox.ID over UUID, but as a string (#51) (2023-11-06)

- Update module github.com/mattn/go-sqlite3 to v1.14.18 (#52) (2023-11-05)

- Update module github.com/spf13/cobra to v1.8.0 (#53) (2023-11-05)

- Gosec workflow (#49) (2023-11-03)

- Add new id based on nanoid (#48) (2023-11-03)

- add templates directory (#44) (2023-11-01)

- Upgrade to golang-jwt/jwt/v5 from v4 (#42) (2023-11-01)

- Remove memberships; make organization hierarchal (#41) (2023-10-31)

- adds pagination, sorting (#40) (2023-10-31)

- add datumclient (#39) (2023-10-30)

- Adds the ability to write to two databases  (#38) (2023-10-29)

- audit should set uuid, not int (#37) (2023-10-29)

- Adds echo-jwt middleware (#32) (2023-10-29)

- Update postgres Docker tag to v16 (#33) (2023-10-29)

- Changes created_by and updated_by to UUIDs, adds custom scaler (#34) (2023-10-29)

- Add privacy (#26) (2023-10-29)

- add atlas.hcl, schema.hcl (#31) (2023-10-29)

- add labeler action (#30) (2023-10-29)

- add ent features for privacy and interceptors (#29) (2023-10-28)

- Update actions/checkout action to v4 (#27) (2023-10-28)

- .github/workflows: add atlas ci workflow (#25) (2023-10-28)

- add excludes logic to graphql generation (#24) (2023-10-27)

- add globaluniqueID (#23) (2023-10-27)

- Groups (#21) (2023-10-27)

- Update module github.com/ogen-go/ogen to v0.77.0 (#22) (2023-10-27)

- Update module github.com/google/uuid to v1.4.0 (#20) (2023-10-27)

- Update module github.com/99designs/gqlgen to v0.17.40 (#18) (2023-10-24)

- Atlas - use atlas config instead of goose (#17) (2023-10-23)

- add validation for org name length, return validation + constraint errors (#16) (2023-10-23)

- Add rover setup for apollo sandbox  (#15) (2023-10-21)

- Adds the basic operations to CRUD resolvers (#13) (2023-10-20)

- ID should be a UUID, not a string (#12) (2023-10-20)

- Add migrations, updates to generated code (#11) (2023-10-20)

- Blowout user schema; add sessions (#10) (2023-10-20)

- Add DB connection, make user schema consistent (#9) (2023-10-18)

- fix linter failures (#8) (2023-10-18)

- Fix go mod openapi (#6) (2023-10-18)

- update entc.go (#5) (2023-10-18)

- add additional graphql schemas, gqlgen, working server (#4) (2023-10-18)

- port over go-template changes (#3) (2023-10-18)

- Template cleanup + base org / member / user (#2) (2023-10-18)

- Initial commit (2023-10-17)
