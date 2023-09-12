FROM golang:1.18-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY main.go main_test.go /app/
RUN go build -v -o go-ping-pong

FROM debian:buster-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/go-ping-pong /app/go-ping-pong

ENTRYPOINT [ "/app/go-ping-pong" ]
