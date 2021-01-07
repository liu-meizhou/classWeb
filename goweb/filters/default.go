package filters

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// 更多过滤器操作 https://beego.me/docs/mvc/controller/filter.md

// 过滤重复调用请求 防止一个ip在一定时间间隔内重复调用相同api
var FilterReCall = func(ctx *context.Context) {
	logs.Info("记得做: 防止重复调用后端API")
}

func init() {
	// 配置过滤器规则基本和路由一致
	web.InsertFilter("/*", web.BeforeRouter, FilterReCall)
}
