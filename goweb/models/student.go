package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
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
//	o.LoadRelated(student.Courses[course], "ClassGroups")
//}
//return nil

func GetStudentCourse(student *StudentInfo) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}

	// 构建查询对象
	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseTeacherRelColumn(), GetTeacherColumn(), GetCourseClassRelColumn(),
		GetClassColumn(), GetCourseGroupRelColumn(), GetClassGroupColumn(), GetCourseStudentRelColumn()).
		From("course_info").
		LeftJoin("course_base_info").On("course_base_info.course_id=course_info.course_id").
		LeftJoin("course_teacher_rel").On("course_teacher_rel.course_id=course_info.course_id").
		LeftJoin("teacher_info").On("teacher_info.teacher_id=course_teacher_rel.teacher_id").
		LeftJoin("course_class_rel").On("course_class_rel.course_id=course_info.course_id").
		LeftJoin("class_info").On("class_info.class_id=course_class_rel.class_id").
		LeftJoin("course_group_rel").On("course_group_rel.course_id=course_info.course_id").
		LeftJoin("class_group_info").On("class_group_info.class_group_id=course_group_rel.class_group_id").
		InnerJoin("course_student_rel").On("course_info.course_id = course_student_rel.course_id").
		Where("course_student_rel.student_id = ?")

	// 导出 SQL 语句
	sql := qb.String()
	logs.Debug(sql)

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

const canChooseCourse = "专业选修课程"

func GetChooseCourse(student *StudentInfo) ([]*CourseInfo, error) {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// 构建查询对象
	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseTeacherRelColumn(), GetTeacherColumn(),
		GetCourseStudentRelColumn()).
		From("course_info").
		LeftJoin("course_base_info").On("course_base_info.course_id=course_info.course_id").
		LeftJoin("course_teacher_rel").On("course_teacher_rel.course_id=course_info.course_id").
		LeftJoin("teacher_info").On("teacher_info.teacher_id=course_teacher_rel.teacher_id").
		//LeftJoin("course_class_rel").On("course_class_rel.course_id=course_info.course_id").
		//LeftJoin("class_info").On("class_info.class_id=course_class_rel.class_id").
		//LeftJoin("course_group_rel").On("course_group_rel.course_id=course_info.course_id").
		//LeftJoin("class_group_info").On("class_group_info.class_group_id=course_group_rel.class_group_id").
		LeftJoin("course_student_rel").On("course_info.course_id = course_student_rel.course_id").
		Where("course_info.course_properties = ?").
		Limit(10).Offset(0)

	// 导出 SQL 语句
	sql := qb.String()
	//logs.Debug(sql)

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(sql, canChooseCourse).Values(&maps)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	logs.Info(num)

	courses := ParseCourses(maps)

	return courses, nil
}

func ChooseCourse(rel *CourseStudentRel) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}
	// 构建查询对象
	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseStudentRelColumn()).
		From("course_info").
		LeftJoin("course_base_info").On("course_base_info.course_id=course_info.course_id").
		LeftJoin("course_student_rel").On("course_info.course_id = course_student_rel.course_id").
		//LeftJoin("student_info").On("course_student_rel.student_id = student_info.student_id").
		Where("course_info.course_id = ?")

	// 导出 SQL 语句
	sql := qb.String()
	//logs.Debug(sql)

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(sql, rel.Course.CourseId).Values(&maps)
	if err != nil {
		logs.Error(err)
		return err
	}
	// 课程是否存在
	if num == 0 {
		return fmt.Errorf("你选择的课程不存在")
	}
	course := ParseCourse(maps)
	if course == nil {
		return fmt.Errorf("你选择的课程不存在")
	}
	// 判断课程是否可选
	if course.CourseProperties != canChooseCourse {
		return fmt.Errorf("你选择的课程不可选择, 不是%v", canChooseCourse)
	}
	// 判断选课人数是否超过最大限制

	// 判断课程是否已经选
	if course.Students != nil {
		for _, student := range course.Students {
			if student.StudentId == rel.Student.StudentId {
				//if num > 1 {
				//	logs.Warning("一个学生出现多条相同的课程了, 学生号: %v, 课程号: %v", rel.Student.StudentId, rel.Course.CourseId)
				//}
				return fmt.Errorf("你已经选上了,不可重复选择")
			}
		}
	}
	id, err := o.Insert(rel)
	if err!=nil {
		logs.Error(err)
		return err
	}
	rel.CourseStudentRelId = int(id)
	rel.Course = course
	return nil
}
