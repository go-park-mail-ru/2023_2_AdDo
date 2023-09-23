# syntax=docker/dockerfile:1.4
FROM golang:1.21 AS builder

WORKDIR /code

ENV CGO_ENABLED 0
ENV GOOS linux

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy


COPY . .

RUN go build -o bin/backend main.go

CMD ["/code/bin/backend"]

EXPOSE 8080
