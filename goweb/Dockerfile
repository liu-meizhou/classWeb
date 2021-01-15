FROM golang:1.15.6-alpine3.12
WORKDIR /go_web

COPY go.mod .
COPY go.sum .
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY="https://goproxy.io,direct" \
    && go mod vendor

# 执行的时候需要挂上数据卷(源路径是项目路径,目标路径是/goweb)，如不使用数据卷，请添加下面的命令
# COPY . .

CMD go build -o main main.go && ./main

# 如docker启动
# docker build -t go_web .
# docker run -d -p 8080:8080 -v ~/project/classWeb/goweb:/go_web -e TZ=Asia/Shanghai --name go_web go_web

# 用docker-compose启动
# docker-compose up -d
