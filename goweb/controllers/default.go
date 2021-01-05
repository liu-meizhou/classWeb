package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"goweb/utils"
)

// 相关控制器操作 https://beego.me/docs/mvc/controller/controller.md
// 获取参数: https://beego.me/docs/mvc/controller/params.md
// Session 控制: https://beego.me/docs/mvc/controller/session.md
//

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	this.Data["json"] = utils.NewReturnJson(200, "V1版本", "success")
	this.ServeJSON()
}

//func (this *MainController) Get() {
//	this.Data["Website"] = "beego.me"
//	this.Data["Email"] = "astaxie@gmail.com"
//	this.TplName = "index.tpl"
//}
