module goweb

go 1.15

require (
	github.com/beego/beego/v2 v2.0.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/lib/pq v1.9.0
)

// go get -insecure github.com/beego/beego/v2@v2.0.0

//go mod init  # 初始化go.mod
//go mod tidy  # 更新依赖文件
//go mod download  # 下载依赖文件
//go mod vendor  # 将依赖转移至本地的vendor文件
//go mod edit  # 手动修改依赖文件
//go mod graph  # 打印依赖图
//go mod verify  # 校验依赖
