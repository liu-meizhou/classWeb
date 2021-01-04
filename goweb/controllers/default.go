package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

// 相关控制器操作 https://beego.me/docs/mvc/controller/controller.md
// 获取参数: https://beego.me/docs/mvc/controller/params.md
// Session 控制: https://beego.me/docs/mvc/controller/session.md
//

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello")
}

//func (this *MainController) Get() {
//	this.Data["Website"] = "beego.me"
//	this.Data["Email"] = "astaxie@gmail.com"
//	this.TplName = "index.tpl"
//}
