FROM golang:1.15.6-alpine3.12
WORKDIR /goweb

COPY go.mod,go.sum .
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY="https://goproxy.io,direct" \
    && go mod vendor

# 执行的时候需要挂上数据卷(源路径是项目路径,目标路径是/goweb)，如不使用数据卷，请添加下面的命令
# COPY . .

CMD go build -o main main.go && ./main
