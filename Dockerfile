FROM golang:latest as builder

WORKDIR /go/src/github.com/xxarupakaxx/go-todo-api


RUN go get -u github.com/labstack/echo

COPY . .

ENV CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64

EXPOSE 8080

CMD go run main.go