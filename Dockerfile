# syntax=docker/dockerfile:1

FROM golang:1.18.4-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY .env ./

RUN go build -o /btc-rate-api

EXPOSE 8080

CMD [ "/btc-rate-api" ]