package internal

import (
	"gf-vue-admin/app/model/system/request"
	"gorm.io/gorm"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Paginate 分页器gorm共享代码
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Paginate(info *request.PageInfo) func(db *gorm.DB) *gorm.DB {
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
