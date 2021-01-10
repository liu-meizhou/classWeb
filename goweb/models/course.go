package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func ReadCourse(courseId string) (*CourseInfo, error) {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// 构建查询对象
	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseTeacherRelColumn(), GetTeacherColumn(), GetCourseClassRelColumn(),
		GetClassColumn(), GetCourseGroupRelColumn(), GetCourseGroupColumn(), GetCourseStudentRelColumn(), GetStudentColumn()).
		From("course_info").
		LeftJoin("course_base_info").On("course_base_info.course_id=course_info.course_id").
		LeftJoin("course_teacher_rel").On("course_teacher_rel.course_id=course_info.course_id").
		LeftJoin("teacher_info").On("teacher_info.teacher_id=course_teacher_rel.teacher_id").
		LeftJoin("course_class_rel").On("course_class_rel.course_id=course_info.course_id").
		LeftJoin("class_info").On("class_info.class_id=course_class_rel.class_id").
		LeftJoin("course_group_rel").On("course_group_rel.course_id=course_info.course_id").
		LeftJoin("course_group_info").On("course_group_info.course_group_id=course_group_rel.course_group_id").
		LeftJoin("course_student_rel").On("course_info.course_id = course_student_rel.course_id").
		LeftJoin("student_info").On("course_student_rel.student_id = student_info.student_id").
		Where("course_info.course_id = ?")

	// 导出 SQL 语句
	sql := qb.String()

	o := orm.NewOrm()
	var maps []orm.Params
	_, err = o.Raw(sql, courseId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	course := ParseCourse(maps)

	return course, nil
}

func CreateCourse(course *CourseInfo) error {
	o := orm.NewOrm()
	_, err := o.Insert(course)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func UpdateCourse(course *CourseInfo) error {
	o := orm.NewOrm()
	_, err := o.Update(course)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func GetCourseClass(course *CourseInfo) error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(course, "Classes")
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
