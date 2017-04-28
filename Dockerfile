FROM golang:latest

MAINTAINER JerryLeooo <whilgeek@gmail.com>

ADD . /app

WORKDIR /app

EXPOSE 8080

RUN go build main.go

ENTRYPOINT ["/app/main"]
