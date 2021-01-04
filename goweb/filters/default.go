package filters

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// 更多过滤器操作 https://beego.me/docs/mvc/controller/filter.md

var FilterUser = func(ctx *context.Context) {
	//_, ok := ctx.Input.Session("uid").(int)
	//if !ok && ctx.Request.RequestURI != "/login" {
	//	ctx.Redirect(302, "/login")
	//}

}

func init()  {
	// 配置过滤器规则基本和路由一致
	web.InsertFilter("/*", web.BeforeRouter, FilterUser)
}