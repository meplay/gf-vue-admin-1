FROM node:14.16.0

WORKDIR /go/src/gf-vue-admin/web
COPY . .

RUN npm install -g cnpm --registry=https://registry.npm.taobao.org
RUN cnpm install || npm install
RUN npm run build

FROM nginx:alpine
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=0 /go/src/gf-vue-admin/web/dist /usr/share/nginx/html
RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/my.conf
RUN ls -al /usr/share/nginx/html