package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"goweb/utils"
)

type VisitorController struct {
	web.Controller
}

// 重写方法，匹配Visitor的Get，同理也可以重写post
func (this *VisitorController) Get() {
	this.Data["json"] = utils.NewReturnJson(200, "游客get", "success")
	this.ServeJSON()
}

// 自定义匹配规格
func (this *VisitorController) ShowAPIVersion() {
	this.Data["json"] = utils.NewReturnJson(200, "游客V1版本", "success")
	this.ServeJSON()
}
