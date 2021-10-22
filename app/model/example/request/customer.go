package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/example"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"gorm.io/gorm"
)

type CustomerCreate struct {
	example.Customer
}

type CustomerUpdate struct {
	example.Customer
}

type CustomerSearch struct {
	UserAuthorityID string `json:"-"`
	common.PageInfo
}

func (r *CustomerSearch) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB { // 如果有条件搜索 下方会自动创建搜索语句
		return db
	}
}
