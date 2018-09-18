FROM golang:1.11-alpine
RUN apk add git

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./

RUN time go mod download

# this fails
RUN go build ./...
# => go: warning: "./..." matched no packages

COPY . .

# this now works but isn't needed
# RUN time go build ./...

# compile app along with any unbuilt deps
RUN time go build
