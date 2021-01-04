package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (this *UserController) Get() {
	this.Ctx.WriteString("用户get")
}

func(this *UserController) ShowAPIVersion() {
	this.Ctx.WriteString("用户V1版本")
}
