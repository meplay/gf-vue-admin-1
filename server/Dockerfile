FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /go/src/gf-vue-admin
COPY server/ ./

RUN go env && go list && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/gf-vue-admin

COPY --from=0 /go/src/gf-vue-admin/server ./
COPY --from=0 /go/src/gf-vue-admin/i18n ./i18n
COPY --from=0 /go/src/gf-vue-admin/public ./public
COPY --from=0 /go/src/gf-vue-admin/config ./config
COPY --from=0 /go/src/gf-vue-admin/template ./template

ENTRYPOINT ./server
# 根据Dockerfile生成Docker镜像

# docker build -t gva-server:1.0 .

#- 根据Docker镜像启动Docker容器
#    - 后台运行
#    - ```
#    docker run -d -p 8888:8888 --name gva-server-v1 gva-server:1.0
#      ```
#    - 以可交互模式运行, Ctrl + p + q 后台运行
#    - ```
#    docker run -it -p 8888:8888 --name gva-server-v1 gva-server:1.0
#      ```