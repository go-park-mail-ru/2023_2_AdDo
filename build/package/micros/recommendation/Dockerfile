FROM alpine:3.17 AS release

COPY --from=registry.musicon.space/base /source/bin/recommendation /recommendation
COPY --from=registry.musicon.space/base /models/classifier.keras classifier.keras

CMD ["/recommendation"]
