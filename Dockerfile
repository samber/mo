
FROM golang:1.23.0-bullseye

WORKDIR /go/src/github.com/samber/mo

COPY Makefile go.* ./

RUN make tools
