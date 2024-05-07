
FROM golang:1.20
WORKDIR /app
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
COPY go.mod go.sum ./
RUN go mod download
RUN echo 'PATH=$PATH:/foo/bar' > ~/.env
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-connect
CMD ["/docker-go-connect"]


