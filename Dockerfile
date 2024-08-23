FROM golang:1.23.0

RUN go install github.com/air-verse/air@latest

WORKDIR /usr/src/app
COPY . .
RUN go mod tidy