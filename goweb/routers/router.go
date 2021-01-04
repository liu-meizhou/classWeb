package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"goweb/controllers"
)

func init() {
	// 支持 基本路由, 固定路由和正则路由
	web.Get("/index", func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})
	web.Router("/", &controllers.MainController{}, "get:Get;post:Get;*:Get")
	web.Router("/api/?:id", &controllers.MainController{})
}
