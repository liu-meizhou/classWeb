package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func GetStudentCourse(student *StudentInfo) error {
	//sql := `select * from course_info where course_id in
	//		(select course_id from course_student_rel where student_id=?)`
	// 或者: `select * from course_info as course left join course_student_rel as rel on course.course_id=rel.course_id
	//			where rel.student_id=?`
	//o := orm.NewOrm()
	//var courses []*CourseInfo
	//num, err := o.Raw(sql, student.StudentId).QueryRows(&courses)
	//if err != nil {
	//	logs.Error(err)
	//	return err
	//}
	//logs.Info("课程查询:", num)
	//student.Courses = courses
	//logs.Info(student)
	//return nil

	o := orm.NewOrm()

	// 获取该学生的课程信息
	//var courses []*CourseInfo

	_, err := o.LoadRelated(student, "Courses")
	if err != nil {
		logs.Error(err)
		return err
	}
	for course := range student.Courses{
		courseStudentRel := &CourseStudentRel{}
		o.QueryTable("CourseStudentRel").
			Filter("Course__CourseId", student.Courses[course].CourseId).Filter("Student__StudentId", student.StudentId).
			One(courseStudentRel)
		student.Courses[course].StudentResults = courseStudentRel.StudentResults
		student.Courses[course].StudentPoint = courseStudentRel.StudentPoint
		o.LoadRelated(student.Courses[course], "CourseBases")
		o.LoadRelated(student.Courses[course], "Classes")
		o.LoadRelated(student.Courses[course], "Teachers")
		o.LoadRelated(student.Courses[course], "ClassGroups")
	}
	return nil
}