# CHANGELOG

## v0.0.1 (2023-12-21)

### Others

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

## v0.0.1 (2023-12-21)

### Others

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
