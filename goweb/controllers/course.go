package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/prometheus/common/log"
	"goweb/models"
	"goweb/utils"
)

type CourseController struct {
	web.Controller
}

// CourseInfo Get请求获取某个课的详细信息
// Post请求带上id则更新课程信息
func (this *CourseController) CourseInfo() {
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
				courseId := this.GetString("courseId")
				if courseId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入课程号")
					break
				}
				course, err := models.ReadCourse(courseId)
				if err != nil {
					log.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(course)
			} else {
				course := new(models.CourseInfo)
				err := utils.ParseBody(&this.Controller, course)
				if err != nil {
					log.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				if course.CourseId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入课程号")
					break
				}
				// 修改学生
				err = models.UpdateCourse(course)
				if err != nil {
					log.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(course)
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

// CreateCourse Post创建课程
func (this *CourseController) CreateCourse() {
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
			course := new(models.CourseInfo)
			err := utils.ParseBody(&this.Controller, course)
			if err != nil {
				log.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			if course.CourseId == "" {
				this.Data["json"] = utils.ErrorReJson("请输入课程号")
				break
			}
			err = models.CreateCourse(course)
			if err != nil {
				log.Error(err)
				this.Data["json"] = utils.ErrorReJson(err.Error())
				break
			}
			this.Data["json"] = utils.SuccessReJson(course)
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// GetCourse 获取课表
func (this *CourseController) GetCourse() {
	// 获取token
	token := this.Ctx.Input.Header("token")
	// 从缓存获取当前登录用户
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
			studentInfo := user.User.(*models.StudentInfo)
			if studentInfo.Courses == nil {
				_ = models.GetStudentCourse(studentInfo)
			}
			this.Data["json"] = utils.SuccessReJson(studentInfo.Courses)
			break
		}
	case utils.TEACHER:
		{
			// 老师
			teacherInfo := user.User.(*models.TeacherInfo)
			if teacherInfo.Courses == nil {
				_ = models.GetTeacherCourse(teacherInfo)
			}
			this.Data["json"] = utils.SuccessReJson(teacherInfo.Courses)
			break
		}
	case utils.TEACHER_HEAD:
		{
			// 系主任
			teacherInfo := user.User.(*models.TeacherInfo)
			// 获取要看的人的类型和id，如没有就是查看自己的课表
			userType := this.GetString("userType", user.UserType)
			userId := this.GetString("userId", teacherInfo.TeacherId)
			if userType == utils.STUDENT {
				// 获取学生课表
				student := &models.StudentInfo{StudentId: userId}
				err := models.GetStudentCourse(student)
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
			err := models.GetTeacherCourse(teacher)
			if err != nil {
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

// ExportCourse 导出课表
func (this *CourseController) ExportCourse() {
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

// ChooseCourse 选课  学生: Get请求带courseId参数请求进行选课 Post带PageInfo参数获取选课列表
func (this *CourseController) ChooseCourse() {
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
			this.Data["json"] = utils.NoFoundReJson("你不是管理员...")
			break
		}
	case utils.STUDENT:
		{
			// 学生
			studentInfo := user.User.(*models.StudentInfo)
			if method == "POST" {
				// 解析查询条件
				pageInfo := new(utils.PageInfo)
				course := new(models.CourseInfo)
				err := utils.ParsePageInfo(&this.Controller, pageInfo, course)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				err = models.GetChooseCourse(studentInfo, pageInfo, course)
				if err != nil {
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(pageInfo)
			} else {
				// 获取请求参数
				courseId := this.GetString("courseId")
				rel := &models.CourseStudentRel{Student: studentInfo, Course: &models.CourseInfo{CourseId: courseId}}
				err := models.ChooseCourse(rel)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(rel)
			}
			break
		}
	case utils.TEACHER:
		{
			// 老师
			this.Data["json"] = utils.NoFoundReJson("此功能老师暂无...")
			break
		}
	case utils.TEACHER_HEAD:
		{
			// 系主任
			this.Data["json"] = utils.NoFoundReJson("此功能系主任暂无...")
			break
		}
	default:
		{
			this.Data["json"] = utils.NoFoundReJson("未知用户...")
		}
	}
	this.ServeJSON()
}

// CourseGrade 获取课程学生及其成绩信息  Get请求获取学生课程成绩, Post请求进行设置成绩
func (this *CourseController) CourseGrade() {
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
	case utils.TEACHER, utils.TEACHER_HEAD:
		{
			// 老师,系主任
			teacherInfo := user.User.(*models.TeacherInfo)
			if method == "GET" {
				// 获取要查询的课程号
				courseId := this.GetString("courseId")
				if courseId == "" {
					this.Data["json"] = utils.ErrorReJson("输入课程号")
					break
				}
				course := &models.CourseInfo{CourseId: courseId}
				if teacherInfo.TeacherType == utils.TEACHER {
					err := models.IsTeacherCourse(&models.CourseTeacherRel{Course: course, Teacher: teacherInfo})
					if err != nil {
						logs.Error(err)
						this.Data["json"] = utils.ErrorReJson(err.Error())
						break
					}
				}
				pageInfo := new(utils.PageInfo)
				err := models.GetGradeCourse(teacherInfo, pageInfo, course)
				if err != nil {
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson(pageInfo)
			} else if method == "POST" {
				courseStudentRel := new(utils.CourseStudentRel)
				err := utils.ParseBody(&this.Controller, courseStudentRel)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				if courseStudentRel.StudentResults == 0 || courseStudentRel.CourseId == "" || courseStudentRel.StudentId == ""{
					logs.Debug(courseStudentRel)
					this.Data["json"] = utils.ErrorReJson("输入数据必须有课程id,学号和新成绩")
					break
				}
				err = models.IsTeacherCourse(&models.CourseTeacherRel{Course: &models.CourseInfo{CourseId: courseStudentRel.CourseId}, Teacher: teacherInfo})
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				// 设置成绩
				err = models.SetStudentGradeRel(courseStudentRel)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				// 更新缓存中的学生成绩  待做
				this.Data["json"] = utils.SuccessReJson(courseStudentRel)
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

// CourseClass 获取课程班级信息,带上课程id
// Get无班级id参数获取该课程所有上课班, 有班级id进行班级统一选课
func (this *CourseController) CourseClass() {
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
	case utils.TEACHER, utils.TEACHER_HEAD:
		{
			// 老师,系主任
			teacherInfo := user.User.(*models.TeacherInfo)
			if method == "GET" {
				courseId := this.GetString("courseId")
				if courseId == "" {
					this.Data["json"] = utils.ErrorReJson("请输入课程号")
					break
				}
				course := &models.CourseInfo{CourseId: courseId}
				classId := this.GetString("classId")
				if classId == "" {
					// 获取该课程的所有班级信息
					if teacherInfo.TeacherType == utils.TEACHER {
						err := models.IsTeacherCourse(&models.CourseTeacherRel{Course: course, Teacher: teacherInfo})
						if err != nil {
							logs.Error(err)
							this.Data["json"] = utils.ErrorReJson(err.Error())
							break
						}
					}
					err := models.GetCourseClass(course)
					if err != nil {
						logs.Error(err)
						this.Data["json"] = utils.ErrorReJson(err.Error())
						break
					}
					this.Data["json"] = utils.ErrorReJson(course.Classes)
					break
				}
				// 给班级统一选课
				if teacherInfo.TeacherType == utils.TEACHER {
					this.Data["json"] = utils.ErrorReJson("你无权操作")
					break
				}
				err := models.SetClassCourse(&models.ClassInfo{ClassId: courseId}, course)
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson(err.Error())
					break
				}
				this.Data["json"] = utils.SuccessReJson("成功选上")
			} else if method == "POST" {

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

