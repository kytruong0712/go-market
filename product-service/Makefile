# ----------------------------
# Env Variables
# ----------------------------
DOCKER_COMPOSE_FILE ?= docker-compose.yaml
DATABASE_CONTAINER ?= database

## run: starts containers to run api server
run: pg-create api-create

## api-setup: executes pre-defined steps to setup api server
api-setup: pg-create pg-migrate api-gen-models

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
	docker compose --profile tools run --rm migrate drop

## api-create: starts api server
api-create:
	@echo Starting Api container in background
	docker compose -p="product-service" up ${API_CONTAINER}
	@echo Api container started!

## api-gen-models: executes CLI command to generator new database models
api-gen-models:
	@echo Starting generate db model...
	cd api && sh -c 'sqlboiler psql'
	@echo Done!

## down: stop and remove all docker containers
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!
