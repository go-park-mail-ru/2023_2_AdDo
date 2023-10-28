FROM golang:1.21-alpine AS builder

WORKDIR /source

ENV CGO_ENABLED 0
ENV GOOS linux

COPY . .

RUN go mod download \
    && go mod tidy \
    && go build -o bin/user cmd/microservices/user/main.go \
    && go build -o bin/session cmd/microservices/session/main.go \
    && go build -o bin/artist cmd/microservices/artist/main.go \
    && go build -o bin/track cmd/microservices/track/main.go \
    && go build -o bin/playlist cmd/microservices/playlist/main.go \
    && go build -o bin/album cmd/microservices/album/main.go \
    && go build -o bin/entrypoint cmd/microservices/entrypoint/main.go