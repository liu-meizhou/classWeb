package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func ReadClass(classId string) (*ClassInfo, error) {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// 构建查询对象
	qb.Select(GetClassColumn(), GetTeacherColumn(), GetCourseClassRelColumn(),
		GetCourseColumn(), GetStudentColumn()).
		From("class_info").
		LeftJoin("teacher_info").On("class_info.teacher_id=teacher_info.teacher_id").
		LeftJoin("course_class_rel").On("class_info.class_id=course_class_rel.class_id").
		LeftJoin("course_info").On("course_class_rel.course_id=course_info.course_id").
		LeftJoin("student_info").On("class_info.class_id=student_info.class_id").
		Where("class_info.class_id = ?")

	// 导出 SQL 语句
	sql := qb.String()

	o := orm.NewOrm()
	var maps []orm.Params
	_, err = o.Raw(sql, classId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	class := ParseClass(maps)

	return class, nil
}

func CreateClass(class *ClassInfo) error {
	o := orm.NewOrm()
	_, err := o.Insert(class)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func UpdateClass(class *ClassInfo) error {
	o := orm.NewOrm()
	_, err := o.Update(class)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

