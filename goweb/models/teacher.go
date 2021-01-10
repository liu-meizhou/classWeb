package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"goweb/utils"
)

func ReadTeacher(teacherId string) (*TeacherInfo, error) {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// 构建查询对象
	qb.Select(GetTeacherColumn(), GetClassColumn(),
		GetCourseGroupRelColumn(), GetCourseGroupColumn(),
		GetCourseTeacherRelColumn(), GetCourseColumn()).
		From("teacher_info").
		LeftJoin("class_info").On("class_info.teacher_id=teacher_info.teacher_id").
		LeftJoin("course_group_rel").On("course_group_rel.teacher_id=teacher_info.teacher_id").
		LeftJoin("course_group_info").On("class_group_info.class_group_id=course_group_rel.class_group_id").
		LeftJoin("course_teacher_rel").On("course_teacher_rel.teacher_id=teacher_info.teacher_id").
		LeftJoin("course_info").On("course_info.course_id=course_student_rel.course_id").
		Where("teacher_info.teacher_id = ?")

	// 导出 SQL 语句
	sql := qb.String()

	o := orm.NewOrm()
	var maps []orm.Params
	_, err = o.Raw(sql, teacherId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	teacher := ParseTeacher(maps)

	return teacher, nil
}

func CreateTeacher(teacher *TeacherInfo) error {
	o := orm.NewOrm()
	_, err := o.Insert(teacher)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func UpdateTeacher(teacher *TeacherInfo) error {
	o := orm.NewOrm()
	_, err := o.Update(teacher)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func DeleteTeacher(student *TeacherInfo)  {
}

// GetTeacherCourse
func GetTeacherCourse(teacher *TeacherInfo) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}

	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseTeacherRelColumn(), GetTeacherColumn(), GetCourseClassRelColumn(),
		GetClassColumn(), GetCourseGroupRelColumn(), GetCourseGroupColumn(), GetCourseStudentRelColumn()).
		From("course_info").
		LeftJoin("course_base_info").On("course_base_info.course_id=course_info.course_id").
		InnerJoin("course_teacher_rel").On("course_info.course_id=course_teacher_rel.course_id").
		LeftJoin("teacher_info").On("teacher_info.teacher_id=course_teacher_rel.teacher_id").
		LeftJoin("course_class_rel").On("course_class_rel.course_id=course_info.course_id").
		LeftJoin("class_info").On("class_info.class_id=course_class_rel.class_id").
		LeftJoin("course_group_rel").On("course_group_rel.course_id=course_info.course_id").
		LeftJoin("course_group_info").On("class_group_info.class_group_id=course_group_rel.class_group_id").
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
func GetGradeCourse(teacher *TeacherInfo, pageInfo *utils.PageInfo, course *CourseInfo) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
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
	_, err = o.Raw(sql, course.CourseId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return err
	}
	students := ParseCourseStudent(maps)
	pageInfo.Lists = students
	return nil
}

// IsTeacherCourse 判断输入是否在库
func IsTeacherCourse(rel *CourseTeacherRel) error {
	o := orm.NewOrm()
	err := o.Read(rel, "Course", "Teacher")
	if err == orm.ErrNoRows {
		return fmt.Errorf("这课不是你的,或者课程号输入错误")
	} else if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// SetStudentGradeRel
func SetStudentGradeRel(courseStudentRel *utils.CourseStudentRel) error {
	// 计算根据成绩绩点
	courseStudentRel.StudentPoint = (courseStudentRel.StudentResults - 50)/10
	if courseStudentRel.StudentPoint < 0 {
		courseStudentRel.StudentPoint = 0
	} else if courseStudentRel.StudentPoint > 4 {
		courseStudentRel.StudentPoint = 4
	}
	o := orm.NewOrm()
	_, err := o.QueryTable("course_student_rel").
		Filter("student_id", courseStudentRel.StudentId).Filter("course_id", courseStudentRel.CourseId).
		Update(orm.Params{
			"student_results": courseStudentRel.StudentResults,
			"student_point":   courseStudentRel.StudentPoint,
		})
	if err!=nil {
		logs.Error(err)
		return err
	}
	return nil
}

