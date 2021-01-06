package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

// {{课程名}}/{{上课周数}}-{{结课周数}}周/{{上课校区}} {{上课地点}}/({{上课学年}}-{{结课学年}}-{{上课学期}})-{{课程号}}-{{xxx}}/{{教学班级}}/{{课程类型}}

/*
数据结构实验/9-16周/大学城 电子楼517/(2020-2021-1)-180600008-8/软件193/专业必修课程/40/40/考查/无/实验:16/2/16
*/

/*
人工智能编程语言/(1-4节)1-16周/大学城 电子楼516人工智能/许鹏/(2020-2021-1)-200602001-1/计科182/专业选修课程/考查/无/4/64/2.0
*/

//// RoleInfo 角色信息表
//type RoleInfo struct {
//	RoleId	      int           `orm:"pk"`                                                  // 主键号,id,主键
//	RoleName    string        `orm:"size(50)"`                                            // 角色名字 Admin,学生,老师,系主任
//	CreatedTime    time.Time     `orm:"auto_now_add;type(datetime)"`
//	UpdatedTime    time.Time     `orm:"auto_now;type(datetime)"`
//}

// StudentInfo 学生信息表
type StudentInfo struct {
	StudentId      int           `orm:"pk"`                                                  // 学生学号,id,主键
	Class          *ClassInfo    `orm:"null;rel(fk);on_delete(set_null)"`                    // 所在班级
	Courses        []*CourseInfo `orm:"rel(m2m);rel_through(goweb/models.CourseStudentRel)"` // 学生拥有的课程
	StudentName    string        `orm:"size(50)"`                                            // 学生名字
	StudentSex     string        `orm:"size(10)"`                                            // 学生性别
	StudentCollege string        `orm:"size(50)"`                                            // 学生所在学院
	StudentBirth   time.Time     `orm:"auto_now;type(datetime)"`                             // 学生出生日期
	StudentTime    time.Time     `orm:"auto_now_add;type(datetime)"`                         // 学生入学日期
	CreatedTime    time.Time     `orm:"auto_now_add;type(datetime)"`
	UpdatedTime    time.Time     `orm:"auto_now;type(datetime)"`
}

// ClassInfo 班级表
type ClassInfo struct {
	ClassId     int            `orm:"pk"`                                                // 班级号,id,主键
	ClassName   string         `orm:"size(20)"`                                          // 班级名
	Students    []*StudentInfo `orm:"reverse(many)"`                                     // 班级内的学生
	Courses     []*CourseInfo  `orm:"rel(m2m);rel_through(goweb/models.CourseClassRel)"` // 班级共同拥有的课程
	Teacher     *TeacherInfo   `orm:"null;rel(fk);on_delete(set_null)"`                  // 班主任
	CreatedTime time.Time      `orm:"auto_now_add;type(datetime)"`
	UpdatedTime time.Time      `orm:"auto_now;type(datetime)"`
}

// TeacherInfo 学生信息表
type TeacherInfo struct {
	TeacherId      int               `orm:"pk"`                                                      // 教工号,id,主键
	TeacherName    string            `orm:"size(50)"`                                                // 老师名字
	Classes        []*ClassInfo      `orm:"reverse(many)"`                                           // 教学班主
	Courses        []*CourseInfo     `orm:"rel(m2m);rel_through(goweb/models.CourseTeacherRel)"`     // 该老师教的课程
	ClassGroups    []*ClassGroupInfo `orm:"rel(m2m);rel_through(goweb/models.ClassGroupTeacherRel)"` // 该老师的课组
	TeacherSex     string            `orm:"size(10)"`                                                // 老师性别
	TeacherCollege string            `orm:"size(50)"`                                                // 老师所在学院
	TeacherBirth   time.Time         `orm:"auto_now;type(datetime)"`                                 // 老师出生日期
	TeacherTime    time.Time         `orm:"auto_now_add;type(datetime)"`                             // 老师加入学校日期
	IsCharge       bool              `orm:"-"`
	CreatedTime    time.Time         `orm:"auto_now_add;type(datetime)"`
	UpdatedTime    time.Time         `orm:"auto_now;type(datetime)"`
}

// CourseInfo 课程总表
type CourseInfo struct {
	CourseId         int               `orm:"pk;auto"`               // 课程号,id,主键,自增
	Students         []*StudentInfo    `orm:"reverse(many)"`         // 选该课程的学生
	Classes          []*ClassInfo      `orm:"reverse(many)"`         // 选该课程的班级
	Teachers         []*TeacherInfo    `orm:"reverse(many)"`         // 教该课的老师,1门课程可由多个老师来教
	ClassGroups      []*ClassGroupInfo `orm:"reverse(many)"`         // 该课在哪些课组中
	CourseName       string            `orm:"size(50)"`              // 课程名
	CourseProperties string            `orm:"size(50)"`              // 课程性质 专业必修，专业选修
	CourseBases      []*CourseBaseInfo `orm:"reverse(many)"`         // 上课时间地点
	CourseScores     float64           `orm:"digits(4);decimals(2)"` // 课程学分
	CourseWay        string            `orm:"size(10)"`              // 考核方式
	CourseCount      float64           `orm:"digits(5);decimals(2)"` // 学时, 单位小时
	StudentResults   float64           `orm:"-"`                     // 课程成绩
	StudentPoint     float64           `orm:"-"`                     // 课程绩点
	CreatedTime      time.Time         `orm:"auto_now_add;type(datetime)"`
	UpdatedTime      time.Time         `orm:"auto_now;type(datetime)"`
}

// CourseBaseInfo 上课时间地点
type CourseBaseInfo struct {
	CourseBaseId        int         `orm:"pk;auto"`                          // 主键号,id,主键,自增
	Course              *CourseInfo `orm:"null;rel(fk);on_delete(set_null)"` // 所属课程
	CourseStartYear     int         // 开始学期
	CourseEndYear       int         // 结束学期
	CourseYear          int         // 课程学年
	CourseStartWeek     int
	CourseEndWeek       int
	CourseWeek          int
	CourseStartCount    int
	CourseEndCount      int
	CourseSchool        string    `orm:"size(50)"` // 上课校区
	CourseAddress       string    `orm:"size(50)"` // 上课地点
	CourseAddressFloor  int       // 上课楼层号
	CourseAddressNumber int       // 上课教室号
	CreatedTime         time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedTime         time.Time `orm:"auto_now;type(datetime)"`
}

// ClassGroupInfo 课组管理表
type ClassGroupInfo struct {
	ClassGroupId   int            `orm:"pk;auto"`                                           // 课组id,主键,自增
	ClassGroupName string         `orm:"size(50)"`                                          // 课组名
	Courses        []*CourseInfo  `orm:"rel(m2m);rel_through(goweb/models.CourseGroupRel)"` // 该课组中的课程
	Teachers       []*TeacherInfo `orm:"reverse(many)"`                                     // 该课组的老师
	CreatedTime    time.Time      `orm:"auto_now_add;type(datetime)"`
	UpdatedTime    time.Time      `orm:"auto_now;type(datetime)"`
}

// CourseStudentRel 学生课程联系表
type CourseStudentRel struct {
	CourseStudentRelId int          `orm:"pk;auto"` // 学生课程联系表id,主键,自增
	Student            *StudentInfo `orm:"null;rel(fk);on_delete(set_null)"`
	Course             *CourseInfo  `orm:"null;rel(fk);on_delete(set_null)"`
	StudentResults     float64      `orm:"default(-1);digits(5);decimals(2)"` // 课程成绩
	StudentPoint       float64      `orm:"default(-1);digits(4);decimals(2)"` // 课程绩点
	CreatedTime        time.Time    `orm:"auto_now_add;type(datetime)"`
	UpdatedTime        time.Time    `orm:"auto_now;type(datetime)"`
}

// ClassGroupTeacherRel 课组老师联系表
type ClassGroupTeacherRel struct {
	ClassGroupTeacherRelId int             `orm:"pk;auto"` // 课组老师表id,主键,自增
	ClassGroup             *ClassGroupInfo `orm:"null;rel(fk);on_delete(set_null)"`
	Teacher                *TeacherInfo    `orm:"null;rel(fk);on_delete(set_null)"`
	IsCharge               bool            // 是否负责该课组
	CreatedTime            time.Time       `orm:"auto_now_add;type(datetime)"`
	UpdatedTime            time.Time       `orm:"auto_now;type(datetime)"`
}

// CourseClassRel 课程班级联系表
type CourseClassRel struct {
	CourseClassRelId int         `orm:"pk;auto"` // 课程班级表id,主键,自增
	Course           *CourseInfo `orm:"null;rel(fk);on_delete(set_null)"`
	Class            *ClassInfo  `orm:"null;rel(fk);on_delete(set_null)"`
	CreatedTime      time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdatedTime      time.Time   `orm:"auto_now;type(datetime)"`
}

// CourseTeacherRel 课程老师联系表
type CourseTeacherRel struct {
	CourseTeacherRelId int          `orm:"pk;auto"` // 课程老师表id,主键,自增
	Course             *CourseInfo  `orm:"null;rel(fk);on_delete(set_null)"`
	Teacher            *TeacherInfo `orm:"null;rel(fk);on_delete(set_null)"`
	CreatedTime        time.Time    `orm:"auto_now_add;type(datetime)"`
	UpdatedTime        time.Time    `orm:"auto_now;type(datetime)"`
}

// CourseGroupRel 课程与课程组联系表
type CourseGroupRel struct {
	CourseGroupRelId int             `orm:"pk;auto"` // 课程与课程组表id,主键,自增
	ClassGroup       *ClassGroupInfo `orm:"null;rel(fk);on_delete(set_null)"`
	Course           *CourseInfo     `orm:"null;rel(fk);on_delete(set_null)"`
	CreatedTime      time.Time       `orm:"auto_now_add;type(datetime)"`
	UpdatedTime      time.Time       `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(StudentInfo), new(ClassInfo), new(TeacherInfo),
		new(CourseInfo), new(ClassGroupInfo), new(CourseStudentRel), new(ClassGroupTeacherRel),
		new(CourseClassRel), new(CourseTeacherRel), new(CourseGroupRel), new(CourseBaseInfo))
}
