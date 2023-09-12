FROM golang:1.18-buster

WORKDIR /app

COPY go.mod go.sum main.go main_test.go /app/

RUN go mod download
RUN go build

EXPOSE 3000

CMD ./go-ping-pong
