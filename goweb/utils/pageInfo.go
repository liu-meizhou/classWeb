package utils

// PageInfo 页面信息
type PageInfo struct {
	PageNo  int         `json:"pageNo" form:"pageNo"` // 当前页面
	Total   int64       `json:"total" form:"total"`   // 数据库总数
	Limit   int         `json:"limit" form:"limit"`   // 一页长度
	Offset  int         `json:"offset" form:"offset"` // 当PageNo为-1时，Offset生效，偏移量
	OrderBy string      `json:"orderBy" form:"orderBy"`
	IsDesc  bool        `json:"isDesc" form:"isDesc"` // 当orderBy不为"" 生效，true表示降序，false表示升序
	Lists   interface{} `json:"lists" form:"lists"`   // 查询结果列表
}
