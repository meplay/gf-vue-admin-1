package common

type PageResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	PageInfo
}

func NewPageResult(list interface{}, total int64, info PageInfo) *PageResult {
	return &PageResult{List: list, Total: total, PageInfo: info}
}
