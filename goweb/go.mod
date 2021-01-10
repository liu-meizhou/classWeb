module goweb

go 1.15

require (
	github.com/beego/beego/v2 v2.0.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/extrame/ole2 v0.0.0-20160812065207-d69429661ad7 // indirect
	github.com/extrame/xls v0.0.1
	github.com/lib/pq v1.9.0
	github.com/prometheus/common v0.10.0
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee // indirect
)

//go mod init  # 初始化go.mod
//go mod tidy  # 更新依赖文件
//go mod download  # 下载依赖文件
//go mod vendor  # 将依赖转移至本地的vendor文件
//go mod edit  # 手动修改依赖文件
//go mod graph  # 打印依赖图
//go mod verify  # 校验依赖
