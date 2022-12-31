# Include variable from the .envrc file
# the envrc looks like below:
# export GREENLIGHT_DB_DSN='postgres://greenlight:passwor@localhost:5432/greenlight?sslmode=disable'

include .envrc

# =============================================== #
# HELPERS
# =============================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N]' && read ans && [ $${ans:-N} = y ]

# =============================================== #
# DEVELOPMENT
# =============================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${GREENLIGHT_DB_DSN}

## db/migration/up: apply all up database migrations
.PHONY: db/migration/up
db/migration/up: confirm
	@echo "Running up migrations..."
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up

## db/migration/new name=$1: create a new database migration
.PHONY: db/migration/new
db/migration/new:
	@echo "Create migration files for ${name}..."
	migrate create -seq -ext=.sql -dir=./migrations ${name}

# =============================================== #
# QUALITY CONTROL
# =============================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo '=== Tidying and verifying module dependencies ==='
	go mod tidy
	go mod verify
	@echo '=== Formatting code ==='
	go fmt ./...
	@echo '=== Vetting code ==='
	go vet ./...
	staticcheck ./...
	@echo '=== Running tests ==='
	go test -race -vet=off ./...
