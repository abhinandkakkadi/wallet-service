# build stage
FROM golang:1.19-alpine3.16 AS builder
#maintainer info
LABEL maintainer="Sethukumarj <Sethukumarj.76@gmail.com>"
#installing git
RUN apk update && apk add --no-cache git

WORKDIR /auth_svc

COPY . .

RUN apk add --no-cache make

RUN make deps
RUN go mod vendor
RUN make build



# Run stage
FROM alpine:3.16

WORKDIR /auth_svc
COPY go.mod .
COPY go.sum .
COPY --from=builder /auth_svc/build/bin/api .


CMD [ "/auth_svc/api"] 