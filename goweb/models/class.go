package models

import (
	"fmt"
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
	if err != nil && err.Error() != "<Ormer> last insert id is unavailable" {
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

func GetClassList() ([]*ClassInfo, error) {
	var classes []*ClassInfo
	o := orm.NewOrm()
	_, err := o.QueryTable("class_info").All(&classes)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return classes,nil
}

func GetClassCourse(class *ClassInfo) error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(class, "Courses")
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// SetClassCourse 给整个班级进行统一选课
func SetClassCourse(class *ClassInfo, course *CourseInfo) error {
	o := orm.NewOrm()
	// 查询是否有该课程存在
	err := o.Read(course)
	if err != nil {
		logs.Error(err)
		return err
	}
	// 由班级获取学生列表
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}
	// 构建查询对象
	qb.Select(GetStudentColumn()).
		From("student_info").
		LeftJoin("class_info").On("student_info.class_id=class_info.class_id").
		Where("class_info.class_id = ?")
	// 导出 SQL 语句
	sql := qb.String()
	var maps []orm.Params
	_, err = o.Raw(sql, class.ClassId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return err
	}
	students := ParseStudents(maps)
	if students == nil || len(students) == 0 {
		return fmt.Errorf("该班级不存在或者该班级无学生")
	}
	courseClassRel := &CourseClassRel{Course: course, Class: class}
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	created, _, err := o.ReadOrCreate(courseClassRel, "Course", "Class")
	if err != nil {
		logs.Error(err)
		return err
	}
	if !created {
		return fmt.Errorf("该班级已经选过此课")
	}
	// 批量给学生选课
	qs := o.QueryTable("course_student_rel")
	i, _ := qs.PrepareInsert()
	for _, student := range students {
		_, err = i.Insert(&CourseStudentRel{Course: course, Student: student})
		if err != nil {
			logs.Error(err)
		}
	}
	err = i.Close() // 别忘记关闭 statement
	if err != nil {
		logs.Error(err)
	}
	return nil
}

func GetClassStudentSort(class *ClassInfo) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}

	// 构建查询对象
	qb.Select(GetStudentColumn()).
		From("student_info").
		LeftJoin("class_info").On("student_info.class_id=class_info.class_id").
		Where("class_info.class_id = ?").
		OrderBy("student_all_point").Desc()

	// 导出 SQL 语句
	sql := qb.String()

	o := orm.NewOrm()
	var maps []orm.Params
	_, err = o.Raw(sql, class.ClassId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return err
	}
	class.Students = ParseStudents(maps)
	return nil
}
