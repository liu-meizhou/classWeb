package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"goweb/utils"
)

func Login(loginInfo *utils.TokenInfo) (*utils.User, error) {
	err := utils.CheckTokenInfo(loginInfo)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	userTypeInfo := utils.NewUser()
	userTypeInfo.UserType = loginInfo.UserType
	o := orm.NewOrm()
	switch loginInfo.UserType {
	case utils.ADMIN:
		return nil, fmt.Errorf("你不是管理员")
	case utils.STUDENT:
		{
			user := &StudentInfo{StudentId: loginInfo.LoginId}
			err := o.Read(user)
			if err != nil {
				logs.Error(err)
				return nil, err
			}
			if user.StudentPassword != loginInfo.Password {
				return nil, fmt.Errorf("密码不正确")
			}
			if user.StudentType != loginInfo.UserType {
				return nil, fmt.Errorf("用户类型不匹配")
			}
			userTypeInfo.User = user
			break
		}
	case utils.TEACHER, utils.TEACHER_HEAD:
		{
			user := &TeacherInfo{TeacherId: loginInfo.LoginId}
			err := o.Read(user)
			if err != nil {
				logs.Error(err)
				return nil, err
			}
			if user.TeacherPassword != loginInfo.Password {
				return nil, fmt.Errorf("密码不正确")
			}
			if user.TeacherType != loginInfo.UserType {
				return nil, fmt.Errorf("用户类型不匹配")
			}
			userTypeInfo.User = user
			break
		}
	default:
		return nil, fmt.Errorf("未知用户类型")
	}
	return userTypeInfo, nil
}
