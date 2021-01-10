package models

import (
	"github.com/beego/beego/v2/client/orm"
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
	StudentId          string        `json:"id" orm:"pk" form:"id"`                                                    // 学生学号,id,主键
	StudentPassword    string        `json:"-" form:"password"`                                                        // 登录密码
	StudentType        string        `json:"userType" orm:"size(5)" form:"userType"`                                   // 学生用户类型
	Class              *ClassInfo    `json:"class" orm:"null;rel(fk);on_delete(set_null)"`                             // 所在班级
	Courses            []*CourseInfo `json:"courses" orm:"rel(m2m);rel_through(goweb/models.CourseStudentRel)"`        // 学生拥有的课程
	StudentName        string        `json:"name" orm:"size(50)" form:"name"`                                          // 学生名字
	StudentSex         string        `json:"sex" orm:"size(10)" form:"sex"`                                            // 学生性别
	StudentCollege     string        `json:"college" orm:"size(50)" form:"college"`                                    // 学生所在学院
	StudentAllPoint    float64       `json:"allPoint" orm:"default(0);digits(4);decimals(2)"`                                  // 学生总绩点
	StudentResults     float64       `json:"grade" orm:"-"`                                                            // 课程成绩
	StudentPoint       float64       `json:"point" orm:"-"`                                                            // 课程绩点
	StudentBirth       time.Time     `json:"birth" orm:"auto_now;type(datetime)" form:"birth"`                         // 学生出生日期
	StudentTime        time.Time     `json:"enterSchoolTime" orm:"auto_now_add;type(datetime)" form:"enterSchoolTime"` // 学生入学日期
	StudentCreatedTime time.Time     `json:"createdTime" orm:"auto_now_add;type(datetime)"`
	StudentUpdatedTime time.Time     `json:"updatedTime" orm:"auto_now;type(datetime)"`
}

// ClassInfo 班级表 连表优先 老师->课程->学生
type ClassInfo struct {
	ClassId          string         `json:"id" orm:"pk" form:"id"`                                           // 班级号,id,主键
	ClassName        string         `json:"name" orm:"size(20)" form:"name"`                                 // 班级名
	Students         []*StudentInfo `json:"students" orm:"reverse(many)"`                                    // 班级内的学生
	Courses          []*CourseInfo  `json:"courses" orm:"rel(m2m);rel_through(goweb/models.CourseClassRel)"` // 班级共同拥有的课程
	Teacher          *TeacherInfo   `json:"teacher" orm:"null;rel(fk);on_delete(set_null)" form:"teacher"`   // 班主任
	ClassCreatedTime time.Time      `json:"createdTime" orm:"auto_now_add;type(datetime)"`
	ClassUpdatedTime time.Time      `json:"updatedTime" orm:"auto_now;type(datetime)"`
}

// TeacherInfo 老师信息表 连表优先 班级->课程->课组 (具体不知道课组的需求量)
type TeacherInfo struct {
	TeacherId          string             `json:"id" orm:"pk" form:"id"`                                                                 // 教工号,id,主键
	TeacherPassword    string             `json:"-" form:"password"`                                                                           // 登录密码
	TeacherType        string             `json:"userType" orm:"size(5)" form:"userType"`                                                      // 老师用户类型
	TeacherName        string             `json:"name" orm:"size(50)" form:"name"`                                                         // 老师名字
	Classes            []*ClassInfo       `json:"classes" orm:"reverse(many)"`                                                 // 教学班主
	Courses            []*CourseInfo      `json:"courses" orm:"rel(m2m);rel_through(goweb/models.CourseTeacherRel)"`           // 该老师教的课程
	CourseGroups       []*CourseGroupInfo `json:"courseGroups" orm:"rel(m2m);rel_through(goweb/models.CourseGroupTeacherRel)"` // 该老师的课组
	TeacherSex         string             `json:"sex" orm:"size(10)" form:"sex"`                                                          // 老师性别
	TeacherCollege     string             `json:"college" orm:"size(50)" form:"college"`                                                      // 老师所在学院
	TeacherBirth       time.Time          `json:"birth" orm:"auto_now;type(datetime)" form:"birth"`                                         // 老师出生日期
	TeacherTime        time.Time          `json:"enterSchoolTime" orm:"auto_now_add;type(datetime)" form:"enterSchoolTime"`                           // 老师加入学校日期
	IsCharge           bool               `json:"isChargeGroup" orm:"-"`
	TeacherCreatedTime time.Time          `json:"createdTime" orm:"auto_now_add;type(datetime)"`
	TeacherUpdatedTime time.Time          `json:"updatedTime" orm:"auto_now;type(datetime)"`
}

// CourseInfo 课程总表 连表优先级 上课基本信息 -> 老师 -> 班级 -> 课组 -> 学生
type CourseInfo struct {
	CourseId         string             `json:"id" orm:"pk" form:"id"`                          // 课程号,id,主键
	Students         []*StudentInfo     `json:"students" orm:"reverse(many)"`                   // 选该课程的学生
	Classes          []*ClassInfo       `json:"classes" orm:"reverse(many)"`                    // 选该课程的班级
	Teachers         []*TeacherInfo     `json:"teachers" orm:"reverse(many)"`                   // 教该课的老师,1门课程可由多个老师来教
	CourseGroups     []*CourseGroupInfo `json:"courseGroups" orm:"reverse(many)"`               // 该课在哪些课组中
	CourseName       string             `json:"name" orm:"size(50)" form:"name"`                // 课程名
	CourseProperties string             `json:"property" orm:"size(50)"`                        // 课程性质 专业必修，专业选修
	CourseBases      []*CourseBaseInfo  `json:"baseInfos" orm:"reverse(many)"`                  // 上课时间地点
	CourseScores     float64            `json:"score" orm:"digits(4);decimals(2)" form:"score"` // 课程学分
	CourseWay        string             `json:"checkWay" orm:"size(10)" form:"checkWay"`                        // 考核方式
	CourseCount      float64            `json:"courseHour" orm:"digits(5);decimals(2)"`         // 学时, 单位小时
	StudentResults   float64            `json:"grade" orm:"-"`                                  // 课程成绩
	StudentPoint     float64            `json:"point" orm:"-"`                                  // 课程绩点
	//CourseMaxNumber      int           	`json:"courseMaxNumber" orm:"max_number"`                     // 课程的上线人数 -1为无上限
	CourseCreatedTime time.Time `json:"createdTime" orm:"auto_now_add;type(datetime)"`
	CourseUpdatedTime time.Time `json:"updatedTime" orm:"auto_now;type(datetime)"`
}

// CourseBaseInfo 上课时间地点
type CourseBaseInfo struct {
	CourseBaseId          int         `json:"id" orm:"pk;auto"`                         // 主键号,id,主键,自增
	Course                *CourseInfo `json:"-" orm:"null;rel(fk);on_delete(set_null)"` // 所属课程
	CourseStartYear       int         `json:"startYear"`                                // 开始学期
	CourseEndYear         int         `json:"endYear"`                                  // 结束学期
	CourseYear            int         `json:"year"`                                     // 课程学年
	CourseStartWeek       int         `json:"startWeek"`
	CourseEndWeek         int         `json:"endWeek"`
	CourseWeek            int         `json:"week"`
	CourseStartCount      int         `json:"startCount"`
	CourseEndCount        int         `json:"endCount"`
	CourseSchool          string      `json:"school" orm:"size(50)"`  // 上课校区
	CourseAddress         string      `json:"address" orm:"size(50)"` // 上课地点
	CourseAddressFloor    int         `json:"floor"`                  // 上课楼层号
	CourseAddressNumber   int         `json:"number"`                 // 上课教室号
	CourseBaseCreatedTime time.Time   `json:"createdTime" orm:"auto_now_add;type(datetime)"`
	CourseBaseUpdatedTime time.Time   `json:"updatedTime" orm:"auto_now;type(datetime)"`
}

// CourseGroupInfo 课组管理表 连表优先级 课程 -> 老师
type CourseGroupInfo struct {
	CourseGroupId          int            `json:"id" orm:"pk;auto"`    // 课组id,主键,自增
	CourseGroupName        string         `json:"name" orm:"size(50)"` // 课组名
	IsCharge               bool           `json:"isCharge" orm:"-"`
	Courses                []*CourseInfo  `json:"courses" orm:"rel(m2m);rel_through(goweb/models.CourseGroupRel)"` // 该课组中的课程
	Teachers               []*TeacherInfo `json:"teachers" orm:"reverse(many)"`                                    // 该课组的老师
	CourseGroupCreatedTime time.Time      `json:"createdTime" orm:"auto_now_add;type(datetime)"`
	CourseGroupUpdatedTime time.Time      `json:"updatedTime" orm:"auto_now;type(datetime)"`
}

// CourseStudentRel 学生课程联系表
type CourseStudentRel struct {
	CourseStudentRelId          int          `json:"id" orm:"pk;auto"` // 学生课程联系表id,主键,自增
	Student                     *StudentInfo `json:"student" orm:"null;rel(fk);on_delete(set_null)" form:"studentId"`
	Course                      *CourseInfo  `json:"course" orm:"null;rel(fk);on_delete(set_null)" form:"courseId"`
	StudentResults              float64      `json:"grade" orm:"default(-1);digits(5);decimals(2)" form:"grade"` // 课程成绩
	StudentPoint                float64      `json:"point" orm:"default(-1);digits(4);decimals(2)"`              // 课程绩点
	CourseStudentRelCreatedTime time.Time    `json:"createdTime" orm:"auto_now_add;type(datetime)"`
	CourseStudentRelUpdatedTime time.Time    `json:"updatedTime" orm:"auto_now;type(datetime)"`
}

// CourseGroupTeacherRel 课组老师联系表
type CourseGroupTeacherRel struct {
	CourseGroupTeacherRelId          int              `orm:"pk;auto"` // 课组老师表id,主键,自增
	CourseGroup                      *CourseGroupInfo `orm:"null;rel(fk);on_delete(set_null)"`
	Teacher                          *TeacherInfo     `orm:"null;rel(fk);on_delete(set_null)"`
	IsCharge                         bool             // 是否负责该课组
	CourseGroupTeacherRelCreatedTime time.Time        `orm:"auto_now_add;type(datetime)"`
	CourseGroupTeacherRelUpdatedTime time.Time        `orm:"auto_now;type(datetime)"`
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
	CourseGroupRelId          int              `orm:"pk;auto"` // 课程与课程组表id,主键,自增
	CourseGroup               *CourseGroupInfo `orm:"null;rel(fk);on_delete(set_null)"`
	Course                    *CourseInfo      `orm:"null;rel(fk);on_delete(set_null)"`
	CourseGroupRelCreatedTime time.Time        `orm:"auto_now_add;type(datetime)"`
	CourseGroupRelUpdatedTime time.Time        `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(StudentInfo), new(ClassInfo), new(TeacherInfo),
		new(CourseInfo), new(CourseGroupInfo), new(CourseStudentRel), new(CourseGroupTeacherRel),
		new(CourseClassRel), new(CourseTeacherRel), new(CourseGroupRel), new(CourseBaseInfo))
}
