package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"goweb/models"
	"goweb/utils"
)

type CourseController struct {
	web.Controller
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
				this.Data["json"] = utils.ErrorReJson(err)
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

// ChooseCourse 选课  学生: Get请求无参数获取选课列表 带courseId参数请求进行选课
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
			// 获取请求参数
			courseId := this.GetString("courseId")
			if courseId == "" {
				courses, err := models.GetChooseCourse(studentInfo)
				if err != nil {
					this.Data["json"] = utils.ErrorReJson(err)
					break
				}
				this.Data["json"] = utils.SuccessReJson(courses)
			} else {
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

// CourseGrade 获取课程成绩等信息  Get请求获取学生课程成绩, Post请求进行设置成绩
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
						this.Data["json"] = utils.ErrorReJson("你无权限查询")
						break
					}
				}
				students, err := models.GetGradeCourse(course)
				if err != nil {
					this.Data["json"] = utils.ErrorReJson(err)
					break
				}
				this.Data["json"] = utils.SuccessReJson(students)
			} else if method == "POST" {
				body := this.Ctx.Input.RequestBody
				if len(body)==0 {
					this.Data["json"] = utils.ErrorReJson("传入不可为空")
					break
				}
				courseStudentRel := new(utils.CourseStudentRel)
				err := json.Unmarshal(body, courseStudentRel)
				if err!=nil {
					err2 := this.ParseForm(courseStudentRel)
					if err2!=nil {
						err2 = fmt.Errorf(err.Error() + "   " + err2.Error())
						logs.Error(err2)
						this.Data["json"] = utils.ErrorReJson(err2.Error())
						break
					}
				}
				if courseStudentRel.StudentResults == 0 || courseStudentRel.CourseId == "" || courseStudentRel.StudentId == ""{
					logs.Debug(courseStudentRel)
					this.Data["json"] = utils.ErrorReJson("输入数据必须有课程id,学号和新成绩")
					break
				}
				err = models.IsTeacherCourse(&models.CourseTeacherRel{Course: &models.CourseInfo{CourseId: courseStudentRel.CourseId}, Teacher: teacherInfo})
				if err != nil {
					logs.Error(err)
					this.Data["json"] = utils.ErrorReJson("你无权限查询")
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

