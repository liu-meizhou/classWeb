package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "goweb/routers"
)


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

	web.Run()
}