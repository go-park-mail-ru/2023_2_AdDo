FROM alpine:3.17 AS release

COPY --from=2023_2_addo-musicon/base /source/bin/album /album

CMD ["/album"]
