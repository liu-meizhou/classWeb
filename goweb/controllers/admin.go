package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"goweb/DB"
	"goweb/utils"
)

type AdminController struct {
	web.Controller
}

func (this *AdminController) Get() {
	isTable := this.GetString("isTable", "true")
	if isTable == "true" {
		DB.GenTable()
	}
	err := DB.GenData()
	if err != nil {
		logs.Debug(err)
		this.Data["json"] = utils.SuccessReJson(err.Error())
		this.ServeJSON()
		return
	}
	this.Data["json"] = utils.SuccessReJson("成功生成！！！")
	this.ServeJSON()
}

func (this *AdminController) ShowAPIVersion() {
	this.Ctx.WriteString("管理员V1版本")
}
