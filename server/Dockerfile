FROM golang

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /go/src/gf-vue-admin/server
COPY . .

RUN go mod tidy
RUN go build -tags "mysql" -o gfva cmd/main.go
RUN go env
RUN go build -tags "mysql" -o server .

FROM alpine:latest
MAINTAINER SliverHorn <sliver_horn@qq.com>

ENV GFVA_CONFIG=DockerCompose

WORKDIR go/src/gf-vue-admin/

# copy server
COPY --from=0 /go/src/gf-vue-admin/server/gfva ./
COPY --from=0 /go/src/gf-vue-admin/server/server ./
COPY --from=0 /go/src/gf-vue-admin/server/public ./public
COPY --from=0 /go/src/gf-vue-admin/server/config ./config
COPY --from=0 /go/src/gf-vue-admin/server/swagger ./swagger
COPY --from=0 /go/src/gf-vue-admin/server/template ./template

EXPOSE 8888

CMD ./gfva initdb -p config/config.mysql.yaml && ./server