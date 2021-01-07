package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

// 连表查询，小表驱动大表(数量级小的表扫描连接数量级大的表)

//// RoleInfo 角色信息表
//type RoleInfo struct {
//	RoleId	      int           `orm:"pk"`                                                  // 主键号,id,主键
//	RoleName    string        `orm:"size(50)"`                                            // 角色名字 Admin,学生,老师,系主任
//	CreatedTime    time.Time     `orm:"auto_now_add;type(datetime)"`
//	UpdatedTime    time.Time     `orm:"auto_now;type(datetime)"`
//}

// StudentInfo 学生信息表 连表优先 班级->课程
type StudentInfo struct {
	StudentId          string        `orm:"pk"`                                                  // 学生学号,id,主键
	Class              *ClassInfo    `orm:"null;rel(fk);on_delete(set_null)"`                    // 所在班级
	Courses            []*CourseInfo `orm:"rel(m2m);rel_through(goweb/models.CourseStudentRel)"` // 学生拥有的课程
	StudentName        string        `orm:"size(50)"`                                            // 学生名字
	StudentSex         string        `orm:"size(10)"`                                            // 学生性别
	StudentCollege     string        `orm:"size(50)"`                                            // 学生所在学院
	StudentResults     float64       `orm:"-"`                                                   // 课程成绩
	StudentPoint       float64       `orm:"-"`                                                   // 课程绩点
	StudentBirth       time.Time     `orm:"auto_now;type(datetime)"`                             // 学生出生日期
	StudentTime        time.Time     `orm:"auto_now_add;type(datetime)"`                         // 学生入学日期
	StudentCreatedTime time.Time     `orm:"auto_now_add;type(datetime)"`
	StudentUpdatedTime time.Time     `orm:"auto_now;type(datetime)"`
}

// ClassInfo 班级表 连表优先 老师->课程->学生
type ClassInfo struct {
	ClassId          string         `orm:"pk"`                                                // 班级号,id,主键
	ClassName        string         `orm:"size(20)"`                                          // 班级名
	Students         []*StudentInfo `orm:"reverse(many)"`                                     // 班级内的学生
	Courses          []*CourseInfo  `orm:"rel(m2m);rel_through(goweb/models.CourseClassRel)"` // 班级共同拥有的课程
	Teacher          *TeacherInfo   `orm:"null;rel(fk);on_delete(set_null)"`                  // 班主任
	ClassCreatedTime time.Time      `orm:"auto_now_add;type(datetime)"`
	ClassUpdatedTime time.Time      `orm:"auto_now;type(datetime)"`
}

// TeacherInfo 老师信息表 连表优先 班级->课程->课组 (具体不知道课组的需求量)
type TeacherInfo struct {
	TeacherId          string            `orm:"pk"`                                                      // 教工号,id,主键
	TeacherName        string            `orm:"size(50)"`                                                // 老师名字
	Classes            []*ClassInfo      `orm:"reverse(many)"`                                           // 教学班主
	Courses            []*CourseInfo     `orm:"rel(m2m);rel_through(goweb/models.CourseTeacherRel)"`     // 该老师教的课程
	ClassGroups        []*ClassGroupInfo `orm:"rel(m2m);rel_through(goweb/models.ClassGroupTeacherRel)"` // 该老师的课组
	TeacherSex         string            `orm:"size(10)"`                                                // 老师性别
	TeacherCollege     string            `orm:"size(50)"`                                                // 老师所在学院
	TeacherBirth       time.Time         `orm:"auto_now;type(datetime)"`                                 // 老师出生日期
	TeacherTime        time.Time         `orm:"auto_now_add;type(datetime)"`                             // 老师加入学校日期
	IsCharge           bool              `orm:"-"`
	TeacherCreatedTime time.Time         `orm:"auto_now_add;type(datetime)"`
	TeacherUpdatedTime time.Time         `orm:"auto_now;type(datetime)"`
}

// CourseInfo 课程总表 连表优先级 上课基本信息 -> 老师 -> 班级 -> 课组 -> 学生
type CourseInfo struct {
	CourseId          string            `orm:"pk"`                    // 课程号,id,主键
	Students          []*StudentInfo    `orm:"reverse(many)"`         // 选该课程的学生
	Classes           []*ClassInfo      `orm:"reverse(many)"`         // 选该课程的班级
	Teachers          []*TeacherInfo    `orm:"reverse(many)"`         // 教该课的老师,1门课程可由多个老师来教
	ClassGroups       []*ClassGroupInfo `orm:"reverse(many)"`         // 该课在哪些课组中
	CourseName        string            `orm:"size(50)"`              // 课程名
	CourseProperties  string            `orm:"size(50)"`              // 课程性质 专业必修，专业选修
	CourseBases       []*CourseBaseInfo `orm:"reverse(many)"`         // 上课时间地点
	CourseScores      float64           `orm:"digits(4);decimals(2)"` // 课程学分
	CourseWay         string            `orm:"size(10)"`              // 考核方式
	CourseCount       float64           `orm:"digits(5);decimals(2)"` // 学时, 单位小时
	StudentResults    float64           `orm:"-"`                     // 课程成绩
	StudentPoint      float64           `orm:"-"`                     // 课程绩点
	//StudentNumber      int           	`orm:"-"`                     // 选该课程的人数
	CourseCreatedTime time.Time         `orm:"auto_now_add;type(datetime)"`
	CourseUpdatedTime time.Time         `orm:"auto_now;type(datetime)"`
}

// CourseBaseInfo 上课时间地点
type CourseBaseInfo struct {
	CourseBaseId          int         `orm:"pk;auto"`                          // 主键号,id,主键,自增
	Course                *CourseInfo `orm:"null;rel(fk);on_delete(set_null)"` // 所属课程
	CourseStartYear       int         // 开始学期
	CourseEndYear         int         // 结束学期
	CourseYear            int         // 课程学年
	CourseStartWeek       int
	CourseEndWeek         int
	CourseWeek            int
	CourseStartCount      int
	CourseEndCount        int
	CourseSchool          string    `orm:"size(50)"` // 上课校区
	CourseAddress         string    `orm:"size(50)"` // 上课地点
	CourseAddressFloor    int       // 上课楼层号
	CourseAddressNumber   int       // 上课教室号
	CourseBaseCreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	CourseBaseUpdatedTime time.Time `orm:"auto_now;type(datetime)"`
}

// ClassGroupInfo 课组管理表 连表优先级 课程 -> 老师
type ClassGroupInfo struct {
	ClassGroupId          int            `orm:"pk;auto"`  // 课组id,主键,自增
	ClassGroupName        string         `orm:"size(50)"` // 课组名
	IsCharge              bool           `orm:"-"`
	Courses               []*CourseInfo  `orm:"rel(m2m);rel_through(goweb/models.CourseGroupRel)"` // 该课组中的课程
	Teachers              []*TeacherInfo `orm:"reverse(many)"`                                     // 该课组的老师
	ClassGroupCreatedTime time.Time      `orm:"auto_now_add;type(datetime)"`
	ClassGroupUpdatedTime time.Time      `orm:"auto_now;type(datetime)"`
}

// CourseStudentRel 学生课程联系表
type CourseStudentRel struct {
	CourseStudentRelId          int          `orm:"pk;auto"` // 学生课程联系表id,主键,自增
	Student                     *StudentInfo `orm:"null;rel(fk);on_delete(set_null)"`
	Course                      *CourseInfo  `orm:"null;rel(fk);on_delete(set_null)"`
	StudentResults              float64      `orm:"default(-1);digits(5);decimals(2)"` // 课程成绩
	StudentPoint                float64      `orm:"default(-1);digits(4);decimals(2)"` // 课程绩点
	CourseStudentRelCreatedTime time.Time    `orm:"auto_now_add;type(datetime)"`
	CourseStudentRelUpdatedTime time.Time    `orm:"auto_now;type(datetime)"`
}

// ClassGroupTeacherRel 课组老师联系表
type ClassGroupTeacherRel struct {
	ClassGroupTeacherRelId          int             `orm:"pk;auto"` // 课组老师表id,主键,自增
	ClassGroup                      *ClassGroupInfo `orm:"null;rel(fk);on_delete(set_null)"`
	Teacher                         *TeacherInfo    `orm:"null;rel(fk);on_delete(set_null)"`
	IsCharge                        bool            // 是否负责该课组
	ClassGroupTeacherRelCreatedTime time.Time       `orm:"auto_now_add;type(datetime)"`
	ClassGroupTeacherRelUpdatedTime time.Time       `orm:"auto_now;type(datetime)"`
}

// CourseClassRel 课程班级联系表
type CourseClassRel struct {
	CourseClassRelId          int         `orm:"pk;auto"` // 课程班级表id,主键,自增
	Course                    *CourseInfo `orm:"null;rel(fk);on_delete(set_null)"`
	Class                     *ClassInfo  `orm:"null;rel(fk);on_delete(set_null)"`
	CourseClassRelCreatedTime time.Time   `orm:"auto_now_add;type(datetime)"`
	CourseClassRelUpdatedTime time.Time   `orm:"auto_now;type(datetime)"`
}

// CourseTeacherRel 课程老师联系表
type CourseTeacherRel struct {
	CourseTeacherRelId          int          `orm:"pk;auto"` // 课程老师表id,主键,自增
	Course                      *CourseInfo  `orm:"null;rel(fk);on_delete(set_null)"`
	Teacher                     *TeacherInfo `orm:"null;rel(fk);on_delete(set_null)"`
	CourseTeacherRelCreatedTime time.Time    `orm:"auto_now_add;type(datetime)"`
	CourseTeacherRelUpdatedTime time.Time    `orm:"auto_now;type(datetime)"`
}

// CourseGroupRel 课程与课程组联系表
type CourseGroupRel struct {
	CourseGroupRelId          int             `orm:"pk;auto"` // 课程与课程组表id,主键,自增
	ClassGroup                *ClassGroupInfo `orm:"null;rel(fk);on_delete(set_null)"`
	Course                    *CourseInfo     `orm:"null;rel(fk);on_delete(set_null)"`
	CourseGroupRelCreatedTime time.Time       `orm:"auto_now_add;type(datetime)"`
	CourseGroupRelUpdatedTime time.Time       `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(StudentInfo), new(ClassInfo), new(TeacherInfo),
		new(CourseInfo), new(ClassGroupInfo), new(CourseStudentRel), new(ClassGroupTeacherRel),
		new(CourseClassRel), new(CourseTeacherRel), new(CourseGroupRel), new(CourseBaseInfo))
}

func GetStudentColumn() string {
	return `student_info.student_id",
"student_info.student_name",
"student_info.student_sex", 
"student_info.student_college",
"student_info.student_birth",
"student_info.student_time",
"student_info.student_created_time",
"student_info.student_updated_time`
}

func GetClassColumn() string {
	return `class_info.class_id",
"class_info.class_name",
"class_info.class_created_time",
"class_info.class_updated_time`
}

func GetTeacherColumn() string {
	return `teacher_info.teacher_id",
"teacher_info.teacher_name",
"teacher_info.teacher_sex", 
"teacher_info.teacher_college",
"teacher_info.teacher_birth",
"teacher_info.teacher_time",
"teacher_info.teacher_created_time",
"teacher_info.teacher_updated_time`
}

func GetCourseColumn() string {
	return `course_info.course_id",
"course_info.course_name",
"course_info.course_properties",
"course_info.course_scores",
"course_info.course_way",
"course_info.course_count", 
"course_info.course_created_time",
"course_info.course_updated_time`
}

func GetCourseBaseColumn() string {
	return `course_base_info.course_base_id",
"course_base_info.course_start_year",
"course_base_info.course_end_year",
"course_base_info.course_year",
"course_base_info.course_start_week",
"course_base_info.course_end_week",
"course_base_info.course_week",
"course_base_info.course_start_count",
"course_base_info.course_end_count",
"course_base_info.course_school",
"course_base_info.course_address",
"course_base_info.course_address_floor",
"course_base_info.course_address_number",
"course_base_info.course_base_created_time",
"course_base_info.course_base_updated_time`
}

func GetClassGroupColumn() string {
	return `class_group_info.class_group_id",
"class_group_info.class_group_name",
"class_group_info.class_group_created_time",
"class_group_info.class_group_updated_time`
}

func GetCourseStudentRelColumn() string {
	return `course_student_rel.course_student_rel_id",
"course_student_rel.student_id",
"course_student_rel.student_results",
"course_student_rel.student_point`
}

func GetClassGroupTeacherRelColumn() string {
	return `class_group_teacher_rel.is_charge`
}

func GetCourseClassRelColumn() string {
	return ""
}

func GetCourseTeacherRelColumn() string {
	return ""
}

func GetCourseGroupRelColumn() string {
	return ""
}

func ParseStudentInfo(param orm.Params) *StudentInfo {
	studentId := param["student_id"]
	if studentId == nil {
		return nil
	}
	student := &StudentInfo{StudentId: studentId.(string)}
	studentName := param["student_name"]
	if studentName != nil {
		student.StudentName = studentName.(string)
	}
	studentSex := param["student_sex"]
	if studentSex != nil {
		student.StudentSex = studentSex.(string)
	}
	studentCollege := param["student_college"]
	if studentCollege != nil {
		student.StudentCollege = studentCollege.(string)
	}
	studentBirth := param["student_birth"]
	if studentBirth != nil {
		student.StudentBirth, _ = time.Parse(time.RFC3339Nano, studentBirth.(string))
	}
	studentTime := param["student_time"]
	if studentTime != nil {
		student.StudentTime, _ = time.Parse(time.RFC3339Nano, studentTime.(string))
	}
	studentResults := param["student_results"]
	if studentResults != nil {
		student.StudentResults, _ = strconv.ParseFloat(studentResults.(string), 64)
	}
	studentPoint := param["student_point"]
	if studentPoint != nil {
		student.StudentPoint, _ = strconv.ParseFloat(studentPoint.(string), 64)
	}
	studentCreatedTime := param["student_created_time"]
	if studentCreatedTime != nil {
		student.StudentCreatedTime, _ = time.Parse(time.RFC3339Nano, studentCreatedTime.(string))
	}
	studentUpdatedTime := param["student_updated_time"]
	if studentUpdatedTime != nil {
		student.StudentUpdatedTime, _ = time.Parse(time.RFC3339Nano, studentUpdatedTime.(string))
	}
	return student
}

func ParseClassInfo(param orm.Params) *ClassInfo {
	classId := param["class_id"]
	if classId == nil {
		return nil
	}
	class := &ClassInfo{ClassId: classId.(string)}
	className := param["class_name"]
	if className != nil {
		class.ClassName = className.(string)
	}
	classCreatedTime := param["class_created_time"]
	if classCreatedTime != nil {
		class.ClassCreatedTime, _ = time.Parse(time.RFC3339Nano, classCreatedTime.(string))
	}
	classUpdatedTime := param["class_updated_time"]
	if classUpdatedTime != nil {
		class.ClassUpdatedTime, _ = time.Parse(time.RFC3339Nano, classUpdatedTime.(string))
	}
	return class
}

func ParseTeacherInfo(param orm.Params) *TeacherInfo {
	teacherId := param["teacher_id"]
	if teacherId == nil {
		return nil
	}
	teacher := &TeacherInfo{TeacherId: teacherId.(string)}
	teacherName := param["teacher_name"]
	if teacherName != nil {
		teacher.TeacherName = teacherName.(string)
	}
	teacherSex := param["teacher_sex"]
	if teacherSex != nil {
		teacher.TeacherSex = teacherSex.(string)
	}
	teacherCollege := param["teacher_college"]
	if teacherCollege != nil {
		teacher.TeacherCollege = teacherCollege.(string)
	}
	teacherBirth := param["teacher_birth"]
	if teacherBirth != nil {
		teacher.TeacherBirth, _ = time.Parse(time.RFC3339Nano, teacherBirth.(string))
	}
	teacherTime := param["teacher_time"]
	if teacherTime != nil {
		teacher.TeacherTime, _ = time.Parse(time.RFC3339Nano, teacherTime.(string))
	}
	isCharge := param["is_charge"]
	if isCharge != nil {
		teacher.IsCharge = isCharge.(string) == "true"
	}
	teacherCreatedTime := param["teacher_created_time"]
	if teacherCreatedTime != nil {
		teacher.TeacherCreatedTime, _ = time.Parse(time.RFC3339Nano, teacherCreatedTime.(string))
	}
	teacherUpdatedTime := param["teacher_updated_time"]
	if teacherUpdatedTime != nil {
		teacher.TeacherUpdatedTime, _ = time.Parse(time.RFC3339Nano, teacherUpdatedTime.(string))
	}
	return teacher
}

func ParseCourseInfo(param orm.Params) *CourseInfo {
	courseId := param["course_id"]
	if courseId == nil {
		return nil
	}
	course := &CourseInfo{CourseId: courseId.(string)}
	courseName := param["course_name"]
	if courseName != nil {
		course.CourseName = courseName.(string)
	}
	courseProperties := param["course_properties"]
	if courseProperties != nil {
		course.CourseProperties = courseProperties.(string)
	}
	courseScores := param["course_scores"]
	if courseScores != nil {
		course.CourseScores, _ = strconv.ParseFloat(courseScores.(string), 64)
	}
	courseWay := param["course_way"]
	if courseWay != nil {
		course.CourseWay = courseWay.(string)
	}
	courseCount := param["course_count"]
	if courseCount != nil {
		course.CourseCount, _ = strconv.ParseFloat(courseCount.(string), 64)
	}
	studentResults := param["student_results"]
	if studentResults != nil {
		course.StudentResults, _ = strconv.ParseFloat(studentResults.(string), 64)
	}
	studentPoint := param["student_point"]
	if studentPoint != nil {
		course.StudentPoint, _ = strconv.ParseFloat(studentPoint.(string), 64)
	}
	courseCreatedTime := param["course_created_time"]
	if courseCreatedTime != nil {
		course.CourseCreatedTime, _ = time.Parse(time.RFC3339Nano, courseCreatedTime.(string))
	}
	courseUpdatedTime := param["course_updated_time"]
	if courseUpdatedTime != nil {
		course.CourseUpdatedTime, _ = time.Parse(time.RFC3339Nano, courseUpdatedTime.(string))
	}
	return course
}

func ParseCourseBaseInfo(param orm.Params) *CourseBaseInfo {
	courseBaseIdStr := param["course_base_id"]
	if courseBaseIdStr == nil {
		return nil
	}
	courseBaseId, _ := strconv.Atoi(courseBaseIdStr.(string))
	courseBase := &CourseBaseInfo{CourseBaseId: courseBaseId}
	courseStartYear := param["course_start_year"]
	if courseStartYear != nil {
		courseBase.CourseStartYear, _ = strconv.Atoi(courseStartYear.(string))
	}
	courseEndYear := param["course_end_year"]
	if courseEndYear != nil {
		courseBase.CourseEndYear, _ = strconv.Atoi(courseEndYear.(string))
	}
	courseYear := param["course_year"]
	if courseYear != nil {
		courseBase.CourseYear, _ = strconv.Atoi(courseYear.(string))
	}
	courseStartWeek := param["course_start_week"]
	if courseStartWeek != nil {
		courseBase.CourseStartWeek, _ = strconv.Atoi(courseStartWeek.(string))
	}
	courseEndWeek := param["course_end_week"]
	if courseEndWeek != nil {
		courseBase.CourseEndWeek, _ = strconv.Atoi(courseEndWeek.(string))
	}
	courseWeek := param["course_week"]
	if courseWeek != nil {
		courseBase.CourseWeek, _ = strconv.Atoi(courseWeek.(string))
	}
	courseStartCount := param["course_start_count"]
	if courseStartCount != nil {
		courseBase.CourseStartCount, _ = strconv.Atoi(courseStartCount.(string))
	}
	courseEndCount := param["course_end_count"]
	if courseEndCount != nil {
		courseBase.CourseEndCount, _ = strconv.Atoi(courseEndCount.(string))
	}
	courseSchool := param["course_school"]
	if courseSchool != nil {
		courseBase.CourseSchool = courseSchool.(string)
	}
	courseAddress := param["course_address"]
	if courseAddress != nil {
		courseBase.CourseAddress = courseAddress.(string)
	}
	courseAddressFloor := param["course_address_floor"]
	if courseAddressFloor != nil {
		courseBase.CourseAddressFloor, _ = strconv.Atoi(courseAddressFloor.(string))
	}
	courseAddressNumber := param["course_address_number"]
	if courseAddressNumber != nil {
		courseBase.CourseAddressNumber, _ = strconv.Atoi(courseAddressNumber.(string))
	}
	courseBaseCreatedTime := param["course_base_created_time"]
	if courseBaseCreatedTime != nil {
		courseBase.CourseBaseCreatedTime, _ = time.Parse(time.RFC3339Nano, courseBaseCreatedTime.(string))
	}
	courseBaseUpdatedTime := param["course_base_updated_time"]
	if courseBaseUpdatedTime != nil {
		courseBase.CourseBaseUpdatedTime, _ = time.Parse(time.RFC3339Nano, courseBaseUpdatedTime.(string))
	}
	return courseBase
}

func ParseClassGroupInfo(param orm.Params) *ClassGroupInfo {
	classGroupIdStr := param["class_group_id"]
	if classGroupIdStr == nil {
		return nil
	}
	classGroupId, _ := strconv.Atoi(classGroupIdStr.(string))
	classGroup := &ClassGroupInfo{ClassGroupId: classGroupId}
	classGroupName := param["class_group_name"]
	if classGroupName != nil {
		classGroup.ClassGroupName = classGroupName.(string)
	}
	isCharge := param["is_charge"]
	if isCharge != nil {
		classGroup.IsCharge = isCharge.(string) == "true"
	}
	classGroupCreatedTime := param["class_group_created_time"]
	if classGroupCreatedTime != nil {
		classGroup.ClassGroupCreatedTime, _ = time.Parse(time.RFC3339Nano, classGroupCreatedTime.(string))
	}
	classGroupUpdatedTime := param["class_group_updated_time"]
	if classGroupUpdatedTime != nil {
		classGroup.ClassGroupUpdatedTime, _ = time.Parse(time.RFC3339Nano, classGroupUpdatedTime.(string))
	}
	return classGroup
}

func ParseCourseStudentRel(param orm.Params) *CourseStudentRel {
	return nil
}

func ParseClassGroupTeacherRel(param orm.Params) *ClassGroupTeacherRel {
	return nil
}

func ParseCourseClassRel(param orm.Params) *CourseClassRel {
	return nil
}

func ParseCourseTeacherRel(param orm.Params) *CourseTeacherRel {
	return nil
}

func ParseCourseGroupRel(param orm.Params) *CourseGroupRel {
	return nil
}

func ParseCourses(params []orm.Params) []*CourseInfo {
	var courses []*CourseInfo
	courseMap := make(map[string]*CourseInfo)
	studentMap := make(map[string]map[string]*StudentInfo)
	classMap := make(map[string]map[string]*ClassInfo)
	teacherMap := make(map[string]map[string]*TeacherInfo)
	classGroupMap := make(map[string]map[int]*ClassGroupInfo)
	for _, param := range params {
		course := ParseCourseInfo(param)
		if course == nil {
			continue
		}
		student := ParseStudentInfo(param)
		class := ParseClassInfo(param)
		teacher := ParseTeacherInfo(param)
		classGroup := ParseClassGroupInfo(param)
		courseBase := ParseCourseBaseInfo(param)
		if oldCourse, ok := courseMap[course.CourseId]; ok {
			if student != nil {
				if _, ok := studentMap[course.CourseId][student.StudentId]; !ok {
					oldCourse.Students = append(oldCourse.Students, student)
					continue
				}
			}
			if class != nil {
				if _, ok := classMap[course.CourseId][class.ClassId]; !ok {
					oldCourse.Classes = append(oldCourse.Classes, class)
					continue
				}
			}
			if teacher != nil {
				if _, ok := teacherMap[course.CourseId][teacher.TeacherId]; !ok {
					oldCourse.Teachers = append(oldCourse.Teachers, teacher)
					continue
				}
			}
			if classGroup != nil {
				if _, ok := classGroupMap[course.CourseId][classGroup.ClassGroupId]; !ok {
					oldCourse.ClassGroups = append(oldCourse.ClassGroups, classGroup)
					continue
				}
			}
			oldCourse.CourseBases = append(oldCourse.CourseBases, courseBase)
			continue
		}
		if _, ok := studentMap[course.CourseId]; !ok {
			studentMap[course.CourseId] = make(map[string]*StudentInfo)
		}
		if _, ok := classMap[course.CourseId]; !ok {
			classMap[course.CourseId] = make(map[string]*ClassInfo)
		}
		if _, ok := teacherMap[course.CourseId]; !ok {
			teacherMap[course.CourseId] = make(map[string]*TeacherInfo)
		}
		if _, ok := classGroupMap[course.CourseId]; !ok {
			classGroupMap[course.CourseId] = make(map[int]*ClassGroupInfo)
		}
		courseMap[course.CourseId] = course

		if student != nil {
			studentMap[course.CourseId][student.StudentId] = student
			course.Students = append(course.Students, student)
		}
		if class != nil {
			classMap[course.CourseId][class.ClassId] = class
			course.Classes = append(course.Classes, class)
		}
		if teacher != nil {
			teacherMap[course.CourseId][teacher.TeacherId] = teacher
			course.Teachers = append(course.Teachers, teacher)
		}
		if classGroup != nil {
			classGroupMap[course.CourseId][classGroup.ClassGroupId] = classGroup
			course.ClassGroups = append(course.ClassGroups, classGroup)
		}
		if courseBase != nil {
			course.CourseBases = append(course.CourseBases, courseBase)
		}

		courses = append(courses, course)
	}
	return courses
}

func ParseCourseStudent(params []orm.Params) []*StudentInfo {
	var students []*StudentInfo
	for _, param := range params {
		student := ParseStudentInfo(param)
		if student == nil {
			continue
		}
		class := ParseClassInfo(param)
		student.Class = class
		students = append(students, student)
	}
	return students
}

