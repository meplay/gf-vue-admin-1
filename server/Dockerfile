FROM golang:alpine as server

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /go/src/gf-vue-admin
COPY ./ ./

RUN go build -tags "mysql" -o gfva cmd/main.go
RUN go env && go list && go build -tags "mysql" -o server .

FROM alpine:latest
MAINTAINER SliverHorn <sliver_horn@qq.com>

WORKDIR gf-vue-admin/

# copy server
COPY --from=server /go/src/gf-vue-admin ./

EXPOSE 8888

ENTRYPOINT ./gfva initdb -p config/config.docker-compose.yaml && ./server -c config/config.docker-compose.yaml