# syntax=docker/dockerfile:1
FROM golang:1.21

WORKDIR /musicon

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o musicon/ ./...

CMD ["./musicon/main"]
