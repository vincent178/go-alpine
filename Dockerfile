FROM golang:alpine3.10 AS build_base

ARG HTTPS_PROXY
ARG HTTP_PROXY

ENV GO111MODULE=on
WORKDIR /src
COPY go.mod .

RUN go mod download

FROM build_base

RUN apk add --no-cache iproute2 ipset

ARG HTTPS_PROXY
ARG HTTP_PROXY

COPY . /src
WORKDIR /src
RUN go build -o app
CMD ['./app']