# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin/server cmd/server.go

CMD [ "/app/bin/server" ]
