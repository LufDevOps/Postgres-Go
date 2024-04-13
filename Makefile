DB_URL=postgresql://postgres:secret@localhost:5432/people?sslmode=disable
DOCKER_COMPOSE_FILE ?= docker-compose.dev.yml

network:
	docker network create people-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank
	
dropdb:
	docker exec -it postgres dropdb people

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

# migrate-up: ## Run migrations UP
# migrate-up:
#  docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up

# migrate-down: ## Rollback migrations against non test DB
# migrate-down:
#  docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1

.PHONY:  postgres createdb dropdb migrateup migratedown


migrate:
	sql-migrate up -config=configs/dbconfig.yml