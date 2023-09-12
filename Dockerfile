FROM golang:1.13-buster

WORKDIR /app

COPY . . 

RUN go mod download
RUN go build

EXPOSE 3000

ENTRYPOINT [ "go-ping-pong" ]
