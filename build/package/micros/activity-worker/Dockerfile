FROM alpine:3.17 AS release

COPY --from=registry.musicon.space/base /source/bin/activity_worker /activity_worker

CMD ["/activity_worker"]
