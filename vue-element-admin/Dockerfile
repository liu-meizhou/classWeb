FROM node:alpine as frontend

VOLUME /vue-admin

COPY package.json /vue-admin/

RUN set -x ; cd /vue-admin \
    && npm install --registry=https://registry.npm.taobao.org \
    && npm run build:prod

FROM nginx:alpine as nginx

COPY nginx.conf /etc/nginx/conf
COPY --from=0 /vue-admin/disk/ /usr/share/nginx/html/
