FROM golang:1.21-alpine AS builder

WORKDIR /source

COPY init init
COPY internal internal
COPY cmd cmd
COPY api api
COPY go.mod go.sum ./

RUN go mod download && go mod tidy
