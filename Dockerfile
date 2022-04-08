# syntax=docker/dockerfile:1

# stage 1: build
FROM golang:alpine AS build

WORKDIR /go/src/stargazers-cmc/
COPY ./src/* .
RUN go build -o app main.go

# stage 2: run
FROM golang:alpine
WORKDIR /go/bin/
COPY --from=build /go/src/stargazers-cmc/app ./stargazers-cmc

# run the app, output to stdout
ENTRYPOINT ["./stargazers-cmc"]
