package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
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
				classId := this.GetString("classId")
				if classId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入班级号")
					break
				}
				class, err := models.ReadClass(classId)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(class)
			} else {
				class := new(models.ClassInfo)
				err := utils.ParseBody(&this.Controller, class)
				if err != nil {
					logs.Error(err)
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
					logs.Error(err)
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
			class := new(models.ClassInfo)
			err := utils.ParseBody(&this.Controller, class)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			if class.ClassId == "" {
				this.Data["json"] = utils.ErrorReJson("请输入班级号")
				break
			}
			err = models.CreateClass(class)
			if err != nil {
				logs.Error(err)
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
	method := this.Ctx.Request.Method
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
			if method == "GET"{
				classId := this.GetString("classId")
				if classId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入班级号")
					break
				}
				class := &models.ClassInfo{ClassId: classId}
				err := models.GetClassCourse(class)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(class.Courses)
			} else {
				rel := new(utils.CourseClassRel)
				err := utils.ParseBody(&this.Controller, rel)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				if rel.ClassId == "" || rel.CourseId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入课程号和班级号")
					break
				}
				err = models.SetClassCourse(&models.ClassInfo{ClassId: rel.ClassId},
				&models.CourseInfo{CourseId: rel.CourseId})
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson("成功选上")
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
	method := this.Ctx.Request.Method
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
			if method == "GET"{
				classId := this.GetString("classId")
				if classId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入班级号")
					break
				}
				class := &models.ClassInfo{ClassId: classId}
				err := models.GetClassStudentSort(class)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(class.Students)
			} else {

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

// GetClassList 获取班级列表
func (this *ClassController) GetClassList() {
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
			classes, err := models.GetClassList()
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.NoFoundReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(classes)
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// GetClass 获取班级
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
			this.Data["json"] = utils.SuccessReJson("目前你不能使用该功能...")
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}
