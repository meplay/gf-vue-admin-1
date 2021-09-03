package common

import "gorm.io/gorm"

// Paginate 分页器
// Author [SliverHorn](https://github.com/SliverHorn)
func Paginate(info *PageInfo) func(db *gorm.DB) *gorm.DB {
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
