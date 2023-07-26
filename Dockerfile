FROM golang:1.18-alpine AS builder

# RUN chmod +x /entrypoint.sh

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add bash

RUN apk add git
RUN apk add aws-cli

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o greenhouse_api ./src

FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates bash git aws-cli

COPY --from=builder /build/greenhouse_api /greenhouse_api

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder ["/build/.docker/entrypoint.sh", "/"]

RUN chmod +x /entrypoint.sh /greenhouse_api

EXPOSE 9000

ENTRYPOINT ["/entrypoint.sh"]