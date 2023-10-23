default: all

all: fmt test build



################################# Generate Functions ########################################

ent:
	@echo "******************** generating ent schema ********************"
	go mod tidy
	go generate ./...
	go mod tidy

graph:
	@echo "******************** generating graph schema ********************"
	bash scripts/gen_graphql.sh

gqlgen:
	@echo "******************** generating gqlgen ********************"
	go run github.com/99designs/gqlgen generate --verbose
	go mod tidy
	go run ./gen_schema.go
	@echo "******************* generating gqlgen client ********************"
	go run github.com/Yamashou/gqlgenc generate --configdir schema

generate: ent graph gqlgen

################################# Database Functions ########################################
atlas: atlas-lint atlas-push

atlas-push:
	atlas migrate push datum --dev-url "sqlite://dev?mode=memory&_fk=1" --dir "file://db/migrations" 

atlas-lint: 
	atlas migrate lint --dev-url "sqlite://file?mode=memory&_fk=1" --dir "file://db/migrations" -w   

################################# Build Functions ########################################
fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

test: 
	$(info ******************** running tests ********************)
	go test -v ./...
	
rover:
	rover dev -u http://localhost:17608/query -s schema.graphql -n datum --elv2-license=accept

run-dev:
	go run main.go serve  --debug --pretty --dev

################################# Template Functions ########################################
setup-template:
	@echo "******************** removing template name occurances ********************"
	bash scripts/clean.sh

clean:
	$(info ******************** removing generated files from repo ********************)
	rm -rf internal/ent/generated/^doc.go
	rm -rf internal/api/^resolver.go
	rm -f schema/ent.graphql
	rm -f schema.graphql
	rm -rf internal/testclient/