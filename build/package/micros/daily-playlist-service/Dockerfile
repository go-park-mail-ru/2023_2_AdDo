FROM alpine:3.17 AS release

COPY --from=registry.musicon.space/base /source/bin/daily_playlist_service /daily_playlist_service

CMD ["/daily_playlist_service"]
