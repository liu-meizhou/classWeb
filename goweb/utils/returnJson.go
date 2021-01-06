package utils

type ReturnJson struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// NewReturnJson 创建一个返回json
func NewReturnJson(code int, data interface{}, message string) *ReturnJson {
	return &ReturnJson{
		code,
		data,
		message,
	}
}

func SuccessReJson(data interface{}) *ReturnJson{
	return NewReturnJson(200, data, "success")
}

func ErrorReJson(data interface{}) *ReturnJson{
	return NewReturnJson(500, data, "error")
}

func NoFoundReJson(data interface{}) *ReturnJson{
	return NewReturnJson(404, data, "noFound")
}

func NoIdentifyReJson(data interface{}) *ReturnJson{
	return NewReturnJson(401, data, "无权限")
}
