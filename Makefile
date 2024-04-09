DB_URL=postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable
DOCKER_COMPOSE_FILE ?= docker-compose.dev.yml
postgres:
	docker run --name postgres-14 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=simple_bank -d postgres:14-alpine

createdb:
	docker exec -it book-keeping-api-main-db-1 createdb --username=root --owner=root root
	
dropdb:
	docker exec -it postgres-14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

# migrate-up: ## Run migrations UP
# migrate-up:
#  docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up

# migrate-down: ## Rollback migrations against non test DB
# migrate-down:
#  docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1

.PHONY:  postgres-14 createdb dropdb migrateup migratedown


migrate:
	sql-migrate up -config=configs/dbconfig.yml