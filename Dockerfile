FROM node:12.16.1 as web

WORKDIR /web/
COPY web/ .

RUN cat .env.production
COPY docker/web-handle.sh .
RUN sh ./web-handle.sh
RUN cat .env.production
RUN rm -f web-handle.sh

RUN npm install -g cnpm --registry=https://registry.npm.taobao.org
RUN cnpm install
RUN npm run build

FROM golang:alpine as server

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /go/src/gf-vue-admin
COPY server/ ./

RUN cat ./boot/server.go
RUN cat ./config/config.toml
COPY docker/server-handle.sh .
RUN sh ./server-handle.sh
RUN rm -f server-handle.sh
RUN cat ./boot/server.go
RUN cat ./config/config.toml

RUN go env && go list && go build -o server .


FROM nginx:alpine
MAINTAINER SliverHorn <sliver_horn@qq.com>

WORKDIR gf-vue-admin/

# copy web
COPY --from=web /web/dist ./public/dist
# copy server
COPY --from=server /go/src/gf-vue-admin/server ./
COPY --from=server /go/src/gf-vue-admin/i18n ./i18n
COPY --from=server /go/src/gf-vue-admin/public ./public
COPY --from=server /go/src/gf-vue-admin/config ./config
COPY --from=server /go/src/gf-vue-admin/template ./template

EXPOSE 8888

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