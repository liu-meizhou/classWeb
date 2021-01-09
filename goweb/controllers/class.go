package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"goweb/utils"
)

type ClassController struct {
	web.Controller
}

// ClassInfo Get请求获取某个班级的详细信息
// Post请求带上id则更新班级信息，不带id或者id为-1则添加班级信息
func (this *ClassController) ClassInfo() {
	// 获取token
	token := this.Ctx.Input.Header("token")
	// 从缓存获取当前登录用户
	user := GetUser(token)
	if user == nil {
		this.Data["json"] = utils.NoIdentifyReJson("请登录...")
		this.ServeJSON()
		return
	}
	switch user.UserType {
	case utils.ADMIN:
		{
			// admin
			break
		}
	case utils.STUDENT:
		{
			// 学生
			break
		}
	case utils.TEACHER:
		{
			// 老师
			break
		}
	case utils.TEACHER_HEAD:
		{
			// 系主任
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// GetClass 获取班级列表
func (this *ClassController) GetClass() {
	// 获取token
	token := this.Ctx.Input.Header("token")
	// 从缓存获取当前登录用户
	user := GetUser(token)
	if user == nil {
		this.Data["json"] = utils.NoIdentifyReJson("请登录...")
		this.ServeJSON()
		return
	}
	switch user.UserType {
	case utils.ADMIN:
		{
			// admin
			break
		}
	case utils.STUDENT:
		{
			// 学生
			break
		}
	case utils.TEACHER:
		{
			// 老师
			break
		}
	case utils.TEACHER_HEAD:
		{
			// 系主任
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}
