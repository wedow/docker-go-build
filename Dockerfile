FROM golang:1.11-alpine
RUN apk add git

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

COPY go.mod go.sum dummy/main.go ./

# compile imports from dummy main
RUN time go build

# remove the dummy and add real project code
RUN rm main.go
COPY . .

# compile app along with any unbuilt deps
RUN time go build
