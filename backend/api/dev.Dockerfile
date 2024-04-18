FROM golang:1.22.0-alpine as builder

ENV GO111MODULE=on

RUN apk update \
  && apk upgrade \
  && apk add --no-cache \
  make build-base libtool musl-dev ca-certificates dumb-init curl \
  && update-ca-certificates 2>/dev/null || true

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app
RUN go mod tidy

CMD air -d

EXPOSE 9000