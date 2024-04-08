postgres:
	docker run --name postgres-14 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres-14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-14 dropdb simple_bank

.PHONY:  postgres-14 createdb dropdb 