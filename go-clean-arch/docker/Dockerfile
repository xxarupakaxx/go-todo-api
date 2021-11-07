FROM golang:1.17.3-alpine as server-build
RUN apk add --update --no-cache ca-certificates git

WORKDIR /go/src/github.com/go-todo-api/go-clean-arch

RUN go install github.com/cosmtrek/air@v1.27.3

COPY go.mod go.sum ./
RUN go mod download

CMD["air","-c","go-clean-arch/docker/.air.toml"]
