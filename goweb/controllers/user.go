package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"goweb/models"
	"goweb/utils"
	"sync"
)

type UserTypeInfo struct {
	User interface{}
	UserType int  // userType:  1.Admin, 2.学生, 3.老师, 4.系主任
}

type UserController struct {
	web.Controller
}

var cache sync.Map

func GetUserByTokenInfo(tokenInfo *utils.TokenInfo) (*UserTypeInfo, error) {
	userTypeInfo := &UserTypeInfo{
		nil,
		tokenInfo.UserType,
	}
	o := orm.NewOrm()
	switch tokenInfo.UserType {
	case 1: {
		// admin
		break
	}
	case 2: {
		// 学生
		user := &models.StudentInfo{StudentId: tokenInfo.Id}
		err := o.Read(user)
		if err == orm.ErrNoRows {
			return nil, fmt.Errorf("查询不到")
		} else if err == orm.ErrMissPK {
			return nil, fmt.Errorf("找不到主键")
		}
		userTypeInfo.User = user
		break
	}
	case 3:
	case 4: {
		// 老师和系主任
		user := &models.TeacherInfo{TeacherId: tokenInfo.Id}
		err := o.Read(user)
		if err == orm.ErrNoRows {
			return nil, fmt.Errorf("查询不到")
		} else if err == orm.ErrMissPK {
			return nil, fmt.Errorf("找不到主键")
		}
		userTypeInfo.User = user
		break
	}
	default:{
		return nil, fmt.Errorf("未知身份")
	}
	}
	return userTypeInfo, nil
}

////store 方法,添加元素
//	cache.Store(1,"a")
//	//Load 方法，获得value
//	if v,ok:=cache.Load(1);ok{
//		fmt.Println(v)
//	}
//	//LoadOrStore方法，获取或者保存
//	//参数是一对key：value，如果该key存在且没有被标记删除则返回原先的value（不更新）和true；不存在则store，返回该value 和false
//	if vv,ok:=cache.LoadOrStore(1,"c");ok{
//		fmt.Println(vv)
//	}
//	if vv,ok:=cache.LoadOrStore(2,"c");!ok{
//		fmt.Println(vv)
//	}
//	//遍历该map，参数是个函数，该函数参的两个参数是遍历获得的key和value，返回一个bool值，当返回false时，遍历立刻结束。
//	cache.Range(func(k,v interface{})bool{
//		fmt.Print(k)
//		fmt.Print(":")
//		fmt.Print(v)
//		fmt.Println()
//		return true
//	})

//ParseForm(obj interface{}) error
//
//将表单反序列化到 obj 对象中。

//GetXXX(key string, def…) XXX, err
//
//从传入参数中，读取某个值。如果传入了默认值，那么在读取不到的情况下，返回默认值，否则返回错误。XXX 可以是 golang 所支持的基本类型，或者是 string, File 对象

// Identify 识别身份的过滤器
func Identify(ctx *context.Context) bool {
	token := ctx.Input.Header("token")

	// 查看缓存是否有用户
	if _ ,ok:=cache.Load(token);ok{
		return true
	}

	// 走后门 学生刘佳合
	if token == "1865400006" {
		userTypeInfo, err := GetUserByTokenInfo(utils.NewTokenInfo(1865400006, 2))
		if err != nil {
			logs.Error(err)
			ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
			return false
		}
		cache.Store(token, userTypeInfo)
		return true
	}
	// 老师 杨朔
	if token == "111666" {
		userTypeInfo, err := GetUserByTokenInfo(utils.NewTokenInfo(111666, 3))
		if err != nil {
			logs.Error(err)
			ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
			return false
		}
		cache.Store(token, userTypeInfo)
		return true
	}
	// 系主任 李传中
	if token == "100755" {
		userTypeInfo, err := GetUserByTokenInfo(utils.NewTokenInfo(100755, 4))
		if err != nil {
			logs.Error(err)
			ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
			return false
		}
		cache.Store(token, userTypeInfo)
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
	// 查看缓存是否有用户
	if _ ,ok:=cache.Load(token);ok{
		return true
	}
	// 访问数据库
	userTypeInfo, err := GetUserByTokenInfo(tokenInfo)
	if err != nil {
		logs.Error(err)
		ctx.Output.JSON(utils.NoIdentifyReJson(err.Error()), web.BConfig.RunMode != web.PROD, true)
		return false
	}
	cache.Store(token, userTypeInfo)
	// cache有内存泄漏的风险
	return true
}

func (this *UserController) Prepare() {
	logs.Info("开始请求普通用户")
}

func (this *UserController) Finish() {
	logs.Info("结束请求普通用户")
}

// GetCourse 获取课表
func (this *UserController) GetCourse() {
	// 获取token
	token := this.Ctx.Input.Header("token")
	// 从缓存获取当前登录用户
	userTypeInfoInterface, ok := cache.Load(token)
	if !ok {
		this.Data["json"] = utils.NoIdentifyReJson( "请登录...")
		this.ServeJSON()
		return
	}
	userTypeInfo := userTypeInfoInterface.(*UserTypeInfo)
	switch userTypeInfo.UserType {
	case 1: {
		// admin
		break
	}
	case 2: {
		// 学生
		studentInfo := userTypeInfo.User.(*models.StudentInfo)
		if studentInfo.Courses == nil {
			_ = models.GetStudentCourse(studentInfo)
		}
		this.Data["json"] = utils.SuccessReJson(studentInfo.Courses)
		break
	}
	case 3: {
		// 老师
		break
	}
	case 4: {
		// 系主任
		break
	}
	default: {
		this.Data["json"] = utils.NoFoundReJson( "未知用户...")
	}
	}
	this.ServeJSON()
}

// ExportCourse 导出课表
func (this *UserController) ExportCourse() {
	this.Data["json"] = utils.NewReturnJson(200, "Get请求导出课表", "success")
	this.ServeJSON()
}

// ChooseCourse 选课
func (this *UserController) ChooseCourse() {
	this.Data["json"] = utils.NewReturnJson(200, "Get请求获取选课列表，Post请求进行选课", "success")
	this.ServeJSON()
}

// CourseGrade 获取课程成绩等信息
func (this *UserController) CourseGrade() {
	//req := struct{ Title string }{}
	//if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
	//	this.Ctx.Output.SetStatus(400)
	//	this.Ctx.Output.Body([]byte("empty title"))
	//	return
	//}
	//_ = this.ParseForm(&req)
	logs.Info(this.Ctx.Request.Method)
	this.Data["json"] = utils.NewReturnJson(200, "Get请求获取学生课程成绩, Post请求进行设置成绩", "success")
	_ = this.ServeJSON()
}
