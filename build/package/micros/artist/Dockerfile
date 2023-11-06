FROM alpine:3.17 AS release

COPY --from=musicon-base /source/bin/artist /artist

CMD ["/artist"]
