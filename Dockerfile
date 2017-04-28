FROM golang:latest

MAINTAINER JerryLeooo <whilgeek@gmail.com>

ADD . /go/src/tinyUrl

WORKDIR /go/src/tinyUrl

RUN go get github.com/tools/godep

RUN godep restore

RUN go install tinyUrl

ENTRYPOINT ["/go/bin/tinyUrl"]

EXPOSE 8080
