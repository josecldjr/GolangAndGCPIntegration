FROM golang:1.13.12-alpine3.12

WORKDIR /go/src

COPY . /go/src

RUN apk add bash
