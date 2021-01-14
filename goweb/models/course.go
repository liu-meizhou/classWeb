package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"goweb/utils"
)

func ReadCourse(courseId string) (*CourseInfo, error) {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// 构建查询对象
	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseTeacherRelColumn(false, false),
		GetTeacherColumn(), GetCourseClassRelColumn(false, false),
		GetClassColumn(), GetCourseGroupRelColumn(false, false), GetCourseGroupColumn(),
		GetCourseStudentRelColumn(false, false), GetStudentColumn()).
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
	to, err := o.Begin()
	if err != nil {
		logs.Error("start the transaction failed")
		return fmt.Errorf("start the transaction failed")
	}
	_, err = o.Insert(course)
	if err != nil && err.Error() != "<Ormer> last insert id is unavailable" {
		logs.Error(err)
		errRollBack := to.Rollback()
		if errRollBack != nil {
			logs.Error("roll back transaction failed", errRollBack)
			return errRollBack
		}
		return err
	}
	if course.CourseBases != nil && len(course.CourseBases) != 0 {
		for _, base := range course.CourseBases {
			base.Course = course
			_, err = o.Insert(base)
			if err != nil {
				logs.Error(err)
				errRollBack := to.Rollback()
				if errRollBack != nil {
					logs.Error("roll back transaction failed", errRollBack)
					return errRollBack
				}
				return err
			}
		}
	}
	err = to.Commit()
	if err != nil {
		logs.Error("commit transaction failed.", err)
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

func DeleteCourse(course *CourseInfo) error {
	_, err := orm.NewOrm().Delete(course)
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

func GetChooseCourse(pageInfo *utils.PageInfo, course *CourseInfo) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}
	// 构建查询对象
	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseTeacherRelColumn(false, false),
		GetTeacherColumn(), GetCourseStudentRelColumn(false, true)).
		From("course_info").
		LeftJoin("course_base_info").On("course_base_info.course_id=course_info.course_id").
		LeftJoin("course_teacher_rel").On("course_teacher_rel.course_id=course_info.course_id").
		LeftJoin("teacher_info").On("teacher_info.teacher_id=course_teacher_rel.teacher_id").
		//LeftJoin("course_class_rel").On("course_class_rel.course_id=course_info.course_id").
		//LeftJoin("class_info").On("class_info.class_id=course_class_rel.class_id").
		//LeftJoin("course_group_rel").On("course_group_rel.course_id=course_info.course_id").
		//LeftJoin("course_group_info").On("course_group_info.course_group_id=course_group_rel.course_group_id").
		LeftJoin("course_student_rel").On("course_info.course_id = course_student_rel.course_id").
		Where("1 = 1")
	var args []interface{}
	// 构建查询参数
	if course.CourseName != "" {
		qb.And("course_info.course_name like ?")
		args = append(args, "%" + course.CourseName + "%")
	}
	if course.CourseScores != 0 {
		qb.And("course_info.course_scores = ?")
		args = append(args, course.CourseScores)
	}
	if course.CourseProperties != "" {
		qb.And("course_info.course_properties = ?")
		args = append(args, course.CourseProperties)
	}
	if course.CourseWay != "" {
		qb.And("course_info.course_way = ?")
		args = append(args, course.CourseWay)
	}
	if course.CourseCount != 0 {
		qb.And("course_info.course_count = ?")
		args = append(args, course.CourseCount)
	}

	// 导出 SQL 语句
	sql := qb.String()
	//logs.Debug(sql)

	o := orm.NewOrm()
	var maps []orm.Params
	_, err = o.Raw(sql, args).Values(&maps)
	if err != nil {
		logs.Error(err)
		return err
	}
	courses := ParseCourses(maps)
	pageInfo.Lists = courses
	return nil
}

func ChooseCourse(rel *CourseStudentRel) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}
	// 构建查询对象
	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseStudentRelColumn(false, true)).
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
	//// 判断课程是否可选
	//if course.CourseProperties != canChooseCourse {
	//	return fmt.Errorf("你选择的课程不可选择, 不是%v", canChooseCourse)
	//}
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
