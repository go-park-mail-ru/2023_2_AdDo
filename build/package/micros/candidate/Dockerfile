FROM alpine:3.17 AS release

COPY --from=registry.musicon.space/base /source/bin/candidate /candidate
COPY /db/cluster_data/clustering_data.json clustering_data.json

CMD ["/candidate"]
