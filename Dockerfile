FROM golang:latest

# Debian
RUN mkdir /go/src/app
ADD . /go/src/app