FROM alpine:3.17 AS release

COPY --from=registry.musicon.space/base /source/bin/entrypoint /entrypoint

EXPOSE 8080

CMD ["/entrypoint"]
