FROM alpine:3.17 AS release

COPY --from=musicon-base /source/bin/entrypoint /entrypoint

EXPOSE 8080

CMD ["/entrypoint"]
