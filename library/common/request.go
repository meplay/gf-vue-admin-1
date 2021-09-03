package common

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// GetByID get by ID
type GetByID struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}

type GetByIDs struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by AuthorityI structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}
