FROM golang:1.21-alpine AS builder

WORKDIR /source

ENV CGO_ENABLED 0
ENV GOOS linux

COPY cmd ./cmd
COPY internal/ ./internal
COPY init ./init
COPY go.sum go.mod ./

RUN go mod download \
    && go mod tidy \
    && go build -o bin/artist cmd/microservices/artist/main.go

FROM alpine:3.17 AS release

COPY --from=builder /source/bin/artist /artist

CMD ["/artist"]
