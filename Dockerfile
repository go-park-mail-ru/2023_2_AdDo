FROM golang:1.21-alpine AS builder

WORKDIR /source

ENV CGO_ENABLED 0
ENV GOOS linux

COPY cmd ./cmd
COPY internal ./internal
COPY init ./init
COPY go.mod go.sum ./

RUN go mod download \
    && go mod tidy \
    && go build -o bin/musicon cmd/musicon/main.go

FROM alpine:3.17 AS release

COPY --from=builder /source/bin/musicon /musicon

EXPOSE 8080

CMD ["/musicon"]
