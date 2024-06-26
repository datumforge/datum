## Pre-commit Scripts

These scripts are used primarily to check our code before
committing.

| Script Name                   | Description                                                       |
| ----------------------------- | ----------------------------------------------------------------- |
| `commit-msg`                  | Ensure JIRA issue is tagged to commit message                     |
| `gen-docs-index`              | generate index for documents                                      |
| `pre-commit-go-custom-linter` | run a custom linter against files (passed go files by pre-commit) |
| `pre-commit-go-imports`       | modify imports in go files                                        |
| `pre-commit-go-lint`          | modify go files with linting rules                                |
| `pre-commit-go-mod`           | modify `go.mod` and `go.sum` to match whats in the project        |
| `pre-commit-go-vet`           | analyze code with `go vet`                                        |
| `pre-commit-swagger-validate` | run `swagger validate`                                            |
| `lint-yaml-with-spectral`     | run `spectral` linter on external APIs                            |