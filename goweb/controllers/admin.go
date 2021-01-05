package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type AdminController struct {
	web.Controller
}

func (this *AdminController) Get() {
	this.Ctx.WriteString("管理员get")
}

func (this *AdminController) ShowAPIVersion() {
	this.Ctx.WriteString("管理员V1版本")
}
