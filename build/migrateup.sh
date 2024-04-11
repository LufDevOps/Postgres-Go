#!/bin/bash


echo "Postgres is up - begin migrating"
# Run migration up
migrate -path ./migration -database "postgresql://root:secret@172.19.0.2/simple_bank?sslmode=disable" -verbose up
