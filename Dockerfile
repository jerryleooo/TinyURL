FROM alpine:3.4

MAINTAINER JerryLeooo <whilgeek@gmail.com>

ADD . /app

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/app/main"]
