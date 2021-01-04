package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

//func (this *MainController) Get() {
//	this.Data["Website"] = "beego.me"
//	this.Data["Email"] = "astaxie@gmail.com"
//	this.TplName = "index.tpl"
//}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello")
}
