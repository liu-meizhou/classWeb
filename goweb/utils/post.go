package utils

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

// ParseBody 解析post请求的参数
func ParseBody(this *web.Controller, data interface{}) error {
	body := this.Ctx.Input.RequestBody
	if len(body)==0 {
		return fmt.Errorf("传入不可为空")
	}
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

