package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"goweb/utils"
)

type UserController struct {
	web.Controller
}

//ParseForm(obj interface{}) error
//
//将表单反序列化到 obj 对象中。

//GetXXX(key string, def…) XXX, err
//
//从传入参数中，读取某个值。如果传入了默认值，那么在读取不到的情况下，返回默认值，否则返回错误。XXX 可以是 golang 所支持的基本类型，或者是 string, File 对象

//SetSession(name interface{}, value interface{}) error
//
//往Session中设置值。
//
//GetSession(name interface{}) interface{}
//
//从Session中读取值。
//
//DelSession(name interface{}) error
//
//从Session中删除某项。
//
//SessionRegenerateID() error
//
//重新生成一个SessionId。
//
//DestroySession() error
//
//销毁Session

func (this *UserController) Prepare() {
	logs.Info("开始请求普通用户")
}

func (this *UserController) Finish() {
	logs.Info("结束请求普通用户")
}

// GetCourse 获取课表
func (this *UserController) GetCourse() {
	this.Data["json"] = utils.NewReturnJson(200, "Get请求获取课程列表", "success")
	this.ServeJSON()
}

// ExportCourse 导出课表
func (this *UserController) ExportCourse() {
	this.Data["json"] = utils.NewReturnJson(200, "Get请求导出课表", "success")
	this.ServeJSON()
}

// ChooseCourse 选课
func (this *UserController) ChooseCourse() {
	this.Data["json"] = utils.NewReturnJson(200, "Get请求获取选课列表，Post请求进行选课", "success")
	this.ServeJSON()
}

// CourseGrade 获取课程成绩等信息
func (this *UserController) CourseGrade() {
	this.Data["json"] = utils.NewReturnJson(200, "Get请求获取学生课程成绩, Post请求进行设置成绩", "success")
	this.ServeJSON()
}
