FROM node:10.23.1-stretch-slim as frontend

WORKDIR /vue-admin

COPY package.json .

RUN npm install --registry=https://registry.npm.taobao.org

COPY . .

RUN npm run build:prod

FROM nginx:alpine as nginx

COPY --from=frontend /vue-admin/nginx-docker.conf /etc/nginx/conf.d/default.conf
COPY --from=frontend /vue-admin/dist/ /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]

# 如docker启动
# docker build -t vue_admin_nginx .
# docker run -d -p 9999:80 -e TZ=Asia/Shanghai --name vue_admin_nginx vue_admin_nginx

# 用docker-compose启动
# docker-compose up -d
