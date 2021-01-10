package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/prometheus/common/log"
	"goweb/models"
	"goweb/utils"
)

type TeacherController struct {
	web.Controller
}


// TeacherInfo Get请求获取某个老师的详细信息
// Post请求带上id则更新班级信息
func (this *TeacherController) TeacherInfo() {
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
				teacherId := this.GetString("teacherId")
				if teacherId == ""{
					this.Data["json"] = utils.ErrorReJson("请输入教工号")
					break
				}
				teacher, err := models.ReadTeacher(teacherId)
				if err != nil {
					log.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(teacher)
			} else {
				teacher := new(models.TeacherInfo)
				err := utils.ParseBody(&this.Controller, teacher)
				if err != nil {
					log.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				if teacher.TeacherId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入教工号")
					break
				}
				// 修改老师
				err = models.UpdateTeacher(teacher)
				if err != nil {
					log.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(teacher)
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

// CreateTeacher Posy创建老师信息
func (this *TeacherController) CreateTeacher() {
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
			teacher := new(models.TeacherInfo)
			err := utils.ParseBody(&this.Controller, teacher)
			if err != nil {
				log.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			if teacher.TeacherId == "" {
				this.Data["json"] = utils.ErrorReJson("请输入教工号")
				break
			}
			err = models.CreateTeacher(teacher)
			if err != nil {
				log.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(teacher)
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// GetTeacher 获取老师详细信息
func (this *TeacherController) GetTeacher() {
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
