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
