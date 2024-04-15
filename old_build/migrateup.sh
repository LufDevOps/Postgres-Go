#!/bin/bash

echo "Postgres is up - begin migrating"
# Run migration up

#docker exec -it postgres createdb --username=postgres --owner=postgres postgres

migrate -path ./migration -database "postgresql://$POSTGRESQL_USER:$POSTGRESQL_PASSWORD@$POSTGRESQL_HOST:$POSTGRESQL_PORT/$POSTGRESQL_DB?sslmode=disable" -verbose up
