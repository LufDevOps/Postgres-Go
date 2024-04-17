FROM golang:1.16-alpine

RUN chmod +x /dist/Postgres-Go_0.1.4-SNAPSHOT-cdf6dba_linux_amd64.tar.gz 
#RUN chmod +x ./dist
COPY dist/Postgres-Go_0.1.4-SNAPSHOT-cdf6dba_linux_amd64.tar.gz /app/

WORKDIR /app

RUN tar -xzf Postgres-Go_0.1.4-SNAPSHOT-cdf6dba_linux_amd64.tar.gz

RUN chmod +x /app/Postgres-Go_0.1.4-SNAPSHOT-cdf6dba_linux_amd64.tar.gz 

# Set the entry point and command
ENTRYPOINT ["./app/Postgres-Go_0.1.4-SNAPSHOT-cdf6dba_linux_amd64.tar.gz"]  
# Add executable permissions to the binary file
