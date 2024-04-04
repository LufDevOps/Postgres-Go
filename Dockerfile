

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
COPY go.mod go.sum ./

RUN go mod download
RUN echo 'PATH=$PATH:/foo/bar' > ~/.env

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-connect

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Rungo install golang.org/x/tools/gopls@latest
CMD ["/docker-go-connect"]
ENTRYPOINT [ "/bin/bash","-c" ]