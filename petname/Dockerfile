FROM golang:1.23 AS build

ENV PATH="$PATH:$(go env GOPATH)/bin"
MAINTAINER Jaba "368601@edu.itmo.ru"

COPY . /src/
COPY proto /src/proto

WORKDIR /src/cmd/apiserver

ENV CGO_ENABLED=0
#RUN go build -mod=vendor -o /petname server.go
RUN go build -o /petname server.go

FROM alpine:3.20

COPY --from=build /petname /petname

ENTRYPOINT [ "/petname" ]