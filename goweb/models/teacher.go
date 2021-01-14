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
		GetCourseGroupTeacherRelColumn(false, false), GetCourseGroupColumn(),
		GetCourseTeacherRelColumn(false, false), GetCourseColumn()).
		From("teacher_info").
		LeftJoin("class_info").On("class_info.teacher_id=teacher_info.teacher_id").
		LeftJoin("course_group_teacher_rel").On("course_group_teacher_rel.teacher_id=teacher_info.teacher_id").
		LeftJoin("course_group_info").On("course_group_info.course_group_id=course_group_teacher_rel.course_group_id").
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
	if err != nil && err.Error() != "<Ormer> last insert id is unavailable" {
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

func DeleteTeacher(teacher *TeacherInfo) error {
	_, err := orm.NewOrm().Delete(teacher)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func GetTeachers() ([]*TeacherInfo, error) {
	var teachers []*TeacherInfo
	_, err := orm.NewOrm().QueryTable("teacher_info").All(&teachers)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return teachers,nil
}

// GetTeacherCourse
func GetTeacherCourse(teacher *TeacherInfo) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}

	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseTeacherRelColumn(false, false),
		GetTeacherColumn(), GetCourseClassRelColumn(false, false),
		GetClassColumn(), GetCourseGroupRelColumn(false, false),
		GetCourseGroupColumn(), GetCourseStudentRelColumn(false, true)).
		From("course_info").
		LeftJoin("course_base_info").On("course_base_info.course_id=course_info.course_id").
		InnerJoin("course_teacher_rel").On("course_info.course_id=course_teacher_rel.course_id").
		LeftJoin("teacher_info").On("teacher_info.teacher_id=course_teacher_rel.teacher_id").
		LeftJoin("course_class_rel").On("course_class_rel.course_id=course_info.course_id").
		LeftJoin("class_info").On("class_info.class_id=course_class_rel.class_id").
		LeftJoin("course_group_rel").On("course_group_rel.course_id=course_info.course_id").
		LeftJoin("course_group_info").On("course_group_info.course_group_id=course_group_rel.course_group_id").
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

	qb.Select(GetCourseStudentRelColumn(true, false), GetStudentColumn(), GetClassColumn()).
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


func TeacherChooseCourse(rel *CourseTeacherRel) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}
	// 构建查询对象
	qb.Select(GetCourseColumn(), GetCourseBaseColumn(), GetCourseTeacherRelColumn(false, true)).
		From("course_info").
		LeftJoin("course_base_info").On("course_base_info.course_id=course_info.course_id").
		LeftJoin("course_teacher_rel").On("course_info.course_id = course_teacher_rel.course_id").
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
	// 判断选课人数是否超过最大限制

	// 判断课程是否已经选
	if course.Teachers != nil {
		for _, teacher:= range course.Teachers {
			if teacher.TeacherId == rel.Teacher.TeacherId {
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
	rel.CourseTeacherRelId = int(id)
	rel.Course = course
	return nil
}

func GetTeacherList(pageInfo *utils.PageInfo, teacher *StudentInfo) error {
	qb, err := orm.NewQueryBuilder("postgres")
	if err != nil {
		logs.Error(err)
		return err
	}

	// 构建查询对象
	qb.Select(GetTeacherColumn()).
		From("teacher_info")

	// 导出 SQL 语句
	sql := qb.String()

	o := orm.NewOrm()
	var teachers []*TeacherInfo
	_, err = o.Raw(sql).QueryRows(&teachers)
	if err != nil {
		logs.Error(err)
		return err
	}
	pageInfo.Lists = teachers
	return nil
}


