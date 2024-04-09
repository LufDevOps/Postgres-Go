#!/bin/bash

# Load environment variables from .env file
set -a
source "$(dirname "$0")/.env"
set +a

# Run migration down
migrate -path db/migration -database "postgresql://$POSTGRES_USER:$POSTGRES_DBPASSWORD@localhost:5432/$POSTGRES_DBNAME?sslmode=disable" -verbose down
