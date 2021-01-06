package tests

import (
	"github.com/beego/beego/v2/core/logs"
	"goweb/controllers"
	"goweb/utils"
	"testing"
)

func TestGetUserByTokenInfo(t *testing.T) {
	user, err := controllers.GetUserByTokenInfo(utils.NewTokenInfo(1865400006, 2))
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info(user)
}

