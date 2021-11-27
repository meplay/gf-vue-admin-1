package common

import "gorm.io/gorm"

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" swaggertype:"string" example:"int 页码"`
	PageSize int `json:"pageSize" form:"pageSize" swaggertype:"string" example:"int 每页大小"`
}

// Paginate 分页器
// Author [SliverHorn](https://github.com/SliverHorn)
func (r *PageInfo) Paginate(info *PageInfo) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case info.PageSize > 100:
			info.PageSize = 100
		case info.PageSize < 0:
			info.PageSize = 10
		}
		offset := info.PageSize * (info.Page - 1)
		return db.Offset(offset).Limit(info.PageSize)
	}
}

// GetByID get by ID
type GetByID struct {
	ID uint `json:"id" form:"page" swaggertype:"string" example:"uint 主键ID"`
}

type GetByIDs struct {
	Ids []uint `json:"ids" form:"ids" swaggertype:"array,number"`
}

// GetAuthorityId Get role by AuthorityI structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId" example:"角色ID"`
}

type Empty struct{}
