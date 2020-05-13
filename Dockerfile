# Build Geth in a stock Go builder container
FROM golang:1.10-alpine as construction

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /taipublicchain
RUN cd /taipublicchain && make taipublic

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=construction /taipublicchain/build/bin/taipublic /usr/local/bin/
CMD ["taipublic"]

EXPOSE 7545 7545 9215 9215 30310 30310 30311 30311 30513 30513
ENTRYPOINT ["taipublic"]


