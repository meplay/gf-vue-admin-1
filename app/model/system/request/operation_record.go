package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"gorm.io/gorm"
)

type OperationRecordCreate struct {
	system.OperationRecord
}

type OperationRecordSearch struct {
	Path   string `json:"path" example:"请求路径"`
	Method string `json:"method" example:"请求方法"`
	Status int    `json:"status" example:"7"`
	common.PageInfo
}

func (r *OperationRecordSearch) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB { // 如果有条件搜索 下方会自动创建搜索语句
		if r.Method != "" {
			db = db.Where("method = ?", r.Method)
		}
		if r.Path != "" {
			db = db.Where("path LIKE ?", "%"+r.Path+"%")
		}
		if r.Status != 0 {
			db = db.Where("status = ?", r.Status)
		}
		return db.Order("id desc")
	}
}
