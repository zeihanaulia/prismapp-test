FROM golang:alpine

RUN set -ex \
        && apk add --no-cache git curl vim nano bash

ADD . /go/src/prismapp-test

WORKDIR /go/src/prismapp-test/cmd/realtimed

RUN go get -v
RUN go build

ENTRYPOINT ["./realtimed"]

EXPOSE 7070