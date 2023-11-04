FROM alpine:3.17 AS release

COPY --from=registry.musicon.space/base /source/bin/session /session

CMD ["/session"]
