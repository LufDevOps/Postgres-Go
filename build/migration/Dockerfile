FROM golang:1.21-alpine
COPY ./build/migration/migrate.sh ./
RUN apk add --no-cache postgresql-client jq && \
apk add --no-cache git && \
CGO_ENABLED=0 go install github.com/rubenv/sql-migrate/...@latest && \
chmod +x migrate.sh
COPY ./build/migration/dbconfig.yml ./
COPY ./migrations ./migrations

ENV POSTGRES_HOST="localhost"
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER="postgres"
# POSTGRES_DBPASSWORD="dz1hXCSgxg"
ENV POSTGRES_PASSWORD="secret"
ENV POSTGRES_DB="people"

ENTRYPOINT ["./migrate.sh"]
