package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "goweb/routers"
)


func main() {
	web.Run()
}