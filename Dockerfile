FROM golang:1.13-alpine AS build

WORKDIR /go
RUN apk update \
  && apk add --no-cache git \
  && go get github.com/labstack/echo/... \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/jinzhu/gorm
