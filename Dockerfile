FROM golang:1.17 AS builder
ENV CGO_ENABLED 0
ADD . /app
WORKDIR /app
RUN go build -ldflags "-s -w" -v -o now main.go

FROM alpine:3
RUN apk update && \
    apk add openssl && \
    rm -rf /var/cache/apk/* \
    && mkdir /app

WORKDIR /app

ADD Dockerfile /Dockerfile

COPY --from=builder /app/now /app/now

RUN chown nobody /app/now \
    && chmod 500 /app/now

USER nobody

ENTRYPOINT ["/app/now"]