package utils

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

// ParseBody 解析post请求的参数
func ParseBody(this *web.Controller, data interface{}) error {
	err := this.ParseForm(data)
	if err!=nil {
		body := this.Ctx.Input.RequestBody
		err2 := json.Unmarshal(body, data)
		if err2!=nil {
			return fmt.Errorf(err.Error() + "   " + err2.Error())
		}
	}
	return nil
}

