##
## build
##
FROM golang:1.20-alpine as builder
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN  go build -o server .


##
## deploy
##
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server ./
COPY --from=builder /app/internal/conf/conf.online.yaml ./internal/conf/conf.online.yaml

EXPOSE 8888
ENTRYPOINT ./server
