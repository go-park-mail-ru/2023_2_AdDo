FROM alpine:3.17 AS release

COPY --from=registry.musicon.space/base /source/bin/mailer /mailer

COPY --from=registry.musicon.space/base /source/internal/microservices/mailer/service/server/templates/ /templates

CMD ["/mailer"]
