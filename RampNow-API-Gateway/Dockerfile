# build stage
FROM golang:1.19-alpine3.16 AS builder
#maintainer info
LABEL maintainer="Sethukumarj <Sethukumarj.76@gmail.com>"
#installing git
RUN apk update && apk add --no-cache git

WORKDIR /api_gateway

COPY . .

RUN apk add --no-cache make

RUN make deps
RUN go mod vendor



EXPOSE 3005


CMD [ "make" , "server"] 