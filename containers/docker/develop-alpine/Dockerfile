FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 https://github.com/taiyuechain/taipublicchain && \
  (cd taipublicchain && make taipublic) && \
  cp taipublicchain/build/bin/taipublic /taipublic && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /taipublicchain && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/taipublic"]
