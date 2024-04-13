FROM ubuntu

WORKDIR /app

COPY db/migration ./migration
COPY build/migrateup.sh ./

RUN apt-get -y update; apt-get -y install curl
RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update
RUN apt-get install migrate
RUN chmod +x ./migrateup.sh

# ENV POSTGRES_USER=root
# ENV POSTGRES_PASSWORD=secret
# ENV POSTGRES_HOST=0.0.0.0
# ENV POSTGRES_DB=simple_bank
# ENV POSTGRES_PORT=5433

# CMD [ "ls" ]


ENTRYPOINT ["./migrateup.sh"]

