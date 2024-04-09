#!/bin/bash

echo $POSTGRES_DBNAME
# Run migration up
migrate -path ./migration -database "postgresql://$POSTGRES_USER:$POSTGRES_DBPASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DBNAME?sslmode=disable" -verbose up
#migrate -path ./migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
