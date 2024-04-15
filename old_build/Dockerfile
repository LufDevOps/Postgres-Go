FROM ubuntu

WORKDIR /app

COPY db/migration ./migration
COPY old_build/migrateup.sh ./

RUN apt-get -y update; apt-get -y install curl
RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update
RUN apt-get install migrate
RUN chmod +x ./migrateup.sh

# ENV POSTGRESQL_HOST="host.docker.internal"
# ENV POSTGRESQL_PORT=5432
# ENV POSTGRESQL_USER="root"
# ENV POSTGRESQL_PASSWORD="secret"
# ENV POSTGRESQL_DB="simple_bank"

ENV POSTGRESQL_HOST="postgres-giang-postgresql.postgresql.svc.cluster.local"
ENV POSTGRESQL_PORT=5432
ENV POSTGRESQL_USER="postgres"
ENV POSTGRESQL_PASSWORD="dz1hXCSgxg"
ENV POSTGRESQL_DB="postgres"

# CMD [ "ls" ]


ENTRYPOINT ["./migrateup.sh"]

