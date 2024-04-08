FROM ubuntu

WORKDIR /
COPY db/migration ./migrations
COPY build/migrateup.sh ./

RUN apt-get -y update; apt-get -y install curl
RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -

RUN echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
RUN migrate -version
RUN chmod +x ./migrateup.sh

ENTRYPOINT ["./migrateup.sh"]

