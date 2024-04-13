#!/usr/bin/env sh

set -e

# Assuming POSTGRESQL_PASSWORD is set to your actual password
# Use a command-line tool like jq to URL-encode the password. Ensure jq is installed.
ENCODED_PASSWORD=$(jq -nr --arg pass "$POSTGRESQL_PASSWORD" '$pass|@uri')

until psql postgres://${POSTGRESQL_USERNAME}:${ENCODED_PASSWORD}@${POSTGRESQL_HOST}:${POSTGRESQL_PORT}/${POSTGRESQL_DATABASE} -c '\q'; do
  echo "Postgres is unavailable - waiting..."
  sleep 1
done

echo "Postgres is up - begin migrating"
sql-migrate up -config=dbconfig.yml -env="production"
