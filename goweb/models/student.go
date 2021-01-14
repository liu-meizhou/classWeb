package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"goweb/utils"
)

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

//o := orm.NewOrm()
//
//// 获取该学生的课程信息
////var courses []*CourseInfo
//
//_, err := o.LoadRelated(student, "Courses")
//if err != nil {
//	logs.Error(err)
//	return err
//}
//for course := range student.Courses{
//	courseStudentRel := &CourseStudentRel{}
//	o.QueryTable("CourseStudentRel").
//		Filter("Course__CourseId", student.Courses[course].CourseId).Filter("Student__StudentId", student.StudentId).
//		One(courseStudentRel)
//	student.Courses[course].StudentResults = courseStudentRel.StudentResults
//	student.Courses[course].StudentPoint = courseStudentRel.StudentPoint
//	o.LoadRelated(student.Courses[course], "CourseBases")
//	o.LoadRelated(student.Courses[course], "Classes")
//	o.LoadRelated(student.Courses[course], "Teachers")
//	o.LoadRelated(student.Courses[course], "CourseGroups")
//}
//return nil
const canChooseCourse = "选修"

func ReadStudent(studentId string) (*StudentInfo, error) {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// 构建查询对象
	qb.Select(GetStudentColumn(), GetClassColumn(), GetCourseStudentRelColumn(false, false), GetCourseColumn()).
		From("student_info").
		LeftJoin("class_info").On("class_info.class_id=student_info.class_id").
		LeftJoin("course_student_rel").On("course_student_rel.student_id=student_info.student_id").
		LeftJoin("course_info").On("course_info.course_id=course_student_rel.course_id").
		Where("student_info.student_id = ?")

	// 导出 SQL 语句
	sql := qb.String()

	o := orm.NewOrm()
	var maps []orm.Params
	_, err = o.Raw(sql, studentId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	student := ParseStudent(maps)

	return student, nil
}

func CreateStudent(student *StudentInfo) error {
	o := orm.NewOrm()
	_, err := o.Insert(student)
	if err != nil && err.Error() != "<Ormer> last insert id is unavailable"  {
		logs.Error(err)
		return err
	}
	return nil
}

func UpdateStudent(student *StudentInfo) error {
	o := orm.NewOrm()
	_, err := o.Update(student)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func DeleteStudent(student *StudentInfo) error {
	_, err := orm.NewOrm().Delete(student)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func UpdateStudentClass(studentId, classId string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("student_info").Filter("student_id", studentId).Update(orm.Params{
		"class_id": classId,
	})
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func GetStudentCourse(student *StudentInfo) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}

	// 构建查询对象
	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseTeacherRelColumn(false, false),
		GetTeacherColumn(), GetCourseClassRelColumn(false, false),
		GetClassColumn(), GetCourseGroupRelColumn(false, false),
		GetCourseGroupColumn(), GetCourseStudentRelColumn(false, true)).
		From("course_info").
		LeftJoin("course_base_info").On("course_base_info.course_id=course_info.course_id").
		LeftJoin("course_teacher_rel").On("course_teacher_rel.course_id=course_info.course_id").
		LeftJoin("teacher_info").On("teacher_info.teacher_id=course_teacher_rel.teacher_id").
		LeftJoin("course_class_rel").On("course_class_rel.course_id=course_info.course_id").
		LeftJoin("class_info").On("class_info.class_id=course_class_rel.class_id").
		LeftJoin("course_group_rel").On("course_group_rel.course_id=course_info.course_id").
		LeftJoin("course_group_info").On("course_group_info.course_group_id=course_group_rel.course_group_id").
		InnerJoin("course_student_rel").On("course_info.course_id = course_student_rel.course_id").
		Where("course_student_rel.student_id = ?")

	// 导出 SQL 语句
	sql := qb.String()

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(sql, student.StudentId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return err
	}
	logs.Info(num)

	courses := ParseCourses(maps)
	student.Courses = courses
	return nil
}

func GetStudentList(pageInfo *utils.PageInfo, student *StudentInfo) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}

	// 构建查询对象
	qb.Select(GetStudentColumn(), GetClassColumn()).
		From("student_info").
		LeftJoin("class_info").On("student_info.class_id=class_info.class_id")

	// 导出 SQL 语句
	sql := qb.String()

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(sql).Values(&maps)
	if err != nil {
		logs.Error(err)
		return err
	}
	logs.Info(num)

	students := ParseStudents(maps)
	pageInfo.Lists = students
	return nil
}
