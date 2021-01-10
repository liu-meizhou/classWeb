package utils

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type NullBodyError struct {}

func (e NullBodyError) Error() string {
	return "传入不可为空"
}

var NullBody NullBodyError

// ParseBody 解析post请求的参数
func ParseBody(this *web.Controller, data interface{}) error {
	body := this.Ctx.Input.RequestBody
	if len(body)==0 {
		return NullBody
	}
	logs.Debug(string(body))
	err := json.Unmarshal(body, data)
	if err!=nil {
		err2 := this.ParseForm(data)
		if err2!=nil {
			err = fmt.Errorf(err.Error() + "   " + err2.Error())
			return err
		}
	}
	return nil
}

// ParsePageInfo 解析页面信息
func ParsePageInfo(this *web.Controller, pageInfo *PageInfo, data interface{}) error {
	// 解析查询条件
	err := ParseBody(this, pageInfo)
	if err != nil {
		if err != NullBody {
			return err
		}
	} else {
		err = ParseBody(this, data)
		if err != nil {
			return err
		}
	}
	return nil
}

