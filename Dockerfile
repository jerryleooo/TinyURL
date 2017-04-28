FROM golang:latest

MAINTAINER JerryLeooo <whilgeek@gmail.com>

ADD . /go/src/tinyUrl

WORKDIR /go/src/tinyUrl

RUN go get github.com/tools/godep

RUN godep restore

RUN go install tinyUrl

ENTRYPOINT ["/go/bin/tinyUrl"]

RUN echo $GR34AB0E8080_HOST
RUN echo $GR34AB0E8080_PORT

EXPOSE 8080
