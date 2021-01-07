package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	"goweb/models"
	_ "goweb/routers"
	"os"
	"time"
)

func init() {
	// need to register db driver
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error(err)
		os.Exit(-1)
	}

	// need to register default database
	err = orm.RegisterDataBase("default", "postgres",
		"postgresql://postgres:123456@42.193.143.9:5432/postgres?sslmode=disable&&TimeZone=Asia/Shanghai")
	if err != nil {
		logs.Error(err)
		os.Exit(-1)
	}
}

func genDB() {
	// automatically build table
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		logs.Info(err)
		os.Exit(-1)
	}
	o := orm.NewOrm()

	// 初始化数据
	// 创建学生 刘佳合
	studentInfo1 := new(models.StudentInfo)
	studentInfo1.StudentId = "1865400006"
	studentInfo1.StudentName = "刘佳合"
	studentInfo1.StudentSex = "男"
	studentInfo1.StudentCollege = "计算机学院"
	studentInfo1.StudentBirth, _ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	// 创建班级 计科182
	classInfo1 := new(models.ClassInfo)
	classInfo1.ClassId = "182"
	classInfo1.ClassName = "计科182"
	// 创建老师 李传中
	teacherInfo1 := new(models.TeacherInfo)
	teacherInfo1.TeacherId = "100755"
	teacherInfo1.TeacherName = "李传中"
	teacherInfo1.TeacherSex = "男"
	teacherInfo1.TeacherCollege = "计算机学院"
	_, err = o.Insert(teacherInfo1)
	if err != nil {
		logs.Info(err)
	}
	// 创建老师 杨朔
	teacherInfo2 := new(models.TeacherInfo)
	teacherInfo2.TeacherId = "111666"
	teacherInfo2.TeacherName = "杨朔"
	teacherInfo2.TeacherSex = "男"
	teacherInfo2.TeacherCollege = "计算机学院"
	_, err = o.Insert(teacherInfo2)
	if err != nil {
		logs.Info(err)
	}
	// 创建课程1 机器学习与数据挖掘
	courseInfo1 := new(models.CourseInfo)
	courseInfo1.CourseId = "180600080"
	courseInfo1.CourseName = "机器学习与数据挖掘"
	courseInfo1.CourseProperties = "专业必修课程"
	courseInfo1.CourseScores = 3.0
	courseInfo1.CourseWay = "考试"
	courseInfo1.CourseCount = 48.0
	_, err = o.Insert(courseInfo1)
	if err != nil {
		logs.Info(err)
	}
	// 创建课程2 软件工程导论
	courseInfo2 := new(models.CourseInfo)
	courseInfo2.CourseId = "180600019"
	courseInfo2.CourseName = "软件工程导论"
	courseInfo2.CourseProperties = "专业选修课程"
	courseInfo2.CourseScores = 2.0
	courseInfo2.CourseWay = "考试"
	courseInfo2.CourseCount = 32
	_, err = o.Insert(courseInfo2)
	if err != nil {
		logs.Info(err)
	}
	// 创建上课基本信息1 机器学习与数据挖掘
	courseBaseInfo1 := new(models.CourseBaseInfo)
	courseBaseInfo1.Course = courseInfo1
	courseBaseInfo1.CourseStartYear = 2020
	courseBaseInfo1.CourseEndYear = 2021
	courseBaseInfo1.CourseYear = 1
	courseBaseInfo1.CourseStartWeek = 1
	courseBaseInfo1.CourseEndWeek = 16
	courseBaseInfo1.CourseWeek = 3
	courseBaseInfo1.CourseStartCount = 1
	courseBaseInfo1.CourseEndCount = 2
	courseBaseInfo1.CourseSchool = "大学城"
	courseBaseInfo1.CourseAddress = "理科南"
	courseBaseInfo1.CourseAddressFloor = 3
	courseBaseInfo1.CourseAddressNumber = 15
	_, err = o.Insert(courseBaseInfo1)
	if err != nil {
		logs.Info(err)
	}
	// 创建上课基本信息2 机器学习与数据挖掘
	courseBaseInfo2 := new(models.CourseBaseInfo)
	courseBaseInfo2.Course = courseInfo1
	courseBaseInfo2.CourseStartYear = 2020
	courseBaseInfo2.CourseEndYear = 2021
	courseBaseInfo2.CourseYear = 1
	courseBaseInfo2.CourseStartWeek = 1
	courseBaseInfo2.CourseEndWeek = 8
	courseBaseInfo2.CourseWeek = 4
	courseBaseInfo2.CourseStartCount = 1
	courseBaseInfo2.CourseEndCount = 2
	courseBaseInfo2.CourseSchool = "大学城"
	courseBaseInfo2.CourseAddress = "理科南"
	courseBaseInfo2.CourseAddressFloor = 6
	courseBaseInfo2.CourseAddressNumber = 14
	_, err = o.Insert(courseBaseInfo2)
	if err != nil {
		logs.Info(err)
	}
	// 创建上课基本信息3 软件工程导论
	courseBaseInfo3 := new(models.CourseBaseInfo)
	courseBaseInfo3.Course = courseInfo2
	courseBaseInfo3.CourseStartYear = 2020
	courseBaseInfo3.CourseEndYear = 2021
	courseBaseInfo3.CourseYear = 1
	courseBaseInfo3.CourseStartWeek = 1
	courseBaseInfo3.CourseEndWeek = 10
	courseBaseInfo3.CourseWeek = 4
	courseBaseInfo3.CourseStartCount = 9
	courseBaseInfo3.CourseEndCount = 11
	courseBaseInfo3.CourseSchool = "大学城"
	courseBaseInfo3.CourseAddress = "理科南"
	courseBaseInfo3.CourseAddressFloor = 1
	courseBaseInfo3.CourseAddressNumber = 11
	_, err = o.Insert(courseBaseInfo3)
	if err != nil {
		logs.Info(err)
	}
	// 创建课程组1 计科
	classGroupInfo1 := new(models.ClassGroupInfo)
	classGroupInfo1.ClassGroupName = "计科"
	_, err = o.Insert(classGroupInfo1)
	if err != nil {
		logs.Info(err)
	}
	// 创建学生课程关系1 刘佳合  软件工程导论
	courseStudentRel1 := new(models.CourseStudentRel)
	courseStudentRel1.Student = studentInfo1
	courseStudentRel1.Course = courseInfo2
	_, err = o.Insert(courseStudentRel1)
	if err != nil {
		logs.Info(err)
	}
	// 创建学生课程关系2 刘佳合  机器学习与数据挖掘
	courseStudentRel2 := new(models.CourseStudentRel)
	courseStudentRel2.Student = studentInfo1
	courseStudentRel2.Course = courseInfo1
	_, err = o.Insert(courseStudentRel2)
	if err != nil {
		logs.Info(err)
	}
	// 创建课程组和老师关系  李传中和杨朔在 计科课组
	classGroupTeacherRel1 := new(models.ClassGroupTeacherRel)
	classGroupTeacherRel1.ClassGroup = classGroupInfo1
	classGroupTeacherRel1.Teacher = teacherInfo1
	classGroupTeacherRel1.IsCharge = true
	classGroupTeacherRel2 := new(models.ClassGroupTeacherRel)
	classGroupTeacherRel2.ClassGroup = classGroupInfo1
	classGroupTeacherRel2.Teacher = teacherInfo2
	classGroupTeacherRel2.IsCharge = false
	_, err = o.Insert(classGroupTeacherRel1)
	if err != nil {
		logs.Info(err)
	}
	_, err = o.Insert(classGroupTeacherRel2)
	if err != nil {
		logs.Info(err)
	}
	// 计科182都选了机器学习课程
	courseClassRel1 := new(models.CourseClassRel)
	courseClassRel1.Course = courseInfo1
	courseClassRel1.Class = classInfo1
	_, err = o.Insert(courseClassRel1)
	if err != nil {
		logs.Info(err)
	}
	// 课程老师联系表1 李传中教软件工程导论
	courseTeacherRel1 := new(models.CourseTeacherRel)
	courseTeacherRel1.Course = courseInfo2
	courseTeacherRel1.Teacher = teacherInfo1
	_, err = o.Insert(courseTeacherRel1)
	if err != nil {
		logs.Info(err)
	}
	// 课程老师联系表2 杨朔教机器学习与数据挖掘
	courseTeacherRel2 := new(models.CourseTeacherRel)
	courseTeacherRel2.Course = courseInfo1
	courseTeacherRel2.Teacher = teacherInfo2
	_, err = o.Insert(courseTeacherRel2)
	if err != nil {
		logs.Info(err)
	}
	// 机器学习与数据挖掘和软件工程导论在同个课组
	courseGroupRel1 := new(models.CourseGroupRel)
	courseGroupRel1.ClassGroup = classGroupInfo1
	courseGroupRel1.Course = courseInfo1
	courseGroupRel2 := new(models.CourseGroupRel)
	courseGroupRel2.ClassGroup = classGroupInfo1
	courseGroupRel2.Course = courseInfo2
	_, err = o.Insert(courseGroupRel1)
	if err != nil {
		logs.Info(err)
	}
	_, err = o.Insert(courseGroupRel2)
	if err != nil {
		logs.Info(err)
	}

	classInfo1.Teacher = teacherInfo1
	_, err = o.Insert(classInfo1)
	if err != nil {
		logs.Info(err)
	}
	studentInfo1.Class = classInfo1
	_, err = o.Insert(studentInfo1)
	if err != nil {
		logs.Info(err)
	}
}

func main() {
	// 配置运行跨域，配置不成功0.0, 改为配置前端
	// 前端配置：https://www.jianshu.com/p/43aa317d7683
	//web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowOrigins: []string{"https://*.foo.com"},
	//	AllowMethods: []string{"GET", "POST", "PUT", "PATCH"},
	//	AllowHeaders: []string{"Origin"},
	//	ExposeHeaders: []string{"Content-Length"},
	//	AllowCredentials: true,
	//}))

	//web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
	//	//AllowAllOrigins:  true,
	//	AllowMethods:     []string{"*"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type"},
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
	//	AllowCredentials: true,
	//	AllowOrigins: []string{"http://10.*.*.*:*","http://localhost:*","http://127.0.0.1:*"},
	//}))

	// 生成数据库
	orm.Debug = true
	//genDB()

	// 启动web
	web.Run()
}
