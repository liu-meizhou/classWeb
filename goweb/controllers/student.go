package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"goweb/models"
	"goweb/utils"
)

type StudentController struct {
	web.Controller
}

// StudentInfo Get获取一个学生详细信息
// Post带学号修改学生个人信息
func (this *StudentController) StudentInfo() {
	// 获取token
	token := this.Ctx.Input.Header("token")
	method := this.Ctx.Request.Method
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
			if method == "GET"{
				studentId := this.GetString("studentId")
				if studentId == ""{
					this.Data["json"] = utils.ErrorReJson("请输入学号")
					break
				}
				student, err := models.ReadStudent(studentId)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(student)
			} else {
				student := new(models.StudentInfo)
				err := utils.ParseBody(&this.Controller, student)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				if student.StudentId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入学号")
					break
				}
				// 修改学生
				err = models.UpdateStudent(student)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(student)
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

// CreateStudent Post创建学生
func (this *StudentController) CreateStudent() {
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
			student := new(models.StudentInfo)
			err := utils.ParseBody(&this.Controller, student)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			if student.StudentId == "" {
				this.Data["json"] = utils.ErrorReJson("请输入学号")
				break
			}
			err = models.CreateStudent(student)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(student)
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// DeleteStudent Get请求删除一个学生
func (this *StudentController) DeleteStudent() {
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
			studentId := this.GetString("studentId")
			if studentId == "" {
				this.Data["json"] = utils.ErrorReJson("请输入学生学号")
				break
			}
			err := models.DeleteStudent(&models.StudentInfo{StudentId: studentId})
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(studentId)
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// EditClass Get编辑学生班级信息
func (this *StudentController) EditClass() {
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
			studentId := this.GetString("studentId")
			classId := this.GetString("classId")
			if studentId == "" {
				this.Data["json"] = utils.ErrorReJson("请输入要修改班级的学生学号")
				break
			}
			err := models.UpdateStudentClass(studentId, classId)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(studentId)
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// GetStudentList Get获取学生列表
func (this *StudentController) GetStudentList() {
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
			// 解析查询条件
			pageInfo := new(utils.PageInfo)
			student := new(models.StudentInfo)
			err := utils.ParsePageInfo(&this.Controller, pageInfo, student)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			err = models.GetStudentList(pageInfo, student)
			if err != nil {
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(pageInfo)
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// GetGradeSortStudentList Get获取学生排序列表
func (this *StudentController) GetGradeSortStudentList() {
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
