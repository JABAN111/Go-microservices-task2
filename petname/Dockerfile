FROM golang:1.23 AS build

ENV PATH="$PATH:$(go env GOPATH)/bin"


COPY . /src/
COPY vendor /src/vendor
COPY proto /src/proto

WORKDIR /src

WORKDIR /src/cmd/apiserver

ENV CGO_ENABLED=0
RUN go build -mod=vendor -o /petname server.go

FROM alpine:3.20

COPY --from=build /petname /petname
ENV PETNAME_GRPC_PORT=8085

ENTRYPOINT [ "/petname" ]