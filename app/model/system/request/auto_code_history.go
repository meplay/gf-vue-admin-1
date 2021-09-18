package request

import (
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"gorm.io/gorm"
)

type AutoCodeHistorySearch struct {
	*common.PageInfo
}

func (s *AutoCodeHistorySearch) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		return db.Select("id,created_at,updated_at,struct_name,struct_cn_name,flag,table_name")
	}
}


func (s *AutoCodeHistorySearch) Order() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("updated_at desc")
	}
}
