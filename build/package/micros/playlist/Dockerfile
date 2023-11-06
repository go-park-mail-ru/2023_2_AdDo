FROM alpine:3.17 AS release

COPY --from=musicon-base /source/bin/playlist /playlist

CMD ["/playlist"]
