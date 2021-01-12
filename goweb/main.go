package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	_ "goweb/filters"
	_ "goweb/routers"
	"os"
)
func init() {
	// need to register db driver
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error(err)
		os.Exit(-1)
	}

	// need to register default database
	err = orm.RegisterDataBase("default", "postgres",
		"postgresql://postgres:123456@42.193.143.9:5432/postgres?sslmode=disable&&TimeZone=Asia/Shanghai")
	if err != nil {
		logs.Error(err)
		os.Exit(-1)
	}
}

func main() {
	// 配置运行跨域，配置不成功0.0, 改为配置前端
	// 前端配置：https://www.jianshu.com/p/43aa317d7683
	//web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowOrigins: []string{"https://*.foo.com"},
	//	AllowMethods: []string{"GET", "POST", "PUT", "PATCH"},
	//	AllowHeaders: []string{"Origin"},
	//	ExposeHeaders: []string{"Content-Length"},
	//	AllowCredentials: true,
	//}))

	//web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
	//	//AllowAllOrigins:  true,
	//	AllowMethods:     []string{"*"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type"},
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
	//	AllowCredentials: true,
	//	AllowOrigins: []string{"http://10.*.*.*:*","http://localhost:*","http://127.0.0.1:*"},
	//}))

	// 生成数据库
	orm.Debug = true
	logs.Info("测试自动化部署")

	// 启动web
	web.Run()
}
