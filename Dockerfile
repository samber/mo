
FROM golang:1.23.3-bullseye

WORKDIR /go/src/github.com/samber/mo

COPY Makefile go.* ./

RUN make tools
