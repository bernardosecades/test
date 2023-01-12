all: help

#### LINTER ##
## lint: Run linter. Usage: 'make lint'
lint: ; $(info Linting ...)
	golangci-lint run --fix

#### TEST ACTIONS ##
## test-unit: Run all unit test. Usage: 'make test-unit'
test-unit:
	docker-compose exec service bash -c "go clean -testcache  && go test ./... -tags=unit -v"
	#go test -p 1 ./... -tags=unit -v

## test-e2e: Run all end to end test. Usage: 'make test-e2e'
test-e2e:
	docker-compose exec service bash -c "go clean -testcache  && go test ./... -tags=e2e -v"
	#go test -p 1 ./... -tags=e2e -v

## test-all: Run all test. Usage: 'make test-all'
test-all:
	docker-compose exec service bash -c "go clean -testcache  && go test ./... -tags=unit,e2e"

#### HELP ##
## help: Show this screen
help: Makefile
	@sed -n 's/^##//p' $<