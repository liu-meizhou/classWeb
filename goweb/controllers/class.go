package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/prometheus/common/log"
	"goweb/models"
	"goweb/utils"
)

type ClassController struct {
	web.Controller
}

// ClassInfo Get请求获取某个班级的详细信息
// Post请求带上id则更新班级信息
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
	method := this.Ctx.Request.Method
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
			if method == "GET"{
				classId := this.GetString("classId")
				if classId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入班级号")
					break
				}
				class, err := models.ReadClass(classId)
				if err != nil {
					log.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(class)
			} else {
				class := new(models.ClassInfo)
				err := utils.ParseBody(&this.Controller, class)
				if err != nil {
					log.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				if class.ClassId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入班级号")
					break
				}
				// 修改学生
				err = models.UpdateClass(class)
				if err != nil {
					log.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(class)
			}
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// CreateClass Post创建一个新班级
func (this *ClassController) CreateClass() {
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
			class := new(models.ClassInfo)
			err := utils.ParseBody(&this.Controller, class)
			if err != nil {
				log.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			if class.ClassId == "" {
				this.Data["json"] = utils.ErrorReJson("请输入班级号")
				break
			}
			err = models.CreateClass(class)
			if err != nil {
				log.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(class)
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// ClassCourse Get请求获取班级所选的课程
// Post请求为班级统一选课
func (this *ClassController) ClassCourse() {
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

// ClassSort Get获取班级学生绩点排名列表
func (this *ClassController) ClassSort() {
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
