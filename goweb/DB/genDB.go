package DB

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"goweb/models"
	"goweb/utils"
	"os"
	"time"
)

func GenTable()  {
	// automatically build table
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		logs.Info(err)
		os.Exit(-1)
	}
}

func GenData() (err error) {
	//orm.Debug = true
	teachers := getTeachers()
	for _, teacher := range teachers {
		err = models.CreateTeacher(teacher)
		if err != nil {
			logs.Error(err)
		}
	}

	//orm.Debug = true
	classes := getClasses()
	for _, class := range classes {
		class.Teacher = teachers[0]
		err = models.CreateClass(class)
		if err != nil {
			logs.Error(err)
		}
	}

	students := getStudents()
	for _, student := range students {
		student.Class = classes[0]
		err = models.CreateStudent(student)
		if err != nil {
			logs.Error(err)
		}
	}

	courses := getCourses()
	for _, course := range courses {
		err = models.CreateCourse(course)
		if err != nil {
			logs.Error(err)
		}
	}

	// 给学生0选选修课1
	courseStudentRel := getCourseStudentRel(courses[1], students[0])
	err = models.ChooseCourse(courseStudentRel)
	if err != nil {
		logs.Error(err)
	}

	// 给老师0分配选修课1
	courseTeacherRel := getCourseTeacherRel(courses[1], teachers[0])
	err = models.TeacherChooseCourse(courseTeacherRel)
	if err != nil {
		logs.Error(err)
	}

	// 给班级0统一选必修课0
	err = models.SetClassCourse(classes[0], courses[0])
	if err != nil {
		logs.Error(err)
	}

	return nil
}


func getClasses() []*models.ClassInfo {
	var classes []*models.ClassInfo
	// 创建班级 计科181
	classInfo1 := new(models.ClassInfo)
	classInfo1.ClassId = "1"
	classInfo1.ClassName = "计科181"
	classes = append(classes, classInfo1)
	// 创建班级 计科182
	classInfo2 := new(models.ClassInfo)
	classInfo2.ClassId = "2"
	classInfo2.ClassName = "计科182"
	classes = append(classes, classInfo2)
	return classes
}

func getStudents() []*models.StudentInfo {
	var students []*models.StudentInfo
	// 初始化数据
	// 创建学生1
	studentInfo1 := new(models.StudentInfo)
	studentInfo1.StudentId = "1"
	studentInfo1.StudentPassword = "123456"
	studentInfo1.StudentType = utils.STUDENT
	studentInfo1.StudentName = "学生1号"
	studentInfo1.StudentSex = "男"
	studentInfo1.StudentCollege = "计算机学院"
	studentInfo1.StudentBirth, _ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	studentInfo1.StudentTime,_ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	students = append(students, studentInfo1)

	// 创建学生2
	studentInfo2 := new(models.StudentInfo)
	studentInfo2.StudentId = "2"
	studentInfo2.StudentPassword = "123456"
	studentInfo2.StudentType = utils.STUDENT
	studentInfo2.StudentName = "学生2号"
	studentInfo2.StudentSex = "男"
	studentInfo2.StudentCollege = "计算机学院"
	studentInfo2.StudentBirth, _ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	studentInfo2.StudentTime,_ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	students = append(students, studentInfo2)

	// 创建学生3
	studentInfo3 := new(models.StudentInfo)
	studentInfo3.StudentId = "3"
	studentInfo3.StudentPassword = "123456"
	studentInfo3.StudentType = utils.STUDENT
	studentInfo3.StudentName = "学生3号"
	studentInfo3.StudentSex = "男"
	studentInfo3.StudentCollege = "计算机学院"
	studentInfo3.StudentBirth, _ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	studentInfo3.StudentTime,_ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	students = append(students, studentInfo3)

	return students
}

func getTeachers() []*models.TeacherInfo {
	var teachers []*models.TeacherInfo
	// 初始化数据
	// 创建老师1
	teacherInfo1 := new(models.TeacherInfo)
	teacherInfo1.TeacherId = "1"
	teacherInfo1.TeacherPassword = "123456"
	teacherInfo1.TeacherType = utils.TEACHER_HEAD
	teacherInfo1.TeacherName = "老师1号"
	teacherInfo1.TeacherSex = "男"
	teacherInfo1.TeacherCollege = "计算机学院"
	teacherInfo1.TeacherBirth, _ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	teacherInfo1.TeacherTime,_ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	teachers = append(teachers, teacherInfo1)

	// 创建老师3
	teacherInfo2 := new(models.TeacherInfo)
	teacherInfo2.TeacherId = "2"
	teacherInfo2.TeacherPassword = "123456"
	teacherInfo2.TeacherType = utils.TEACHER
	teacherInfo2.TeacherName = "老师2号"
	teacherInfo2.TeacherSex = "男"
	teacherInfo2.TeacherCollege = "计算机学院"
	teacherInfo2.TeacherBirth, _ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	teacherInfo2.TeacherTime,_ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	teachers = append(teachers, teacherInfo2)

	// 创建老师3
	teacherInfo3 := new(models.TeacherInfo)
	teacherInfo3.TeacherId = "3"
	teacherInfo3.TeacherPassword = "123456"
	teacherInfo3.TeacherType = utils.TEACHER
	teacherInfo3.TeacherName = "老师3号"
	teacherInfo3.TeacherSex = "男"
	teacherInfo3.TeacherCollege = "计算机学院"
	teacherInfo3.TeacherBirth, _ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	teacherInfo3.TeacherTime,_ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
	teachers = append(teachers, teacherInfo3)

	return teachers
}

func getCourses() []*models.CourseInfo {
	var courses []*models.CourseInfo
	// 创建课程1
	courseInfo1 := new(models.CourseInfo)
	courseInfo1.CourseId = "1"
	courseInfo1.CourseName = "课程1"
	courseInfo1.CourseProperties = "必修"
	courseInfo1.CourseScores = 2
	courseInfo1.CourseWay = "考试"
	courseInfo1.CourseCount = 32
	courseBaseInfo11 := new(models.CourseBaseInfo)
	courseBaseInfo11.CourseStartYear = 2020
	courseBaseInfo11.CourseEndYear = 2021
	courseBaseInfo11.CourseYear = 2020
	courseBaseInfo11.CourseStartWeek = 1
	courseBaseInfo11.CourseEndWeek = 16
	courseBaseInfo11.CourseWeek = 3
	courseBaseInfo11.CourseStartCount = 1
	courseBaseInfo11.CourseEndCount = 2
	courseBaseInfo11.CourseSchool = "大学城"
	courseBaseInfo11.CourseAddress = "理科南"
	courseBaseInfo11.CourseAddressFloor = 2
	courseBaseInfo11.CourseAddressNumber = 11
	courseInfo1.CourseBases = append(courseInfo1.CourseBases, courseBaseInfo11)
	courseBaseInfo12 := new(models.CourseBaseInfo)
	courseBaseInfo12.CourseStartYear = 2020
	courseBaseInfo12.CourseEndYear = 2021
	courseBaseInfo12.CourseYear = 2020
	courseBaseInfo12.CourseStartWeek = 1
	courseBaseInfo12.CourseEndWeek = 16
	courseBaseInfo12.CourseWeek = 5
	courseBaseInfo12.CourseStartCount = 3
	courseBaseInfo12.CourseEndCount = 4
	courseBaseInfo12.CourseSchool = "大学城"
	courseBaseInfo12.CourseAddress = "理科南"
	courseBaseInfo12.CourseAddressFloor = 4
	courseBaseInfo12.CourseAddressNumber = 12
	courseInfo1.CourseBases = append(courseInfo1.CourseBases, courseBaseInfo12)
	courses = append(courses, courseInfo1)

	// 创建课程2
	courseInfo2 := new(models.CourseInfo)
	courseInfo2.CourseId = "2"
	courseInfo2.CourseName = "课程2"
	courseInfo2.CourseProperties = "选修"
	courseInfo2.CourseScores = 2
	courseInfo2.CourseWay = "考察"
	courseInfo2.CourseCount = 32
	courseBaseInfo21 := new(models.CourseBaseInfo)
	courseBaseInfo21.CourseStartYear = 2020
	courseBaseInfo21.CourseEndYear = 2021
	courseBaseInfo21.CourseYear = 2020
	courseBaseInfo21.CourseStartWeek = 1
	courseBaseInfo21.CourseEndWeek = 16
	courseBaseInfo21.CourseWeek = 4
	courseBaseInfo21.CourseStartCount = 1
	courseBaseInfo21.CourseEndCount = 2
	courseBaseInfo21.CourseSchool = "大学城"
	courseBaseInfo21.CourseAddress = "理科南"
	courseBaseInfo21.CourseAddressFloor = 2
	courseBaseInfo21.CourseAddressNumber = 11
	courseInfo2.CourseBases = append(courseInfo2.CourseBases, courseBaseInfo21)
	courses = append(courses, courseInfo2)

	// 创建课程3
	courseInfo3 := new(models.CourseInfo)
	courseInfo3.CourseId = "3"
	courseInfo3.CourseName = "课程3"
	courseInfo3.CourseProperties = "选修"
	courseInfo3.CourseScores = 2
	courseInfo3.CourseWay = "考察"
	courseInfo3.CourseCount = 32
	courseBaseInfo31 := new(models.CourseBaseInfo)
	courseBaseInfo31.CourseStartYear = 2020
	courseBaseInfo31.CourseEndYear = 2021
	courseBaseInfo31.CourseYear = 2020
	courseBaseInfo31.CourseStartWeek = 1
	courseBaseInfo31.CourseEndWeek = 16
	courseBaseInfo31.CourseWeek = 3
	courseBaseInfo31.CourseStartCount = 1
	courseBaseInfo31.CourseEndCount = 2
	courseBaseInfo31.CourseSchool = "大学城"
	courseBaseInfo31.CourseAddress = "理科南"
	courseBaseInfo31.CourseAddressFloor = 2
	courseBaseInfo31.CourseAddressNumber = 11
	courseInfo3.CourseBases = append(courseInfo3.CourseBases, courseBaseInfo31)
	courses = append(courses, courseInfo3)

	return courses
}

func getCourseGroups() []*models.CourseGroupInfo {
	// 创建课程组1 计科
	courseGroupInfo1 := new(models.CourseGroupInfo)
	courseGroupInfo1.CourseGroupName = "计科"
	return []*models.CourseGroupInfo{courseGroupInfo1}
}

func getCourseStudentRel(course *models.CourseInfo, student *models.StudentInfo) *models.CourseStudentRel {
	// 创建学生课程关系
	courseStudentRel := new(models.CourseStudentRel)
	courseStudentRel.Student = student
	courseStudentRel.Course = course
	return courseStudentRel
}

func getCourseTeacherRel(course *models.CourseInfo, teacher *models.TeacherInfo) *models.CourseTeacherRel {
	// 创建课程老师关系
	courseTeacherRel := new(models.CourseTeacherRel)
	courseTeacherRel.Teacher = teacher
	courseTeacherRel.Course = course
	return courseTeacherRel
}

func getCourseClassRel(course *models.CourseInfo, class *models.ClassInfo) *models.CourseClassRel {
	// 创建课程班级关系
	courseClassRel := new(models.CourseClassRel)
	courseClassRel.Class = class
	courseClassRel.Course = course
	return courseClassRel
}




//if !isData {
//	return
//}
//o := orm.NewOrm()
//
//// 初始化数据
//// 创建学生 刘佳合
//studentInfo1 := new(models.StudentInfo)
//studentInfo1.StudentId = "1865400006"
//studentInfo1.StudentPassword = "123456"
//studentInfo1.StudentType = utils.STUDENT
//studentInfo1.StudentName = "刘佳合"
//studentInfo1.StudentSex = "男"
//studentInfo1.StudentCollege = "计算机学院"
//studentInfo1.StudentBirth, _ = time.Parse("2006-01-02 15:04:05", "2000-01-22 12:00:00")
//// 创建班级 计科182
//classInfo1 := new(models.ClassInfo)
//classInfo1.ClassId = "182"
//classInfo1.ClassName = "计科182"
//// 创建老师 李传中
//teacherInfo1 := new(models.TeacherInfo)
//teacherInfo1.TeacherId = "100755"
//teacherInfo1.TeacherPassword = "123456"
//teacherInfo1.TeacherType = utils.TEACHER_HEAD
//teacherInfo1.TeacherName = "李传中"
//teacherInfo1.TeacherSex = "男"
//teacherInfo1.TeacherCollege = "计算机学院"
//_, err = o.Insert(teacherInfo1)
//if err != nil {
//	logs.Info(err)
//}
//// 创建老师 杨朔
//teacherInfo2 := new(models.TeacherInfo)
//teacherInfo2.TeacherId = "111666"
//teacherInfo2.TeacherPassword = "123456"
//teacherInfo2.TeacherType = utils.TEACHER
//teacherInfo2.TeacherName = "杨朔"
//teacherInfo2.TeacherSex = "男"
//teacherInfo2.TeacherCollege = "计算机学院"
//_, err = o.Insert(teacherInfo2)
//if err != nil {
//	logs.Info(err)
//}
//// 创建课程1 机器学习与数据挖掘
//courseInfo1 := new(models.CourseInfo)
//courseInfo1.CourseId = "180600080"
//courseInfo1.CourseName = "机器学习与数据挖掘"
//courseInfo1.CourseProperties = "专业必修课程"
//courseInfo1.CourseScores = 3.0
//courseInfo1.CourseWay = "考试"
//courseInfo1.CourseCount = 48.0
//_, err = o.Insert(courseInfo1)
//if err != nil {
//	logs.Info(err)
//}
//// 创建课程2 软件工程导论
//courseInfo2 := new(models.CourseInfo)
//courseInfo2.CourseId = "180600019"
//courseInfo2.CourseName = "软件工程导论"
//courseInfo2.CourseProperties = "专业选修课程"
//courseInfo2.CourseScores = 2.0
//courseInfo2.CourseWay = "考试"
//courseInfo2.CourseCount = 32
//_, err = o.Insert(courseInfo2)
//if err != nil {
//	logs.Info(err)
//}
//// 创建上课基本信息1 机器学习与数据挖掘
//courseBaseInfo1 := new(models.CourseBaseInfo)
//courseBaseInfo1.Course = courseInfo1
//courseBaseInfo1.CourseStartYear = 2020
//courseBaseInfo1.CourseEndYear = 2021
//courseBaseInfo1.CourseYear = 1
//courseBaseInfo1.CourseStartWeek = 1
//courseBaseInfo1.CourseEndWeek = 16
//courseBaseInfo1.CourseWeek = 3
//courseBaseInfo1.CourseStartCount = 1
//courseBaseInfo1.CourseEndCount = 2
//courseBaseInfo1.CourseSchool = "大学城"
//courseBaseInfo1.CourseAddress = "理科南"
//courseBaseInfo1.CourseAddressFloor = 3
//courseBaseInfo1.CourseAddressNumber = 15
//_, err = o.Insert(courseBaseInfo1)
//if err != nil {
//	logs.Info(err)
//}
//// 创建上课基本信息2 机器学习与数据挖掘
//courseBaseInfo2 := new(models.CourseBaseInfo)
//courseBaseInfo2.Course = courseInfo1
//courseBaseInfo2.CourseStartYear = 2020
//courseBaseInfo2.CourseEndYear = 2021
//courseBaseInfo2.CourseYear = 1
//courseBaseInfo2.CourseStartWeek = 1
//courseBaseInfo2.CourseEndWeek = 8
//courseBaseInfo2.CourseWeek = 4
//courseBaseInfo2.CourseStartCount = 1
//courseBaseInfo2.CourseEndCount = 2
//courseBaseInfo2.CourseSchool = "大学城"
//courseBaseInfo2.CourseAddress = "理科南"
//courseBaseInfo2.CourseAddressFloor = 6
//courseBaseInfo2.CourseAddressNumber = 14
//_, err = o.Insert(courseBaseInfo2)
//if err != nil {
//	logs.Info(err)
//}
//// 创建上课基本信息3 软件工程导论
//courseBaseInfo3 := new(models.CourseBaseInfo)
//courseBaseInfo3.Course = courseInfo2
//courseBaseInfo3.CourseStartYear = 2020
//courseBaseInfo3.CourseEndYear = 2021
//courseBaseInfo3.CourseYear = 1
//courseBaseInfo3.CourseStartWeek = 1
//courseBaseInfo3.CourseEndWeek = 10
//courseBaseInfo3.CourseWeek = 4
//courseBaseInfo3.CourseStartCount = 9
//courseBaseInfo3.CourseEndCount = 11
//courseBaseInfo3.CourseSchool = "大学城"
//courseBaseInfo3.CourseAddress = "理科南"
//courseBaseInfo3.CourseAddressFloor = 1
//courseBaseInfo3.CourseAddressNumber = 11
//_, err = o.Insert(courseBaseInfo3)
//if err != nil {
//	logs.Info(err)
//}
//// 创建课程组1 计科
//classGroupInfo1 := new(models.CourseGroupInfo)
//classGroupInfo1.CourseGroupName = "计科"
//_, err = o.Insert(classGroupInfo1)
//if err != nil {
//	logs.Info(err)
//}
//// 创建学生课程关系1 刘佳合  软件工程导论
//courseStudentRel1 := new(models.CourseStudentRel)
//courseStudentRel1.Student = studentInfo1
//courseStudentRel1.Course = courseInfo2
//_, err = o.Insert(courseStudentRel1)
//if err != nil {
//	logs.Info(err)
//}
//// 创建学生课程关系2 刘佳合  机器学习与数据挖掘
//courseStudentRel2 := new(models.CourseStudentRel)
//courseStudentRel2.Student = studentInfo1
//courseStudentRel2.Course = courseInfo1
//_, err = o.Insert(courseStudentRel2)
//if err != nil {
//	logs.Info(err)
//}
//// 创建课程组和老师关系  李传中和杨朔在 计科课组
//classGroupTeacherRel1 := new(models.CourseGroupTeacherRel)
//classGroupTeacherRel1.CourseGroup = classGroupInfo1
//classGroupTeacherRel1.Teacher = teacherInfo1
//classGroupTeacherRel1.IsCharge = true
//classGroupTeacherRel2 := new(models.CourseGroupTeacherRel)
//classGroupTeacherRel2.CourseGroup = classGroupInfo1
//classGroupTeacherRel2.Teacher = teacherInfo2
//classGroupTeacherRel2.IsCharge = false
//_, err = o.Insert(classGroupTeacherRel1)
//if err != nil {
//	logs.Info(err)
//}
//_, err = o.Insert(classGroupTeacherRel2)
//if err != nil {
//	logs.Info(err)
//}
//// 计科182都选了机器学习课程
//courseClassRel1 := new(models.CourseClassRel)
//courseClassRel1.Course = courseInfo1
//courseClassRel1.Class = classInfo1
//_, err = o.Insert(courseClassRel1)
//if err != nil {
//	logs.Info(err)
//}
//// 课程老师联系表1 李传中教软件工程导论
//courseTeacherRel1 := new(models.CourseTeacherRel)
//courseTeacherRel1.Course = courseInfo2
//courseTeacherRel1.Teacher = teacherInfo1
//_, err = o.Insert(courseTeacherRel1)
//if err != nil {
//	logs.Info(err)
//}
//// 课程老师联系表2 杨朔教机器学习与数据挖掘
//courseTeacherRel2 := new(models.CourseTeacherRel)
//courseTeacherRel2.Course = courseInfo1
//courseTeacherRel2.Teacher = teacherInfo2
//_, err = o.Insert(courseTeacherRel2)
//if err != nil {
//	logs.Info(err)
//}
//// 机器学习与数据挖掘和软件工程导论在同个课组
//courseGroupRel1 := new(models.CourseGroupRel)
//courseGroupRel1.CourseGroup = classGroupInfo1
//courseGroupRel1.Course = courseInfo1
//courseGroupRel2 := new(models.CourseGroupRel)
//courseGroupRel2.CourseGroup = classGroupInfo1
//courseGroupRel2.Course = courseInfo2
//_, err = o.Insert(courseGroupRel1)
//if err != nil {
//	logs.Info(err)
//}
//_, err = o.Insert(courseGroupRel2)
//if err != nil {
//	logs.Info(err)
//}
//
//classInfo1.Teacher = teacherInfo1
//_, err = o.Insert(classInfo1)
//if err != nil {
//	logs.Info(err)
//}
//studentInfo1.Class = classInfo1
//_, err = o.Insert(studentInfo1)
//if err != nil {
//	logs.Info(err)
//}

