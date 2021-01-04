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
