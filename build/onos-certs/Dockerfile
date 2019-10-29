ARG ONOS_CERTS_BASE_VERSION=latest

FROM onosproject/onos-certs-base:$ONOS_CERTS_BASE_VERSION as base

FROM alpine:3.9
RUN apk add libc6-compat

USER nobody

COPY --from=base /go/src/github.com/onosproject/onos-certs/build/_output/onos-certs /usr/local/bin/onos-certs

ENTRYPOINT ["onos-certs"]