FROM node:10.23.1-stretch-slim as frontend

WORKDIR /vue-admin

VOLUME /vue-admin

COPY . .

RUN npm install --registry=https://registry.npm.taobao.org \
    && npm run build:prod

FROM nginx:alpine as nginx

COPY nginx.conf /etc/nginx/conf
COPY --from=frontend ./disk/ /usr/share/nginx/html/

CMD ["nginx", "-g", "daemon off;"]


