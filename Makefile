postgres:
	docker run -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=simple_bank -d postgres:14-alpine

createdb:
	docker exec -it postgres-14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down


.PHONY:  postgres-14 createdb dropdb migrateup migratedown