package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type VisitorController struct {
	web.Controller
}

// 重写方法，匹配Visitor的Get，同理也可以重写post
func(this *VisitorController) Get() {
	this.Ctx.WriteString("游客get")
}

// 自定义匹配规格
func(this *VisitorController) ShowAPIVersion() {
	this.Ctx.WriteString("游客V1版本")
}

