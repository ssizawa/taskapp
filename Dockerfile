FROM golang:1.18.3-alpine AS dev

RUN apk update

WORKDIR /src

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

CMD ["air"]