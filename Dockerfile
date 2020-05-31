FROM golang:1.14-alpine AS builder
WORKDIR /usr/src/app

ARG SSH_PRIVATE_KEY
ENV GO111MODULE=on

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /usr/src/app/bin/main /usr/src/app/main.go

# Main Image
FROM alpine:latest

RUN apk update \
  && apk --no-cache add \
  ca-certificates openssl && update-ca-certificates

# Setting timezone
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Setting folder workdir
WORKDIR /usr/src/app
RUN mkdir -p etc/kube-siera

COPY --from=builder /usr/src/app/bin/main .

CMD ["./main"]
