ARG ONOS_CERTS_BASE_VERSION=latest

FROM onosproject/onos-certs-base:$ONOS_CERTS_BASE_VERSION as base

FROM golang:1.12.6-alpine3.9 as debugBuilder

RUN apk upgrade --update --no-cache && apk add git && \
    go get -u github.com/go-delve/delve/cmd/dlv

FROM alpine:3.9

RUN apk upgrade --update --no-cache && apk add bash bash-completion libc6-compat

COPY --from=base /go/src/github.com/onosproject/onos-certs/build/_output/onos-certs-debug /usr/local/bin/onos-certs
COPY --from=debugBuilder /go/bin/dlv /usr/local/bin/dlv

RUN echo "#!/bin/sh" >> /usr/local/bin/onos-topo-debug && \
    echo "dlv --listen=:40000 --headless=true --accept-multiclient=true --continue --api-version=2 exec /usr/local/bin/onos-certs -- \"\$@\"" >> /usr/local/bin/onos-certs-debug && \
    chmod +x /usr/local/bin/onos-certs-debug

RUN addgroup -S onos-certs && adduser -S -G onos-certs onos-certs
USER onos-certs
WORKDIR /home/onos-certs

ENTRYPOINT ["onos-certs-debug"]
