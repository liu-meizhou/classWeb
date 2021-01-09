package utils

import "time"

// CourseStudentRel 学生课程联系表
type CourseStudentRel struct {
	CourseStudentRelId          int       `json:"id"` // 学生课程联系表id,主键,自增
	StudentId                   string    `json:"studentId" form:"studentId"`
	CourseId                    string    `json:"courseId" form:"courseId"`
	StudentResults              float64   `json:"grade" form:"grade"` // 课程成绩
	StudentPoint                float64   `json:"point"`              // 课程绩点
	CourseStudentRelCreatedTime time.Time `json:"createdTime"`
	CourseStudentRelUpdatedTime time.Time `json:"updatedTime"`
}
