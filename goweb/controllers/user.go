package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"goweb/models"
	"goweb/utils"
	"sync"
)

// 	// 获取token
//	token := this.Ctx.Input.Header("token")
//	// 从缓存获取当前登录用户
//	userTypeInfoInterface, ok := userCache.Load(token)
//	if !ok {
//		this.Data["json"] = utils.NoIdentifyReJson( "请登录...")
//		this.ServeJSON()
//		return
//	}
//	userTypeInfo := userTypeInfoInterface.(*UserTypeInfo)
//	switch userTypeInfo.UserType {
//	case 1: {
//		// admin
//		break
//	}
//	case 2: {
//		// 学生
//		break
//	}
//	case 3: {
//		// 老师
//		break
//	}
//	case 4: {
//		// 系主任
//		break
//	}
//	default: {
//		this.Data["json"] = utils.NoFoundReJson( "未知用户...")
//	}
//	}
//	this.ServeJSON()

type UserController struct {
	web.Controller
}

var userCache sync.Map

//ParseForm(obj interface{}) error
//
//将表单反序列化到 obj 对象中。

//GetXXX(key string, def…) XXX, err
//
//从传入参数中，读取某个值。如果传入了默认值，那么在读取不到的情况下，返回默认值，否则返回错误。XXX 可以是 golang 所支持的基本类型，或者是 string, File 对象

// 后门列表
// 1865400006 学生 刘佳合  2
// 111666 老师 杨朔  3
// 100755 系主任 李传中 4

// Identify 识别身份的过滤器
func Identify(ctx *context.Context) bool {
	token := ctx.Input.Header("token")

	// 查看缓存是否有用户
	if _ ,ok:=userCache.Load(token);ok{
		return true
	}

	// 走后门 学生刘佳合
	if token == "1865400006" {
		userTypeInfo, err := models.Login(utils.NewTokenInfo("1865400006","123456", 2))
		if err != nil {
			logs.Error(err)
			ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
			return false
		}
		userCache.Store(token, userTypeInfo)
		return true
	}
	// 老师 杨朔
	if token == "111666" {
		userTypeInfo, err := models.Login(utils.NewTokenInfo("111666", "123456", 3))
		if err != nil {
			logs.Error(err)
			ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
			return false
		}
		userCache.Store(token, userTypeInfo)
		return true
	}
	// 系主任 李传中
	if token == "100755" {
		userTypeInfo, err := models.Login(utils.NewTokenInfo("100755", "123456", 4))
		if err != nil {
			logs.Error(err)
			ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
			return false
		}
		userCache.Store(token, userTypeInfo)
		return true
	}

	if token == "" {
		logs.Debug("不存在token")
		ctx.Output.JSON(utils.NoIdentifyReJson("不存在token"), web.BConfig.RunMode != web.PROD, true)
		return false
	}
	tokenInfo, err := utils.ParseToken(token)
	if err != nil {
		logs.Debug(err)
		ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
		return false
	}

	// 访问数据库
	userTypeInfo, err := models.Login(tokenInfo)
	if err != nil {
		logs.Error(err)
		ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
		return false
	}
	userCache.Store(token, userTypeInfo)
	// userCache有内存泄漏的风险
	return true
}

func (this *UserController) Prepare() {
	logs.Info("开始请求普通用户")
}

func (this *UserController) Finish() {
	logs.Info("结束请求普通用户")
}

// GetCourse 获取课表
func (this *UserController) GetCourse() {
	// 获取token
	token := this.Ctx.Input.Header("token")
	// 从缓存获取当前登录用户
	userTypeInfoInterface, ok := userCache.Load(token)
	if !ok {
		this.Data["json"] = utils.NoIdentifyReJson( "请登录...")
		this.ServeJSON()
		return
	}
	user := userTypeInfoInterface.(*utils.User)
	switch user.UserType {
	case utils.ADMIN: {
		// admin
		break
	}
	case utils.STUDENT: {
		// 学生
		studentInfo := user.User.(*models.StudentInfo)
		if studentInfo.Courses == nil {
			_ = models.GetStudentCourse(studentInfo)
		}
		this.Data["json"] = utils.SuccessReJson(studentInfo.Courses)
		break
	}
	case utils.TEACHER: {
		// 老师
		teacherInfo := user.User.(*models.TeacherInfo)
		if teacherInfo.Courses == nil {
			_ = models.GetTeacherCourse(teacherInfo)
		}
		this.Data["json"] = utils.SuccessReJson(teacherInfo.Courses)
		break
	}
	case utils.TEACHER_HEAD: {
		// 系主任
		teacherInfo := user.User.(*models.TeacherInfo)
		// 获取要看的人的类型和id，如没有就是查看自己的课表
		userType, err := this.GetInt("userType", user.UserType)
		if err != nil {
			this.Data["json"] = utils.ErrorReJson(err)
			break
		}
		userId := this.GetString("userId", teacherInfo.TeacherId)
		if userType == 2 {
			// 获取学生课表
			student := &models.StudentInfo{StudentId: userId}
			err = models.GetStudentCourse(student)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(student)
			break
		}
		// 获取老师课表
		teacher := &models.TeacherInfo{TeacherId: userId}
		err = models.GetTeacherCourse(teacher)
		if err != nil {
			this.Data["json"] = utils.ErrorReJson(err)
			break
		}
		this.Data["json"] = utils.SuccessReJson(teacher)
		break
	}
	default: {
		this.Data["json"] = utils.NoFoundReJson( "未知用户...")
	}
	}
	this.ServeJSON()
}

// ExportCourse 导出课表
func (this *UserController) ExportCourse() {
	// 获取token
	token := this.Ctx.Input.Header("token")
	// 从缓存获取当前登录用户
	userTypeInfoInterface, ok := userCache.Load(token)
	if !ok {
		this.Data["json"] = utils.NoIdentifyReJson( "请登录...")
		this.ServeJSON()
		return
	}
	user := userTypeInfoInterface.(*utils.User)
	switch user.UserType {
	case utils.ADMIN: {
		// admin
		break
	}
	case utils.STUDENT: {
		// 学生
		break
	}
	case utils.TEACHER: {
		// 老师
		break
	}
	case utils.TEACHER_HEAD: {
		// 系主任
		break
	}
	default: {
		this.Data["json"] = utils.NoFoundReJson( "未知用户...")
	}
	}
	this.ServeJSON()
}

// ChooseCourse 选课  Get请求获取选课列表 Post请求进行选课
func (this *UserController) ChooseCourse() {
	// 获取token
	token := this.Ctx.Input.Header("token")
	// 从缓存获取当前登录用户
	userTypeInfoInterface, ok := userCache.Load(token)
	if !ok {
		this.Data["json"] = utils.NoIdentifyReJson( "请登录...")
		this.ServeJSON()
		return
	}
	// 获取请求类型
	method := this.Ctx.Request.Method
	user := userTypeInfoInterface.(*utils.User)
	switch user.UserType {
	case utils.ADMIN: {
		// admin
		break
	}
	case utils.STUDENT: {
		// 学生
		studentInfo := user.User.(*models.StudentInfo)
		if method == "GET" {
			courses, err := models.GetChooseCourse(studentInfo)
			if err != nil {
				this.Data["json"] = utils.ErrorReJson(err)
				break
			}
			this.Data["json"] = utils.SuccessReJson(courses)
		} else if method == "POST" {

		}
		break
	}
	case utils.TEACHER: {
		// 老师
		break
	}
	case utils.TEACHER_HEAD: {
		// 系主任
		break
	}
	default: {
		this.Data["json"] = utils.NoFoundReJson( "未知用户...")
	}
	}
	this.ServeJSON()
}

// CourseGrade 获取课程成绩等信息  Get请求获取学生课程成绩, Post请求进行设置成绩
func (this *UserController) CourseGrade() {
	// 获取token
	token := this.Ctx.Input.Header("token")
	// 从缓存获取当前登录用户
	userTypeInfoInterface, ok := userCache.Load(token)
	if !ok {
		this.Data["json"] = utils.NoIdentifyReJson( "请登录...")
		this.ServeJSON()
		return
	}
	method := this.Ctx.Request.Method
	user := userTypeInfoInterface.(*utils.User)
	switch user.UserType {
	case utils.ADMIN: {
		// admin
		break
	}
	case utils.STUDENT: {
		// 学生
		break
	}
	case utils.TEACHER: {
		// 老师
		//teacherInfo := userTypeInfo.User.(*models.TeacherInfo)
		if method == "GET" {
			// 获取要查询的课程号
			courseId := this.GetString("courseId")
			if courseId == "" {
				this.Data["json"] = utils.ErrorReJson("输入课程号")
				break
			}
			students, err := models.GetGradeCourse(&models.CourseInfo{CourseId: courseId})
			if err != nil {
				this.Data["json"] = utils.ErrorReJson(err)
				break
			}
			this.Data["json"] = utils.SuccessReJson(students)
		} else if method == "POST" {

		}
		break
	}
	case utils.TEACHER_HEAD: {
		// 系主任
		break
	}
	default: {
		this.Data["json"] = utils.NoFoundReJson( "未知用户...")
	}
	}
	this.ServeJSON()
}
