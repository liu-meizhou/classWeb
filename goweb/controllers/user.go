package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"goweb/models"
	"goweb/utils"
	"sync"
)

// 	// 获取token
//	token := this.Ctx.Input.Header("token")
//	// 从缓存获取当前登录用户
//	userTypeInfoInterface, ok := userCache.Load(token)
//	if !ok {
//		this.Data["json"] = utils.NoIdentifyReJson( "请登录...")
//		this.ServeJSON()
//		return
//	}
//	userTypeInfo := userTypeInfoInterface.(*UserTypeInfo)
//	switch userTypeInfo.UserType {
//	case 1: {
//		// admin
//		break
//	}
//	case 2: {
//		// 学生
//		break
//	}
//	case 3: {
//		// 老师
//		break
//	}
//	case 4: {
//		// 系主任
//		break
//	}
//	default: {
//		this.Data["json"] = utils.NoFoundReJson( "未知用户...")
//	}
//	}
//	this.ServeJSON()

type UserController struct {
	web.Controller
}

var userCache sync.Map

func GetUser(token string) *utils.User {
	user,ok := userCache.Load(token)
	if ok {
		return user.(*utils.User)
	}
	return nil
}

func SetUser(token string, user *utils.User)  {
	// userCache有内存泄漏的风险
	userCache.Store(token, user)
}

func DeleteUser(token string)  {
	userCache.Delete(token)
}

//ParseForm(obj interface{}) error
//
//将表单反序列化到 obj 对象中。

//GetXXX(key string, def…) XXX, err
//
//从传入参数中，读取某个值。如果传入了默认值，那么在读取不到的情况下，返回默认值，否则返回错误。XXX 可以是 golang 所支持的基本类型，或者是 string, File 对象

// 后门列表
// 1865400006 学生 刘佳合  2
// 111666 老师 杨朔  3
// 100755 系主任 李传中 4

// Identify 识别身份的过滤器
func Identify(ctx *context.Context) bool {
	token := ctx.Input.Header("token")

	// 查看缓存是否有用户
	if user := GetUser(token); user != nil {
		return true
	}

	// 走后门 学生刘佳合
	if token == "1865400006" {
		userTypeInfo, err := models.Login(utils.NewTokenInfo("1865400006", "123456", utils.STUDENT))
		if err != nil {
			logs.Error(err)
			ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
			return false
		}
		SetUser(token, userTypeInfo)
		return true
	}
	// 老师 杨朔
	if token == "111666" {
		userTypeInfo, err := models.Login(utils.NewTokenInfo("111666", "123456", utils.TEACHER))
		if err != nil {
			logs.Error(err)
			ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
			return false
		}
		SetUser(token, userTypeInfo)
		return true
	}
	// 系主任 李传中
	if token == "100755" {
		userTypeInfo, err := models.Login(utils.NewTokenInfo("100755", "123456", utils.TEACHER_HEAD))
		if err != nil {
			logs.Error(err)
			ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
			return false
		}
		SetUser(token, userTypeInfo)
		return true
	}

	if token == "" {
		logs.Debug("不存在token")
		ctx.Output.JSON(utils.NoIdentifyReJson("不存在token"), web.BConfig.RunMode != web.PROD, true)
		return false
	}
	tokenInfo, err := utils.ParseToken(token)
	if err != nil {
		logs.Debug(err)
		ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
		return false
	}

	// 访问数据库
	userTypeInfo, err := models.Login(tokenInfo)
	if err != nil {
		logs.Error(err)
		ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
		return false
	}
	SetUser(token, userTypeInfo)
	return true
}

func (this *UserController) Prepare() {
	logs.Info("开始请求普通用户")
}

func (this *UserController) Finish() {
	logs.Info("结束请求普通用户")
}

// Logout 下线
func (this *UserController) Logout() {
	// 获取token
	token := this.Ctx.Input.Header("token")
	// 从缓存删除当前登录用户
	DeleteUser(token)
	this.Data["json"] = utils.SuccessReJson("成功注销...")
	this.ServeJSON()
}

