package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"goweb/models"
	"goweb/utils"
)

type ReturnMessage struct {
	User  interface{} `json:"user"`
	Token string `json:"token"`
}

type VisitorController struct {
	web.Controller
}

func (this *VisitorController) Prepare() {
	logs.Info("开始请求普通游客")
}

func (this *VisitorController) Finish() {
	logs.Info("结束请求普通游客")
}

// 登录
func (this *VisitorController) Login() {
	// 防止重复登录
	loginInfo := new(utils.TokenInfo)
	err := utils.ParseBody(&this.Controller, loginInfo)
	if err!=nil {
		logs.Error(err)
		this.ServeJSON()
		return
	}
	user, err := models.Login(loginInfo)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = utils.ErrorReJson(err.Error())
		this.ServeJSON()
		return
	}
	// 登录成功 创建token并且加入缓存
	token := utils.CreateToken(loginInfo)
	SetUser(token, user)

	this.Data["json"] = utils.SuccessReJson(&ReturnMessage{User: user.User, Token: token})
	this.ServeJSON()
}

// 自定义匹配规格
func (this *VisitorController) Register() {
	this.Data["json"] = utils.SuccessReJson("想注册? 没门")
	this.ServeJSON()
}
