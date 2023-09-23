FROM golang:1.21-alpine AS builder

WORKDIR /code

ENV CGO_ENABLED 0
ENV GOOS linux

COPY . .

RUN go mod download \
    && go mod tidy \
    && go build -o bin/backend main.go

FROM alpine:3.17 AS release

COPY --from=builder /code/bin/backend /backend 
COPY --from=builder /code/config.ini /config.ini

EXPOSE 8080

CMD ["/backend"]
