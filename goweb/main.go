package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	_ "goweb/filters"
	_ "goweb/routers"
	"os"
)
func init() {
	// need to register db driver
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error(err)
		os.Exit(-1)
	}
	DbUrl, err := web.AppConfig.String("DB_URL")
	if err != nil {
		logs.Error(err)
		os.Exit(-1)
	}
	// need to register default database
	err = orm.RegisterDataBase("default", "postgres", DbUrl)
	if err != nil {
		logs.Error(err)
		os.Exit(-1)
	}
}

func main() {
	// 配置运行跨域，配置不成功0.0, 改为配置前端
	// 前端配置：https://www.jianshu.com/p/43aa317d7683

	// 生成数据库
	orm.Debug = true

	// 启动web
	web.Run()
}
