FROM alpine:3.17 AS release

COPY --from=registry.musicon.space/base /source/bin/daily_playlist_worker /daily_playlist_worker

CMD ["/daily_playlist_worker"]
