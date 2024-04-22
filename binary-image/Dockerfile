FROM golang:1.16-alpine

# RUN chmod +x dist/Postgres-Go_0.1.4-SNAPSHOT-5506cf3_linux_arm64.tar.gz

COPY dist/Postgres-Go_0.1.4-SNAPSHOT-cdf6dba_linux_amd64.tar.gz /app/

WORKDIR /app
RUN chmod +x /app
RUN tar -xzf /app/Postgres-Go_0.1.4-SNAPSHOT-cdf6dba_linux_amd64.tar.gz

# RUN cp dist/Postgres-Go_0.1.4-SNAPSHOT-5506cf3_linux_arm64.tar.gz . 
# RUN tar -xzf Postgres-Go_0.1.4-SNAPSHOT-5506cf3_linux_arm64.tar.gz 

RUN chmod +x /app/Postgres-Go

#RUN ./Postgres-Go

ENTRYPOINT ["./Postgres-Go"]  

