# syntax=docker/dockerfile:1

FROM golang:alpine AS build

WORKDIR /go/src/stargazers-cmc/

COPY ./src/* .
RUN go build -o app main.go


FROM golang:alpine
WORKDIR /go/bin/
COPY --from=build /go/src/stargazers-cmc/app ./app-stargazers
ENTRYPOINT ["./app-stargazers"]

