# ----------------------------
# Env Variables
# ----------------------------
DOCKER_COMPOSE_FILE ?= docker-compose.yaml
DATABASE_CONTAINER ?= database

## run: starts containers to run api server
run: pg-create api-create

## api-setup: executes pre-defined steps to setup api server
api-setup: pg-create pg-migrate api-gen-models api-go-generate

## api-create: starts api server
api-create:
	@echo Starting Api container in background
	docker-compose -p="product-service" up ${API_CONTAINER}
	@echo Api container started!

## api-gen-models: executes CLI command to generate new database models
api-gen-models:
	@echo Starting generate db model...
	cd api && sh -c 'sqlboiler psql'
	@echo Done!

## api-gen-models: executes CLI command to generate Mock files
api-gen-mocks:
	@echo Starting generate Mock files...
	cd api && sh -c '\
		mockery --dir internal/controller --all --recursive --inpackage && \
		mockery --dir internal/repository --all --recursive --inpackage'
	@echo Done!

## api-test: executes all test cases
api-test:
	cd api; \
	env $$(cat ./local.env | egrep -v '^#' | xargs) \
    sh -c 'go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./... | grep -v pkg'

## api-go-generate: executes CLI command to generate go based config to handle graphql
api-go-generate:
	@echo Starting generate gqlgenerated file...
	cd api && sh -c 'go generate ./...'
	@echo Done!

## pg-create: starts postgres container
pg-create:
	@echo Starting Postgres database container in background
	docker-compose -p="product-service" up -d ${DATABASE_CONTAINER}
	@echo Database container started!

## new-migration-file: creates a DB migration files. Ex: make new-migration-file name=init
new-migration-file:
	docker-compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migrations $(name)

## pg-migrate: executes latest migration files
pg-migrate:
	@echo Running migration
	docker-compose --profile tools run --rm migrate up
	@echo Migration done!

## pg-drop: reset db to blank
pg-drop:
	@echo Dropping database...
	docker-compose --profile tools run --rm migrate drop
	@echo Done!

## pg-redo: rollbacks migrations against non test DB
pg-redo:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down

## down: stop and remove all docker containers
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!
