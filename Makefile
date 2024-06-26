DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
DOCKER_COMPOSE_FILE ?= docker-compose.dev.yml

network:
	docker network create bank-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
	
dropdb:
	docker exec -it postgres dropdb simple_bank

# migrateup:
# 	migrate -path db/migration -database "$(DB_URL)" -verbose up

# migratedown:
# 	migrate -path db/migration -database "$(DB_URL)" -verbose down

.PHONY:  postgres createdb dropdb migrateup migratedown


migrate:
	sql-migrate up -config=configs/dbconfig.yml