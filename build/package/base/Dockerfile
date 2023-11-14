FROM registry.musicon.space/env AS builder

WORKDIR /source

COPY init init
COPY internal internal
COPY cmd cmd
COPY api api
COPY go.mod go.sum ./

ENV CGO_ENABLED 0
ENV GOOS linux

RUN go build -o bin/user cmd/microservices/user/main.go \
    && go build -o bin/session cmd/microservices/session/main.go \
    && go build -o bin/artist cmd/microservices/artist/main.go \
    && go build -o bin/track cmd/microservices/track/main.go \
    && go build -o bin/playlist cmd/microservices/playlist/main.go \
    && go build -o bin/images cmd/microservices/images/main.go \
    && go build -o bin/album cmd/microservices/album/main.go \
    && go build -o bin/entrypoint cmd/microservices/entrypoint/main.go

FROM alpine:3.17

COPY --from=builder \ 
    /source/bin/user \
    /source/bin/session \
    /source/bin/artist \
    /source/bin/track \
    /source/bin/playlist \
    /source/bin/images \
    /source/bin/album \
    /source/bin/entrypoint \
    /source/bin/
