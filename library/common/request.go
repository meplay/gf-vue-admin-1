package common

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" example:"1"`          // 页码
	PageSize int `json:"pageSize" form:"pageSize" example:"10"` // 每页大小
}

// GetByID get by ID
type GetByID struct {
	ID float64 `json:"id" form:"page" example:"7"` // 主键ID
}

func (c *GetByID) ToUint() uint {
	return uint(c.ID)
}

type GetByIDs struct {
	Ids []uint `json:"ids" form:"ids" swaggertype:"array,number"` // 主键Ids
}

// GetAuthorityId Get role by AuthorityI structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId" example:"角色ID"` // 角色ID
}

type Empty struct{}