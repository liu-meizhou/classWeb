package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// GetTeacherCourse
func GetTeacherCourse(teacher *TeacherInfo) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}

	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseTeacherRelColumn(), GetTeacherColumn(), GetCourseClassRelColumn(),
		GetClassColumn(), GetCourseGroupRelColumn(), GetClassGroupColumn(), GetCourseStudentRelColumn()).
		From("course_info").
		LeftJoin("course_base_info").On("course_base_info.course_id=course_info.course_id").
		InnerJoin("course_teacher_rel").On("course_info.course_id=course_teacher_rel.course_id").
		LeftJoin("teacher_info").On("teacher_info.teacher_id=course_teacher_rel.teacher_id").
		LeftJoin("course_class_rel").On("course_class_rel.course_id=course_info.course_id").
		LeftJoin("class_info").On("class_info.class_id=course_class_rel.class_id").
		LeftJoin("course_group_rel").On("course_group_rel.course_id=course_info.course_id").
		LeftJoin("class_group_info").On("class_group_info.class_group_id=course_group_rel.class_group_id").
		LeftJoin("course_student_rel").On("course_info.course_id=course_student_rel.course_id").
		Where("course_teacher_rel.teacher_id = ?")

	// 导出 SQL 语句
	sql := qb.String()
	//logs.Debug(sql)

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(sql, teacher.TeacherId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return err
	}
	logs.Info(num)

	courses := ParseCourses(maps)
	teacher.Courses = courses
	return nil
}

// GetGradeCourse
func GetGradeCourse(course *CourseInfo) ([]*StudentInfo, error) {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	qb.Select(GetCourseStudentRelColumn(), GetStudentColumn(), GetClassColumn()).
		From("course_student_rel").
		RightJoin("student_info").On("course_student_rel.student_id=student_info.student_id").
		RightJoin("class_info").On("student_info.class_id=class_info.class_id").
		Where("course_student_rel.course_id = ?")

	// 导出 SQL 语句
	sql := qb.String()
	//logs.Debug(sql)

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(sql, course.CourseId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	logs.Info(num)

	students := ParseCourseStudent(maps)
	return students, nil
}
