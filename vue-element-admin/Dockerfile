FROM node:10.23.1-stretch-slim as frontend

WORKDIR /vue-admin

COPY package.json .

RUN npm install --registry=https://registry.npm.taobao.org

COPY . .

RUN npm run build:prod

FROM nginx:alpine as nginx

COPY --from=frontend nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=frontend dist/* /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]

