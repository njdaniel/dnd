MAINTAINER nicholasjdaniel@gmail.com
FROM golang:1.15.3 as builder
ENV GO111MODULE=on

#Copy repo to builder container
RUN mkdir -p /service
WORKDIR /service
COPY . .

#Build
WORKDIR /service/dnd
RUN go build

#============
#Run
FROM alpine:3.12.1
COPY --from=builder /service/dnd/dnd /app/dnd
WORKDIR /app

