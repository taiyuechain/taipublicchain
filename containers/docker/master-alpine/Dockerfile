FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.8 https://github.com/taiyuechain/taipublicchain && \
  (cd taipublicchain && make taipublic) && \
  cp taipublicchain/build/bin/taipublic /taipublic && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /taipublicchain && rm -rf /var/cache/apk/*

EXPOSE 7545
EXPOSE 30513

ENTRYPOINT ["/taipublic"]
