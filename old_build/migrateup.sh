#!/bin/bash

echo "Postgres is up - begin migrating"
# Run migration up

#docker exec -it postgres createdb --username=postgres --owner=postgres postgres

migrate -path ./migration -database "postgresql://postgres:postgrespassword@postgres:5432/shop?sslmode=disable" -verbose up
