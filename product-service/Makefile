# ----------------------------
# Env Variables
# ----------------------------
DOCKER_COMPOSE_FILE ?= docker-compose.yaml
DATABASE_CONTAINER ?= database

## up: starts all related containers to run the service
up:	api-up

## api-setup: starts all backend supported containers
api-up: pg-create

## pg-create: starts postgres container
pg-create:
	@echo Starting Postgres database container in background
	docker-compose -p="product-service" up -d ${DATABASE_CONTAINER}
	@echo Database image started!

## down: stop and remove all docker containers
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!
