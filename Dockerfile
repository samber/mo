
FROM golang:1.22.5-bullseye

WORKDIR /go/src/github.com/samber/mo

COPY Makefile go.* ./

RUN make tools
