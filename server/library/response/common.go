package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

func NewPageResult(list interface{}, total int64, page int, pageSize int) *PageResult {
	return &PageResult{List: list, Total: int(total), Page: page, PageSize: pageSize}
}
