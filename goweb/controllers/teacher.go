package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
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
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.STUDENT:
		{
			// 学生
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.TEACHER:
		{
			// 老师
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
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
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(teacher)
			} else {
				teacher := new(models.TeacherInfo)
				err := utils.ParseBody(&this.Controller, teacher)
				if err != nil {
					logs.Error(err)
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
					logs.Error(err)
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
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.STUDENT:
		{
			// 学生
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.TEACHER:
		{
			// 老师
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.TEACHER_HEAD:
		{
			// 系主任
			teacher := new(models.TeacherInfo)
			err := utils.ParseBody(&this.Controller, teacher)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			if teacher.TeacherId == "" {
				this.Data["json"] = utils.ErrorReJson("请输入教工号")
				break
			}
			err = models.CreateTeacher(teacher)
			if err != nil {
				logs.Error(err)
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

// DeleteTeacher Get请求删除一个老师
func (this *TeacherController) DeleteTeacher() {
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
			this.Data["json"] = utils.SuccessReJson("目前你不能使用该功能...")
			break
		}
	case utils.STUDENT:
		{
			// 学生
			this.Data["json"] = utils.SuccessReJson("目前你不能使用该功能...")
			break
		}
	case utils.TEACHER:
		{
			// 老师
			this.Data["json"] = utils.SuccessReJson("目前你不能使用该功能...")
			break
		}
	case utils.TEACHER_HEAD:
		{
			// 系主任
			teacherId := this.GetString("teacherId")
			if teacherId == "" {
				this.Data["json"] = utils.ErrorReJson("请输入教师号")
				break
			}
			err := models.DeleteTeacher(&models.TeacherInfo{TeacherId: teacherId})
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(teacherId)
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// GetTeachers 获取老师列表
func (this *TeacherController) GetTeachers() {
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
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.STUDENT:
		{
			// 学生
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.TEACHER:
		{
			// 老师
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.TEACHER_HEAD:
		{
			// 系主任
			teachers, err := models.GetTeachers()
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(teachers)
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
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.STUDENT:
		{
			// 学生
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.TEACHER:
		{
			// 老师
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	case utils.TEACHER_HEAD:
		{
			// 系主任
			this.Data["json"] = utils.NoFoundReJson("目前你不能使用该功能...")
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}
