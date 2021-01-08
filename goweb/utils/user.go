package utils

type User struct {
	User     interface{}
	UserType int // userType:  1.Admin, 2.学生, 3.老师, 4.系主任
}

const (
	ADMIN        = 1
	STUDENT      = 2
	TEACHER      = 3
	TEACHER_HEAD = 4
)

func NewUser() *User {
	return new(User)
}
