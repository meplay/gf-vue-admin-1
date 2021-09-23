package common

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" swaggertype:"string" example:"int 页码"`           // 页码
	PageSize int `json:"pageSize" form:"pageSize" swaggertype:"string" example:"int 每页大小"` // 每页大小
}

// GetByID get by ID
type GetByID struct {
	ID uint `json:"id" form:"page" swaggertype:"string" example:"uint 主键ID"` //
}

type GetByIDs struct {
	Ids []uint `json:"ids" form:"ids" swaggertype:"array,number"` // 主键Ids
}

// GetAuthorityId Get role by AuthorityI structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId" example:"角色ID"` // 角色ID
}

type Empty struct{}
