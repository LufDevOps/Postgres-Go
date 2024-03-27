FROM golang:1.16-alpine3.13

WORKDIR /app

COPY . .

RUN go build -o main main.go

EXPOSE 3000

CMD [ "/app/main" ]
