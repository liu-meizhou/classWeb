package utils

// PageInfo 页面信息
type PageInfo struct {
	Limit   int
	Offset  int
	OrderBy string
}

// NewPageInfo
func NewPageInfo(limit, offset int, orderBy string) *PageInfo {
	return &PageInfo{
		limit,
		offset,
		orderBy,
	}
}
