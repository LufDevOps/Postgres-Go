#!/bin/bash

echo $POSTGRES_USER


# Run migration up
migrate -path ./migrations -database "postgresql://$POSTGRES_USER:$POSTGRES_DBPASSWORD@localhost:5433/$POSTGRES_DBNAME?sslmode=disable" -verbose up
