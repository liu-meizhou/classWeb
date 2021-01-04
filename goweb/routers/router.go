package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"goweb/controllers"
)

// 相关路由配置: https://beego.me/docs/mvc/controller/router.md

func init() {
	// 支持 基本路由, 固定路由和正则路由
	web.Get("/index", func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})
	web.Router("/", &controllers.MainController{}, "get:Get;post:Get;*:Get")
	web.Router("/api/?:id", &controllers.MainController{})

	// 命名空间配置路由
	//初始化 游客命名空间
	visitor :=
		web.NewNamespace("/visitor/v1",
			// return true匹配该命名空间, 可以作为拦截器
			web.NSCond(func(ctx *context.Context) bool {
				return true
			}),
			// 执行controller前先执行该函数, 常常用于过滤器, 可以注册多个
			//web.NSBefore(func),
			// 执行controller后执行该函数
			//web.NSAfter(func),
			// 下面都是设置路由
			web.NSGet("/nsGet", func(ctx *context.Context) {
				ctx.Output.Body([]byte("游客nsGet"))
			}),
			web.NSRouter("/get", &controllers.VisitorController{}),
			web.NSRouter("/version", &controllers.VisitorController{}, "get:ShowAPIVersion"),
			// 命名空间内嵌套命名空间
			web.NSNamespace("/show",
				//web.NSBefore(func),
				web.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("展示游客"))
				}),
			),
		)
	//注册 namespace
	web.AddNamespace(visitor)
	//初始化 登录用户命名空间
	user :=
		web.NewNamespace("/user/v1",
			web.NSCond(func(ctx *context.Context) bool {
				return true
			}),
			web.NSGet("/get", func(ctx *context.Context) {
				ctx.Output.Body([]byte("登录用户get"))
			}),
			web.NSRouter("/version", &controllers.UserController{}, "get:ShowAPIVersion"),
			web.NSNamespace("/show",
				web.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("展示登录用户"))
				}),
			),
		)
	//注册 namespace
	web.AddNamespace(user)
	//初始化 admin命名空间
	admin :=
		web.NewNamespace("/admin/v1",
			web.NSCond(func(ctx *context.Context) bool {
				return true
			}),
			web.NSGet("/get", func(ctx *context.Context) {
				ctx.Output.Body([]byte("管理员get"))
			}),
			web.NSRouter("/version", &controllers.AdminController{}, "get:ShowAPIVersion"),
			web.NSNamespace("/show",
				web.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("展示管理员"))
				}),
			),
		)
	//注册 namespace
	web.AddNamespace(admin)
}
